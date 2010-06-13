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
    headersize int 
    PortChans map[int]chan hdr_n_data
}
type ProtoFilter interface {
     Check(*Header) bool
     Call ( *Header, Buf)
}

type ProtoHandler struct {
    HotRoutine
    Proxy *ProtoProxy
    
}
type hdr_n_data struct {
    header *Header
    data Buf
}

func (this *ProtoProxy) Init() {
    this.Buffer = make([]byte, 10000)
    this.SBuffer = make([]byte, 10000)
    
    go this.HotStart()
}

func NewProtoProxy(conn *net.TCPConn) *ProtoProxy {
    p := new(ProtoProxy)
    p.Conn = conn
    p.Init()
    p.PortChans = make(map[int]chan hdr_n_data)
    go p.HotStart()
    return p
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

func SubMsg(data []byte, t int32, encap bool) []byte {
        header := NewSubHeader()
        header.Type = proto.Int32(t)
        header.Encap = proto.Bool(encap)
        hdrdata,_ := proto.Marhal(header)
        hdrlen := uint32(len(hdrdata))
        datalen :=  uint32(len(data))
        buf := make([]byte, 8+hdrlen+datalen)
        binary.Write(this.SBuffer, binary.BigEndian, [2]uint32{hdrlen,datalen})
        copy(buf[8:8+len(hdrdata)],hdrdata)
        copy(buf[8+len(hdrdata):-1], data)
        return buf

}
func UnSubMsg(data []byte) (*SubHeader, []byte) {
    var hdrlen uint32
    var datalen uint32
    hdr,err := data[0:8]
    if err != nil  {
        return nil,err
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

func (this *ProtoProxy) SendMsg(data []byte, port int32, t int32, encap bool) {
    h := NewHot(func(shared map[string]interface{}){
        //self := shared["self"].(*GenericHot)
        header := NewHeader()
        header.Type = proto.Int32(t)
        header.Port = proto.Int32(port)
        header.Encap = proto.Bool(encap)
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

func (this *ProtoProxy) ReadMsg(port int) (*Header, []byte) {
    if this.PortChans[port] == nil {
        this.PortChans[port] = make(chan hdr_n_data)
    }
    m := <- this.PortChans[port]
    return m.header, m.data
}

func (this *ProtoProxy) Main() {
    this.Init()
    for {
        header, data, err := this.readMsg(); 
        if err != nil {
            this.Conn.Close()
            return
        } else {
            port := int(*header.Port)
            if this.PortChans[port] == nil {
                this.PortChans[port] = make(chan hdr_n_data)
            }
            m := hdr_n_data{header,data}
            go func() {this.PortChans[port] <- m }()
        }
        
    }
}

func (this *ProtoHandler) IsAccepted(port int) bool {
    _, data := this.Proxy.ReadMsg(port)
    a := NewAcceptBool()
    err := proto.Unmarshal(data,a)
    if err != nil {return false}
    if *a.Accept == true { return true }
    return false 
}

func (this *ProtoHandler) Acceptbool() {
        msg := NewAcceptBool()
        msg.Accept = proto.Bool(true)
        data,_ := proto.Marshal(msg)
        this.Proxy.SendMsg(data, 0,0,false)
}
func (this *ProtoHandler) Declinebool() {
        msg := NewAcceptBool()
        msg.Accept = proto.Bool(false)
        data,_ := proto.Marshal(msg)
        this.Proxy.SendMsg(data,0,0)
}
