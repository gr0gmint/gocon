package main

//import "termbox"
//import "sdl"
import "net"
import "goprotobuf.googlecode.com/hg/proto"
import pwan "./pwan"
import . "gocon"
import "fmt"
import "time"

type SimpleHandler struct {
    ProtoHandler
}

func main() {

   raddr, _ := net.ResolveTCPAddr("127.0.0.1:7777")
   conn,_ := net.DialTCP("tcp", nil, raddr)
   proxy := NewProtoProxy (conn)
   
   joinmsg := pwan.NewClientJoin()
   joinmsg.Playername = proto.String("weoijoijasdjf")
   data,_ := proto.Marshal(joinmsg)
   fmt.Printf("%s", data)
   proxy.Send(data, "dafsdfsf")
   time.Sleep(10000000)
}
