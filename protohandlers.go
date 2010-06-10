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



func (this *InitProtoHandler) Main() {
    fmt.Printf("InitProtoHandler·Main\n")
    header, data := this.Proxy.ReadMsg(0)
    fmt.Printf("After chan%s\n",header)
    //Player joins
    joinmsg := NewClientJoin() 
    player := new(Player)
    proto.Unmarshal(data, joinmsg)
    player.Name = *joinmsg.Playername
    worldhandler := GlobalRoutines["worldhandler"].(*WorldHandler)
    coord := worldhandler.World.GetCoord(50,50)
    fmt.Printf("%d %d\n", coord.X, coord.Y)
    worldhandler.PlacePlayer(player,coord)
    this.Acceptbool()

}   

