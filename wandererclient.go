package main

//import "termbox"
//import "sdl"
//import "net"
import "goprotobuf.googlecode.com/hg/proto"
import pwan "./pwan"
import . "gocon"
import "fmt"
import "time"
import . "curses"

const (
    E_FATAL = iota
    E_NOTICE
)

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
    } else {

    }
   error:
    fmt.Printf("E: %s", err)

}

func intro_animation() {
    
}


func main() {
    Init()
    defer Shutdown()
intro_animation()
       time.Sleep(1000000000)

    /*
    raddr, _ := net.ResolveTCPAddr("127.0.0.1:7777")
    conn,err := net.DialTCP("tcp", nil, raddr)
    if err != nil { //goto error 
    }
    proxy := NewProtoProxy (conn)
    proxy.Init()
    go proxy.Main()
 
    handler := NewClientHandler(proxy)
    handler.Main()
    
   return
   error:
    fmt.Printf("E: %s", err)
    */
}
