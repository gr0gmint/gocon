package gocon

/**********TEMPLATE FOR HOT FUNCTION*********
    h := NewHot(func(shared map[string]interface{}){
        self := shared["self"].(*GenericHot)
    })
    this.queryHot(h)
    answer:=<-h.Answer
*********************************************/


type Hot interface { //Hot code "swapping"
    Unpack(interface{}) 
}
type NamedHot interface {
    Hot
    Type() string
}
type GenericHot struct {
    F func(map[string]interface{})
    Answer chan *Message
}
func NewHot(f func(map[string]interface{})) *GenericHot {
    h := new(GenericHot)
    h.F = f
    h.Answer = make(chan *Message)
    return h
}
func (this *GenericHot) Unpack(data map[string]interface{}) {
    this.F(data)
}
type HotRoutine struct {
    Routine
    HotChan chan *Hot
    hotlock bool
}
func (this *HotRoutine) queryHot(h *GenericHot) {
    if !this.hotlock {
        go func() { this.Chan<-h; }()
    } else { //We're already in another hot, which means the hot called another hot
        shared := make(map[string]interface{})
        shared["self"] = h
        go h.F(shared)
    }

}

func (this *HotRoutine) HotStart() {
    this.HotChan = make(chan *Hot)
    for {
        h := <-this.HotChan
        shared := make(map[string]interface{})
        shared["self"] = h
        h.Unpack(shared)

    }
}
