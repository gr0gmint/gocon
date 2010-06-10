package main

//import "termbox"
//import "sdl"
import "net"
import "goprotobuf.googlecode.com/hg/proto"
import pwan "./pwan"
import . "gocon"
import "fmt"
//import "termbox"
import "time"

type ObjectHandler struct {
    ProtoHandler
}

type ClientHandler struct {
    ProtoHandler
}
func NewClientHandler(p *ProtoProxy) *ClientHandler {
    c := new(ClientHandler)
    c.Proxy=p
    c.Init()
    return c
}
func (this *ClientHandler) Main() {
    joinmsg := pwan.NewClientJoin()
    joinmsg.Playername = proto.String("Johnny")
    data,err := proto.Marshal(joinmsg)
        if err != nil { goto error }
        
    //Login
    this.Proxy.SendMsg(data,0,0)
       if (this.IsAccepted(0)) {
        fmt.Printf("IT HAS BEEN ACCEPTED HURR DURR\n")
    } else {
        fmt.Printf("NOt know what do")
    }
   error:
    fmt.Printf("E: %s", err)

}

func main() {
    raddr, _ := net.ResolveTCPAddr("127.0.0.1:7777")
    conn,err := net.DialTCP("tcp", nil, raddr)
    if err != nil { goto error }
    proxy := NewProtoProxy (conn)
    proxy.Init()
    go proxy.Main()
    time.Sleep(10000)
    handler := NewClientHandler(proxy)
    handler.Main()
    
   return
   error:
    fmt.Printf("E: %s", err)
}
