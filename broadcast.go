package main

import "container/vector"

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

type (*this BroadcastServer) Setup() {
    
}

type (*this BroadcastServer) AddFilter(filter *Filter) {
    this.Filters.Append(filter)
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
        m := <- r.Chan 
        switch  {
            case m.MsgType == "broadcast":
                
        }
    }
}

func (this *BroadcastServer) Broadcast(b *Broadcast) {
    c := this.Filters.Iter()
    for { filter.(*Filter),ok := <-c; if !ok {break;}
         filer.ParseBroadcast(b)
    }
}
