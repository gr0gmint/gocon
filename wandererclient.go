package main

import "goprotobuf.googlecode.com/hg/proto"
import . "gocon"
import "fmt"
import "time"
import . "curses"
import "net"

const (
    MAN = "男"
    WOMAN =  "女"
)

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

type WorldWindow struct {
    HotRoutine
    Window
}
type TextWindow struct {
    HotRoutine
    Window
}

func NewClientHandler(p *ProtoProxy) *ClientHandler {
    c := new(ClientHandler)
    c.Proxy=p
    c.Init()
    return c
}


func (this *ClientHandler) Main() {
    joinmsg := NewClientJoin()
    joinmsg.Playername = proto.String("Johnny")
    data,err := proto.Marshal(joinmsg)
        if err != nil { fmt.Printf("E: %s", err)  }
    
    //Login
    this.Proxy.SendMsg(data,0,Client_JOIN)
       if (this.IsAccepted(0)) {
        fmt.Printf("Request was accepted by server\n")
        for {
            ch := Stdwin.Getch()
            direction := new( ClientWalk_Direction)
            switch {

                case ch == KEY_UP:
                    *direction = ClientWalk_DIRECTION_UP
                case ch == KEY_DOWN:
                    *direction = ClientWalk_DIRECTION_DOWN
                case ch == KEY_LEFT:
                    *direction = ClientWalk_DIRECTION_LEFT
                case ch == KEY_RIGHT:
                    *direction = ClientWalk_DIRECTION_RIGHT
                default:
                    return
            }
            m := NewClientWalk()
            m.Direction = direction
            mdata,err:= proto.Marshal(m); if err!= nil {Endwin(); fmt.Printf("Couldnt marshal\n")}
            this.Proxy.SendMsg(mdata,0,Client_WALK)
        }
        endfor:
    } else {

    }
   

}



func main() {
    Initscr()
    defer Endwin()
    Cbreak()
    Noecho()
    Stdwin.Keypad(true)
    
    raddr, _ := net.ResolveTCPAddr("127.0.0.1:7777")
    conn,err := net.DialTCP("tcp", nil, raddr)
    if err != nil { //goto error 
    }
    proxy := NewProtoProxy (conn)
    proxy.Init()
    go proxy.Main()
   time.Sleep(1000000000)
    handler := NewClientHandler(proxy)
    handler.Main()
    
   return
   error:
    fmt.Printf("E: %s", err)
    
}
