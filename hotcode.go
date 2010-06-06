package main

import "rand"
import . "gocon"

/**********TEMPLATE FOR HOT FUNCTION*********
    h := NewHot(func(shared map[string]interface{}){
        self := shared["self"].(*GenericHot)
    })
    this.queryHot(h)
    answer:=<-h.Answer
*********************************************/


type Hot interface { //Hot code "swapping"
    Unpack(*Hot,interface{}) 
}
type NamedHot interface {
    Hot
    Type() string
}
type HotPlayerJoin struct {
    Player *Player
}
func (e *HotPlayerJoin) Type() string { return "PlayerJoin" }
func (e *HotPlayerJoin) Unpack(shared map[string]interface{})  int {
    world := shared["world"].(*World)
    rchan := shared["rchan"].(chan *Message)
    randomcoord := world.GetCoord(rand.Int() % WORLD_SIZE_X, rand.Int() % WORLD_SIZE_Y)
   
    m := NewMessage()
    if world.PlacePlayer(e.Player, randomcoord) {

        m.Key = "accepted"
        
    } else {

        m.Key = "declined"
    }
    rchan <- m
    return 0
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
    F(data)
}
type HotRoutine struct {
    Routine
    HotChan chan *Hot
}
func (r *HotRoutine) queryHot(h *Hot) {
    if !this.hotlock {
        go func() { this.Chan<-h; }()
    } else { //We're already in another hot, which means the hot called another hot
        shared := make(map[string]interface{})
        shared["self"] = h
        go h.F(shared)()
    }

}

func (this *HotRoutine) HotStart() {
    this.HotChan := make(chan *Hot)
    for {
        h := <-this.HotChan
        shared := make(map[string]interface{})
        shared["self"] = h
        h.Unpack(shared)

    }
}
