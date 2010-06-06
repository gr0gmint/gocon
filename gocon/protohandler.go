package gocon

import "os"
import "goprotobuf.googlecode.com/hg/proto"
import "net"

type ProtoProxy struct {
    HotRoutine
    Buffer []byte
    SBuffer []byte
    Conn *net.Conn
    Default *IProtoHandler
    Handlers map[string]*IProtoHandler  
    headersize int
    Filter *ProtoFilter
}
type ProtoFilter interface {
     Check(header *Header) bool
}

type ProtoHandler struct {
    HotRoutine
    Proxy *ProtoProxy
    
}
type IProtoHandler interface {
    Handle([]byte)
}

func (this *ProtoProxy) Init() {
    temphdr := NewHeader()
    *temphdr.Size = 4123
    data, _ := proto.Marshal(temphdr)
    this.headersize = len(data)
    this.Buffer = make([]byte, 10000)
    this.SBuffer = make([]byte, 10000)
    go this.HotStart()
    this.Handlers = make(map[string]*IProtoHandler)
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

func (this *ProtoProxy) Read (size int) ([]byte, os.Error) {
    
    if size <= 0 {
        n,err := this.Conn.Read(this.Buffer)
        return this.Buffer[0:n], err 
    } 
    var total int
    for ; total < size; {
            n, err := this.Conn.Read(this.Buffer[total:size])
            total += n
            if err !=nil { return this.Buffer[0:total],err }
        }
        
    return this.Buffer[0:total], nil
}

func (this *ProtoProxy) readMsg() (*Header,[]byte, os.Error) {

    //Read header first
    data,err := this.Read(this.headersize)
    if err != nil  {
        header := NewHeader()
        proto.Unmarshal(data, header)
        data,err := this.Read(int(*header.Size))
        return header,data,err
    } 
    return nil,nil,err
} 
func (this *ProtoProxy) Send(data []byte, handlername string) {
    h := NewHot(func(shared map[string]interface{}){
        //self := shared["self"].(*GenericHot)
        header := NewHeader()
        *header.Handler = handlername
        *header.Size = int32(len(data))
        hdrdata,_ := proto.Marshal(header)
        hdrlen := len(hdrdata)
        copy(this.SBuffer[0:hdrlen], hdrdata)
        copy(this.SBuffer[hdrlen:hdrlen+len(data)], data)
        this.Conn.Write(this.SBuffer[0:hdrlen+len(data)])
    })
    this.queryHot(h)

}


func (this *ProtoProxy) Main() {
    this.Init()
    for {
        header, data, err := this.readMsg(); 
        if err == nil {
            this.Conn.Close()
            return

        } else {
            this.Handlers[*header.Handler].Handle(data)
        }
        
    }
}




/*
func (this *ProtoHandler) Acceptbool() {
        msg := NewAcceptBool()
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

*/



