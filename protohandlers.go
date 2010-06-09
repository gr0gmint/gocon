package main

import "fmt"
import . "gocon"
import "goprotobuf.googlecode.com/hg/proto"


type InitProtoHandler struct {
    ProtoHandler
    Queue chan []byte
}

type InWorldProtoHandler struct { //Need this?
    ProtoHandler
}


func NewInitProtoHandler(p *ProtoProxy) *InWorldProtoHandler {
    
}
func (this *InWorldProtoHandler) Handle(header *Header, data []byte) {
    
}




func NewInitProtoHandler(p *ProtoProxy) *InitProtoHandler {
    h := new(InitProtoHandler)
    h.Queue = make(chan []byte)
    h.Proxy = p
    h.Init()
    return h
}

func (this *InitProtoHandler) Handle(header *Header, data []byte) {
    fmt.Printf("InitProtoHandler·Handle\n")
    this.Queue <- data
}

func (this *InitProtoHandler) Main() {
    data := <-this.Queue
    
    //Player joins
    joinmsg := NewClientJoin() 
    proto.Unmarshal(data, joinmsg)
    playername := *joinmsg.Playername
    
      
}   

