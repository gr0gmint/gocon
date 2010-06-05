package main

import "rand"
import . "gocon"

type Hot interface { //Hot code "swapping"
    Unpack(interface{}) int
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
    F func(interface{})
    Answer chan *Message
}
func NewHot(f func(interface{})) *GenericHot {
    h := new(GenericHot)
    h.F = f
    h.Answer = make(chan *Message)
    return h
}
func (this *GenericHot) Unpack(data interface{}) {
    F(data)
}
