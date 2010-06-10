package main

//import "termbox"
//import "sdl"
//import "net"
import "goprotobuf.googlecode.com/hg/proto"
import pwan "./pwan"
import . "gocon"
import "fmt"
import . "termbox"
import "time"

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

func drawString(x uint, y uint, s string, fg uint16, bg uint16) {
	for _, r := range s {
		ChangeCell(x, y, r, fg, bg)
		x++
	}
}

func drawBox() {
    width := Width()
    height := Height()
    var  i uint 
    for i=0; i <width; i++ {
        ChangeCell(i,0,' ', BLACK,YELLOW)
        ChangeCell(i,height-1, ' ',BLACK,YELLOW)
    }
    for i=0; i <height; i++ {
        ChangeCell(0,i,' ', BLACK,YELLOW)
        ChangeCell(width-1,i,  ' ',BLACK,YELLOW)
    }
}

func intro_animation() {
    drawString(1,1,"这个男人在跑步！！！！！！！！ 呵呵呵呵", YELLOW,BLACK)
    drawBox()
    Present()
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
