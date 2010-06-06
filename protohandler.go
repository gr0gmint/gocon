package main

import "os"
import "goprotobuf.googlecode.com/hg/proto"
import "./pwan"


type ProtoProxy struct {
    HotRoutine
    Default *IProtoHandler
    Handlers map[string]*IProtoHandler  
    headersize int
    Filter *ProtoFilter
}
type ProtoFilter interface {
     Check(header *pwan.Header) bool
}

type ProtoHandler struct {
    HotRoutine
    Proxy *ProtoProxy
    
}
type IProtoHandler interface {
    Handle([]byte)
}


type InitProtoHandler struct {
    ProtoHandler
    Queue chan []byte
}

type InWorldProtoHandler struct {
    ProtoHandler
}


func (this *ProtoProxy) Init() {
    temphdr := pwan.NewHeader()
    temphdr.Size = 0xf00
    data, _ := proto.Marshal(temphdr)
    this.headersize = len(data)
    this.Buffer := make([]byte, 10000)
    this.SBuffer := make([]byte, 10000)
    go this.HotStart()
    p.Handlers = make(map[string]*IProtoHandler)
}

func NewProtoProxy(conn *net.Conn, def *IProtoHandler) *ProtoProxy {
    p := new(ProtoProxy)
    p.Conn = conn
    p.Default = def
    p.Init()
    return p
}

func (this *ProtoProxy) AddHandler(name string, handler *IProtoHandler) {
    this.Handlers[name] = handler
}
func (this *ProtoProxy) SetDefault(def *IProtoHandler) {
    this.Default = def
}
func (this *ProtoProxy) RemoveHandler(name string) {
    this.Handlers[name] = nil
}

func (this *ProtoProxy) Read(size int) []byte, os.Error {
    
    if size == nil || size <= 0 {
        return this.Conn.Read(this.Buffer)
    } else {
        for total := 0; total < size {
            n, err := this.Conn.Read(this.Buffer[total:size])
            total += n
            if err { return this.Buffer[0:total],err }
        }
        return this.Buffer[0:total], nil
    }
}

func (this *ProtoProxy) readMsg() *pwan.Header,[]byte, os.error() {

    //Read header first
    if data,err := this.Read(this.headersize); !err  {
        header := NewHeader()
        proto.Unmarshal(data, header)
        data,err := this.Read(header.Size)
        return header,data,err
    } else {
        return nil,nil,err
    }
} 
func (this *ProtoProxy) Send(data []byte, handlername string = nil) {
    h := NewHot(func(shared map[string]interface{}){
        self := shared["self"].(*GenericHot)
        header := pwan.NewHeader()
        header.Handler = handlername
        header.Size = len(data)
        hdrdata,_ := proto.Marshal(header)
        hdrlen := len(hdrdata)
        *this.SBuffer[0:hdrlen] = *hdrdata
        *this.SBuffer[hdrlen:hdrlen+len(data)] = *data
        this.Conn.Send(this.SBuffer[0:hdrlen+len(data)])
    })
    this.queryHot(h)

}


func (this *ProtoProxy) Main() {
    this.Init()
    for {
        if header, data, err := this.recvMsg(); !err {
            this.Handlers[header.Handler].Handle(data)
        } else {
            this.Conn.Close()
            return
        }
        
    }
}





func (this *ProtoHandler) Acceptbool() {
        msg := pwan.NewAcceptBool()
        msg.Accept = true
        data := proto.Marshal(msg)
        this.Proxy.Send(data)
}
func (this *ProtoHandler) Declinebool() {
        msg := pwan.NewAcceptBool()
        msg.Accept = false
        data := proto.Marshal(msg)
        this.Proxy.Send(data)
}










func NewInitProtoHandler(p *ProtoProxy) *InitProtoHandler {
    h := new(ProtoHandler)
    h.Proxy = p
    h.Queue = make(chan []byte)
    h.Init()
    return h
}

func (this *InitProtoHandler) Handle(data []byte
    this.Queue <- data
}

func (this *InitProtoHandler) Main() {
        data := <-this.Queue 
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
