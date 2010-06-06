package main

import "./pwan"
import "fmt"
import . "gocon"

type InitProtoHandler struct {
    ProtoHandler
    Queue chan []byte
}

type InWorldProtoHandler struct {
    ProtoHandler
}






func NewInitProtoHandler(p *ProtoProxy) *InitProtoHandler {
    h := new(InitProtoHandler)
    h.Proxy = p
    h.Queue = make(chan []byte)
    h.Init()
    return h
}

func (this *InitProtoHandler) Handle(data []byte) {
    this.Queue <- data
}

func (this *InitProtoHandler) Main() {
        data := <-this.Queue
        joinmsg := pwan.NewClientJoin() 
        proto.Unmarshal(data, joinmsg)
        fmt.Printf("%s\n",joinmsg.Playername)
    /*
        if header, _, err := this.RecvMsg(joinmsg); err {
            this.Conn.Close()
            return
        } else {
            this.Player = NewPlayer()
            this.Player.Name = joinmsg.Playername
            
            worldhandler := GlobalRoutines["worldhandler"].(*WorldHandler)
            worldhandler.PlacePlayer(this.Player, world.GetCoord(50,50))
            inworldhandler := new(InWorldProtoHandler)
            inworldhandler.Conn = this.Conn
            go inworldhandler.Main()

        }
        */
    
}   

