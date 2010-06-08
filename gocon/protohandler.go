package gocon

import "os"
import "goprotobuf.googlecode.com/hg/proto"
import "net"
import "fmt"
import binary "encoding/binary"

type Buf []byte
func (b Buf) Write(p []byte) (int, os.Error) {
    copy(b,p)
    return len(p), nil
}
func (b Buf) Read(p []byte) (int, os.Error) {
    copy(p,b)
    return len(b), nil
}

type ProtoProxy struct {
    HotRoutine
    Buffer Buf
    SBuffer Buf
    Conn *net.TCPConn
    Default *IProtoHandler
    Handlers map[int]IProtoHandler  
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
    temphdr.Size = proto.Int32(4831)
    temphdr.Type = proto.Int32(0)
    temphdr.Port = proto.Int32(0)
    data, err := proto.Marshal(temphdr)
    if err != nil {
        fmt.Printf("E: Couldn't marshal header\n")
    }
    this.headersize = len(data)
    this.Buffer = make([]byte, 10000)
    this.SBuffer = make([]byte, 10000)
    
    this.Handlers = make(map[int]IProtoHandler)
    go this.HotStart()
}

func NewProtoProxy(conn *net.TCPConn) *ProtoProxy {
    p := new(ProtoProxy)
    p.Conn = conn
    p.Init()
    go p.HotStart()
    return p
}

func (this *ProtoProxy) AddHandler(port int, handler IProtoHandler) {
    this.Handlers[port] = handler
}
func (this *ProtoProxy) SetDefault(def IProtoHandler) {
    this.Default = &def
}
func (this *ProtoProxy) RemoveHandler(port int) {
    this.Handlers[port] = nil
}

func (this *ProtoProxy) Read (buf Buf) (Buf, os.Error) {
    if buf == nil {
        n,err := this.Conn.Read(buf)
        return this.Buffer[0:n], err 
    } 
    var total int = 0
    buflen := len(buf)
    for ; total < buflen; {
            n, err := this.Conn.Read(this.Buffer[total:buflen])
            fmt.Printf("read %d bytes\n", n)
            total += n
            fmt.Printf("total: %d\n",total)
            if err !=nil { return nil,err }
        }
        
    return this.Buffer[0:total], nil
}

func (this *ProtoProxy) readMsg() (*Header,[]byte, os.Error) {
    //Read header first
    var hdrlen int32
    var datalen int32
    data,err := this.Read(this.Buffer[0:8])
    if err != nil  {
        return nil,nil,err
    }
    fmt.Printf("len(data) = %d\n", len(data))
    binary.Read(data[0:4], binary.BigEndian, &hdrlen)
    binary.Read(data[4:8], binary.BigEndian, &datalen)
    if !(hdrlen < 4096 && datalen < 16000 ) {
        fmt.Printf("hdrlen=%d, datalen=%d.. Far too high\n", hdrlen, datalen)
        return nil,nil,os.ENOMEM
    }
    hdrdata := make(Buf, hdrlen)
    newdata := make(Buf, datalen)
    tmp,err := this.Read(this.Buffer[0:hdrlen])
    if err != nil {
        //this.Conn.Close()
        return nil,nil,err
    }
    copy(hdrdata,tmp)
    header := NewHeader()
    proto.Unmarshal(hdrdata, header)
    tmp,err = this.Read(this.Buffer[0:datalen])
    if err != nil {
        return nil,nil,err
    }
    copy(newdata,tmp)
    
    
    return header, newdata,nil
} 
func (this *ProtoProxy) Send(data []byte, port int32, t int32) {
    h := NewHot(func(shared map[string]interface{}){
        fmt.Printf("inside hot\n")
        //self := shared["self"].(*GenericHot)
        var datalen int32 = int32(len(data))
        binary.Write(this.SBuffer, binary.BigEndian, &datalen)
        copy(this.SBuffer[4:4+len(data)], data)
        //fmt.Printf("Writing this: [%d]%s\n", len(this.SBuffer[0:hdrlen+len(data)]), this.SBuffer[0:hdrlen+len(data)])
        this.Conn.Write(this.SBuffer[0:4+len(data)])
    })
    this.QueryHot(h)

}


func (this *ProtoProxy) Main() {
    this.Init()
    for {
        header, data, err := this.readMsg(); 
        if err != nil {
            this.Conn.Close()
            fmt.Printf("\nConnection closed\n")
            return
        } else {
            fmt.Printf("%s %s\n",header,data)
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



