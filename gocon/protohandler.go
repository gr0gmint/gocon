package gocon

import "os"
import "goprotobuf.googlecode.com/hg/proto"

type ProtoProxy struct {
    HotRoutine
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

func (this *ProtoProxy) Read (size int) ([]byte, os.Error) {
    
    if size == nil || size <= 0 {
        return this.Conn.Read(this.Buffer)
    } else {
        for total := 0; total < size; {
            n, err := this.Conn.Read(this.Buffer[total:size])
            total += n
            if err { return this.Buffer[0:total],err }
        }
        return this.Buffer[0:total], nil
    }
}

func (this *ProtoProxy) readMsg() (*pwan.Header,[]byte, os.Error) {

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
func (this *ProtoProxy) Send(data []byte, handlername string) {
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






