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
    Handle(*Header,[]byte)
}

func (this *ProtoProxy) Init() {
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
            if err != nil {
                return nil, err
            }
            fmt.Printf("read %d bytes\n", n)
            total += n
            fmt.Printf("total: %d\n",total)
            if err !=nil { return nil,err }
        }
        
    return this.Buffer[0:total], nil
}

func (this *ProtoProxy) readMsg() (*Header,[]byte, os.Error) {
    //Read header first
    var hdrlen uint32
    var datalen uint32
    data,err := this.Read(this.Buffer[0:8])
    if err != nil  {
        return nil,nil,err
    }
    fmt.Printf("len(data) = %d\n", len(data))
    err = binary.Read(data[0:4], binary.BigEndian, &hdrlen)
        if err != nil { return nil,nil,err }
    err = binary.Read(data[4:8], binary.BigEndian, &datalen)
        if err != nil { return nil,nil,err }
            fmt.Printf("hdrlen=%d, datalen=%d\n", hdrlen, datalen)
    if !(hdrlen < 96 && datalen < 16000 ) {
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
    err = proto.Unmarshal(hdrdata, header)
    if err != nil {
        return nil,nil,err
    }
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
        header := NewHeader()
        header.Type = proto.Int32(t)
        header.Port = proto.Int32(port)
        hdrdata,err := proto.Marshal(header)
        if err != nil {
            fmt.Printf("%s\n", err)
            return
        }
        hdrlen := uint32(len(hdrdata))
        datalen :=  uint32(len(data))
        binary.Write(this.SBuffer, binary.BigEndian, [2]uint32{hdrlen,datalen})
        copy(this.SBuffer[8:8+len(hdrdata)],hdrdata)
        copy(this.SBuffer[8+len(hdrdata):8+len(hdrdata)+len(data)], data)
        //fmt.Printf("Writing this: [%d]%s\n", len(this.SBuffer[0:hdrlen+len(data)]), this.SBuffer[0:hdrlen+len(data)])
        this.Conn.Write(this.SBuffer[0:8+len(hdrdata)+len(data)])
    })
    this.QueryHot(h)

}


func (this *ProtoProxy) Main() {
    this.Init()
    for {
        header, data, err := this.readMsg(); 
        if err != nil {
            this.Conn.Close()
            fmt.Printf("%s\n", err)
            return
        } else {
                            fmt.Printf("header.Port = %d\n", *header.Port)
            if *header.Port == 0 {

                this.Default.Handle(header, data)
                
            } else {
                this.Handlers[int(*header.Port)].Handle(header,data)
            }
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



