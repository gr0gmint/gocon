package gocon

/**********TEMPLATE FOR HOT FUNCTION*********
    h := NewHot(func(shared map[string]interface{}){
        self := shared["self"].(*GenericHot)
    })
    this.queryHot(h)
    answer:=<-h.Answer
*********************************************/
import "fmt"

type Hot interface { //Hot code "swapping"
    Unpack(map[string]interface{})
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
    HotChan chan Hot
    hotlock bool
}
func (this *HotRoutine) QueryHot(h Hot) {
    //if !this.hotlock {
     this.HotChan<-h
    //} else { //We're already in another hot, which means the hot called another hot
   //     shared := make(map[string]interface{})
    //    shared["self"] = h
    //    go h.Unpack(shared)
    //}

}

func (this *HotRoutine) HotStart() {
    this.HotChan = make(chan Hot)
    for {
        fmt.Printf("DEBUG: listening to this.HotChan\n")
        h := <-this.HotChan
        fmt.Printf("got a hot!\n")
        shared := make(map[string]interface{})
        shared["self"] = h
        fmt.Printf("Trying to unpack\n")
        fmt.Printf("%s\n", h)
        h.Unpack(shared)

    }
}
