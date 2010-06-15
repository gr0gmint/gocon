package main

import "fmt"
import . "gocon"
import "goprotobuf.googlecode.com/hg/proto"

type Session map[string]interface{}

type InitProtoHandler struct {
    ProtoHandler
    Queue chan []byte
}

type InWorldProtoHandler struct { //Need this?
    ProtoHandler
}

type PushProtoHandler struct {
    ProtoHandler
    HotRoutine
}

func NewPushProtoHandler(p *ProtoProxy) *PushProtoHandler  {
    push := new(PushProtoHandler)
    push.Proxy =p 
    push.Init()
    go push.HotStart()
    return push
}
func (this *PushProtoHandler) UpdatePlayerCoordinate(x,y int) {
    fmt.Printf("Updating player coordinate\n")
    m := NewUpdatePlayerCoord()
    m.Coord = new(Coordinate)
    m.Coord.X = proto.Int32(int32(x))
    m.Coord.Y = proto.Int32(int32(y))
    data,err := proto.Marshal(m)
    if err != nil {
        fmt.Printf("E: %s", err)
        return
    }
    this.Proxy.SendMsg(data,PORT_PUSH, Server_UPDATELOCATION,  false)
}
func (this *PushProtoHandler) Main() {
    
}
//func (this *PushProtoHandler) PushVisibleObjects(objects []*GObject) {
//}


func NewInWorldProtoHandler(p *ProtoProxy) *InWorldProtoHandler {
    h := new(InWorldProtoHandler)
    h.Proxy = p
    h.Init()
    return h
}


func NewInitProtoHandler(p *ProtoProxy) *InitProtoHandler {
    h := new(InitProtoHandler)
    h.Proxy = p
    h.Init()
    return h
}

func (this *InitProtoHandler) Cleanup() {
    
}
func (this *InitProtoHandler) Main() {
    defer this.Cleanup()
    fmt.Printf("InitProtoHandler·Main\n")
    header, data := this.Proxy.ReadMsg(0)
    if header == nil || *header.Type != Client_JOIN {
        fmt.Printf("Invalid data!\n")
        return
    }
    //Player joins
    joinmsg := NewClientJoin() 
    player := NewPlayer()
    proto.Unmarshal(data, joinmsg)
    player.Name = *joinmsg.Playername
    worldhandler := GlobalRoutines["worldhandler"].(*WorldHandler)
    coord := worldhandler.World.GetCoord(50,50)
    fmt.Printf("%d %d\n", coord.X, coord.Y)
    
    pusher := NewPushProtoHandler(this.Proxy)
    go pusher.Main()
    player.ProtoHandlers[PORT_PUSH] = pusher
    
    worldhandler.PlacePlayer(player,coord)
    this.Acceptbool()
    //Pass on to another handler
    session := make(Session)
    session["player"] = player
    

    
    
    inworld := NewInWorldProtoHandler(this.Proxy)
    player.ProtoHandlers[PORT_MAIN] = inworld
    inworld.Main(&session)
    

}   



func (this *InWorldProtoHandler) Main(session *Session) {
    fmt.Printf("InWorldProtoHandler·Main\n")
    w := GlobalRoutines["worldhandler"].(*WorldHandler)
    for {
        header,data := this.Proxy.ReadMsg(0)
        if header == nil && data == nil {
            fmt.Printf("Connection closed\n")
            return
        }
        t := *header.Type
        fmt.Printf("%s received\n", Client_Type_name[t])
        
        player := (*session)["player"].(*Player)
        switch {
            case t == Client_WALK:
                m := NewClientWalk()
                err := proto.Unmarshal(data, m)
                if err != nil { fmt.Printf("E:%s\n", err) } else {
                    fmt.Printf("%s\n", ClientWalk_Direction_name[int32(*m.Direction)])
                    switch {
                        case *m.Direction == ClientWalk_DIRECTION_UP:
                            w.PlayerMove(player, DIRECTION_UP)
                        case *m.Direction == ClientWalk_DIRECTION_DOWN:
                            w.PlayerMove(player, DIRECTION_DOWN)
                        case *m.Direction == ClientWalk_DIRECTION_LEFT:
                            w.PlayerMove(player, DIRECTION_LEFT)
                        case *m.Direction == ClientWalk_DIRECTION_RIGHT:
                            w.PlayerMove(player, DIRECTION_RIGHT)
                    }
                }
            
        }
        
    }

}   

