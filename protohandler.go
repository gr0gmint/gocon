package main

import "os"
import "goprotobuf.googlecode.com/hg/proto"
import "./pwan"

type ProtoHandler struct {
    Routine
    Conn *net.Conn
    Player *Player
    Buffer []byte
    headersize int
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

func (this *ProtoHandler) Main() {
    //We want to know how big a Header-marshalized buffer is:
    temphdr := pwan.NewHeader()
    temphdr.Size = 0xf00
    data, _ := proto.Marshal(temphdr)
    this.headersize = len(data)
    
    p.Buffer := make([]byte, 100000)
    
    joinmsg := pwan.NewClientJoin()
    if header, _, err := this.RecvMsg(pwan.New()); err {
        this.Conn.Close()
        return
    } else {
        this.Player = NewPlayer()
        this.Player.Name = joinmsg.Playername
        
    }
    
}   
        


