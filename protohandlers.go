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



func (this *InWorldProtoHandler) Handle(header *Header, data []byte) {
    
}




func NewInitProtoHandler(p *ProtoProxy) *InitProtoHandler {
    h := new(InitProtoHandler)
    h.Queue = make(chan []byte)
    h.Proxy = p
    h.Init()
    return h
}
func ABool(accept bool) []byte{
    a := NewAcceptBool()
    a.Accept = proto.Bool(accept)
    data,_ := proto.Marshal(a)
    return data
}

func (this *InitProtoHandler) Handle(header *Header, data []byte) {
    fmt.Printf("InitProtoHandler·Handle\n")
    this.Queue <- data
}

func (this *InitProtoHandler) Main() {
            fmt.Printf("InitProtoHandler·Main\n")
    data := <-this.Queue
    fmt.Printf("After chan\n")
    //Player joins
    joinmsg := NewClientJoin() 
    fmt.Printf("1\n")
    player := new(Player)
        fmt.Printf("2\n")
    proto.Unmarshal(data, joinmsg)
        fmt.Printf("3\n")
    player.Name = *joinmsg.Playername
        fmt.Printf("4\n")
    worldhandler := GlobalRoutines["worldhandler"].(*WorldHandler)
        fmt.Printf("5\n")
    coord := worldhandler.World.GetCoord(50,50)
        fmt.Printf("6\n")
    worldhandler.PlacePlayer(player,coord)
    fmt.Printf("7\n")
    a := ABool(true)
    this.Proxy.Send(a,0,Server_ANSWERBOOL)

}   

