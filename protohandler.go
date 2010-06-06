package main

import "os"
import "goprotobuf.googlecode.com/hg/proto"
import "./pwan"


type InitProtoHandler struct {
    Routine
    ProtoHandler *net.Conn
    Player *Player
    Buffer []byte
    headersize int
}
type ProtoProxy struct {
    HotRoutine
    Default *ProtoHandler
    Handlers map[string]*ProtoHandler   
}
func NewProtoProxy(conn *net.Conn, def *ProtoHandler) {
    p := new(ProtoProxy)
    p.Conn = conn
    p.Default = def
    p.Handlers = make(map[string]*ProtoHandler)
}

func (this *ProtoProxy) AddHandler(name string, handler *ProtoHandler) {
    h := NewHot(func(shared map[string]interface{}){
        //self := shared["self"].(*GenericHot)
        this.Handlers[name] = handler
    })
    this.queryHot(h)
}
func (this *ProtoProxy) Main() {
    this.Init()
    go this.Start()
    for {
        
    }
}

type ProtoHandler interface {
    Handle(interface{})
}

func NewProtoHandler(c *Conn) *ProtoHandler {
    p := new(ProtoHandler)
    p.Conn = c
    return p
}
func (this *ProtoHandler) Read(size int) []byte, os.Error {
    
    if size == 0 {
        return this.Conn.Read(this.Buffer)
    } else {
        for total := 0; total < n {
            n, err := this.Conn.Read(this.Buffer[total:size])
            total += n
            if err { return this.Buffer[0:total],err }
        }
        return this.Buffer[0:total], nil
    }
}
type ConnHandler struct { 
    Routine
    Conn *Conn
}
func NewConnHandler(conn *Conn) *ConnHandler {
    c := new(ConnHandler)
    c.Conn = conn
    c.Init()
    return c
}
func (this *ConnHandler) Main() {
    for {
        msg, _ := this.ReceiveMessage()
        switch {
            case msg.Key == "hot":
                shared := make(map[string]interface{})
                shared
                msg.Data["hot"].(*Hot)
                hot.Un
        }
    }
}

func (this *ProtoHandler) RecvMsg(msg interface{}) *pwan.Header,[]byte, os.error() {

    //Read header first
    if data,err := this.Read(this.headersize); !err  {
        header := NewHeader()
        proto.Unmarshal(data, header)
        data,err := this.Read(header.Size)
        if msg == nil {
            return header,data,err
        }
        proto.Unmarshal(data, msg)
        return header,nil,err
    } else {
        return nil,nil,err
    }
} 

func (this *ProtoHandler) SendMsg(msg interface{})  os.Error {
    header := NewHeader()
    header.Size = len(data)
    hdrdata,_ := proto.Marshal(header)
    
    this.Conn.Send(hdrdata)
    
    data,_ := proto.Marshal(msg)
    this.Conn.Send(data)
    return nil
}
func (this *ProtoHandler) Acceptbool() {
        msg := pwan.NewAcceptBool()
        msg.Accept = true
        this.SendMsg(msg)
}
func (this *ProtoHandler) Declinebool() {
        msg := pwan.NewAcceptBool()
        msg.Accept = false
        this.SendMsg(msg)
}
func (this *ProtoHandler) Init {
    temphdr := pwan.NewHeader()
    temphdr.Size = 0xf00
    data, _ := proto.Marshal(temphdr)
    this.headersize = len(data)
    
    p.Buffer := make([]byte, 10000)
}



func (this *InitProtoHandler) Main() {
    this.Init()
    
    joinmsg := pwan.NewClientJoin()
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
    
}   

type InWorldProtoHandler struct {
    ProtoHandler
}
func (this *InWorldProtoHandler) Cleanup () {
    //remove player from world
}

func (this *InWorldProtoHandler) Main() {
    this.Init()
    defer this.Cleanup()
    
    
    //go ObjectPusher
    for {
        header, data, err := this.RecvMsg(nil); err {
            this.Conn.Close()
            return
        }
        switch {
            
        }
        
    }
}
