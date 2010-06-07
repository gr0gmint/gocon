package main

import . "gocon"
import . "container/vector"

type Broadcast struct {
    Data map[string]interface{}
}
func NewBroadcast () {
    b := new(Broadcast)
    b.Data = make(map[string]interface{})
}

type BroadcastServer struct {
    HotRoutine
    Filters  *Vector 
}

func    (this *BroadcastServer) Setup() {
    
}

func (this *BroadcastServer) AddFilter(filter *Filter) {
    this.Filters.Push(filter)
}

func NewBroadcastServer() *BroadcastServer {
    b := new(BroadcastServer)
    b.Init()
    return b
}

func (this *BroadcastServer) Main() {
    go this.HotStart()
    this.Filters = new(Vector)
    for {
        m := <- this.Chan 
        switch  {
            case m.Key == "broadcast":
                
        }
    }
}

func (this *BroadcastServer) Broadcast(b *Broadcast) {
    c := this.Filters.Iter()
    for { filter,ok := <-c; if !ok {break;}
         filter.(*Filter).ParseBroadcast(b)
    }
}
