package main

import . "gocon"
import . "container/vector"
import "fmt"

type Broadcast struct {
    Data map[string]interface{}
}
func NewBroadcast () *Broadcast {
    b := new(Broadcast)
    b.Data = make(map[string]interface{})
    return b
}

type BroadcastServer struct {
    HotRoutine
    Filters  *Vector
}
var FDistanceFromPlayer *FilterDistanceFromPlayer
var FPlayer *FilterPlayer
func (this *BroadcastServer) Setup() {
    FDistanceFromPlayer = NewFilterDistanceFromPlayer() 
    this.AddFilter(FDistanceFromPlayer)
    FPlayer := NewFilterPlayer()
    this.AddFilter(FPlayer)
    
}

func (this *BroadcastServer) AddFilter(filter Filter) {
    this.Filters.Push(filter)
}

func NewBroadcastServer() *BroadcastServer {
    b := new(BroadcastServer)
    b.Init()
    b.Filters = new(Vector)
  
    return b
}


func (this *BroadcastServer) Main() {
    go this.HotStart()
    for {
        m := <- this.Chan 
        switch  {
            case m.Key == "broadcast":
                b := m.Data["b"].(*Broadcast)
                this.Broadcast(b)
        }
    }
}

func (this *BroadcastServer) Broadcast(b *Broadcast) {
    fmt.Printf("Broadcasting\n")
    c := this.Filters.Iter()
    for { filter := <-c
        if filter == nil {break }
         filter.(Filter).ParseBroadcast(b)
    }
}
