package main

import "goprotobuf.googlecode.com/hg/proto"
import . "gocon"
import "fmt"
import "time"
import . "termbox"
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
}
type TextWindow struct {
    HotRoutine
}

func NewObjectHandler(p *ProtoProxy) *ObjectHandler {
    c := new(ObjectHandler)
    c.Proxy = p
    c.Init()
    return c
}
func (this *ObjectHandler) Main() {
    for {
        header, data := this.Proxy.ReadMsg(PORT_PUSH)
        t := int(*header.Type)
        switch {
            case t == Server_UPDATELOCATION:
                m := NewUpdatePlayerCoord()
                err := proto.Unmarshal(data, m)
                if err != nil { return }
                DrawEmptyMap()
                
                Addstr(1,1,"%d %d     ",YELLOW,BLACK, *m.Coord.X,*m.Coord.Y)
                Present()
                //Stdwin.Addstr(1,1, "%d, %d\n", 0, *m.Coord.X, *m.Coord.Y)
                //Stdwin.Refresh()
            default:
                //Stdwin.Addstr(0,5, "Wtf man", 0)
        }
    }
}

func NewClientHandler(p *ProtoProxy) *ClientHandler {
    c := new(ClientHandler)
    c.Proxy=p
    c.Init()
    return c
}
func Box(startx,starty, width, height uint, fg, bg uint16,boxcharacters []int ) { //"─│┌┐└┘"
     
    ChangeCell(startx,starty, boxcharacters[2], fg,bg)
    ChangeCell(startx+width,starty, boxcharacters[3], fg,bg)
    ChangeCell(startx,starty+height, boxcharacters[4], fg,bg)
     ChangeCell(startx+width,starty+height, boxcharacters[5], fg,bg)
     for i := startx +1; i < startx+width; i++ {
        ChangeCell(i,starty,boxcharacters[0], fg,bg)
     }
          for i := startx +1; i < startx+width; i++ {
        ChangeCell(i,starty+height,boxcharacters[0], fg,bg)
     }
          for i := starty +1; i < startx+width; i++ {
        ChangeCell(startx,i,boxcharacters[1], fg,bg)
     }
               for i := starty +1; i < startx+width; i++ {
        ChangeCell(startx+width,i,boxcharacters[1], fg,bg)
     }
     
}
func StdBox(startx,starty, width, height uint, fg, bg uint16) {
    Box(startx,starty, width, height, fg, bg, &[...]int{'─', '│', '┌', '┐', '└', '┘'})
}
func DrawEmptyMap() {
    StdBox(2,2,10,10, YELLOW,BLACK)
}
func Addstr(x,y uint, s string, fg, bg uint16, v ... interface{}) {
    l := Width()
    i := x
    j := y
    si := 0
    ns := fmt.Sprintf(s, v)
    ls := len(ns)
    for si<ls {
        if (i >=l) { j += 1; i = 0; continue }
        ChangeCell(i,j,int(ns[si]),fg,bg)
        si += 1
        i += 1
    }

}
func ViewCoordinates(x,y int) {
    
}


func (this *ClientHandler) Main() {
    joinmsg := NewClientJoin()
    joinmsg.Playername = proto.String("Johnny")
    data,err := proto.Marshal(joinmsg)
        if err != nil { fmt.Printf("E: %s", err)  }
    
    //Login
    this.Proxy.SendMsg(data,0,Client_JOIN,false)
       if (this.IsAccepted(0)) {
        for {
            e := new(Event)
            PollEvent(e)
            ch := e.Key
            direction := new( ClientWalk_Direction)
            switch {

                case ch == KEY_ARROW_UP:
                    *direction = ClientWalk_DIRECTION_UP
                case ch == KEY_ARROW_DOWN:
                    *direction = ClientWalk_DIRECTION_DOWN
                case ch == KEY_ARROW_LEFT:
                    *direction = ClientWalk_DIRECTION_LEFT
                case ch == KEY_ARROW_RIGHT:
                    *direction = ClientWalk_DIRECTION_RIGHT
                default:
                    return
            }
            m := NewClientWalk()
            m.Direction = direction
            mdata,err:= proto.Marshal(m); if err!= nil { fmt.Printf("Couldnt marshal\n")}
            this.Proxy.SendMsg(mdata,0,Client_WALK,false)
        }
        endfor:
    } else {

    }
   

}



func main() {
    Init()
    
    defer Shutdown()
    //Cbreak()
    
    raddr, _ := net.ResolveTCPAddr("127.0.0.1:7777")
    conn,err := net.DialTCP("tcp", nil, raddr)
    if err != nil { //goto error 
    }
    proxy := NewProtoProxy (conn)
    proxy.Init()
    go proxy.Main()
   time.Sleep(100000000)
    handler := NewClientHandler(proxy)
    ohandler := NewObjectHandler(proxy)
    go ohandler.Main()
    handler.Main()
    
   return
   error:
    fmt.Printf("E: %s", err)
    
}
