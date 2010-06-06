package main

//import "termbox"
//import "sdl"
import "fmt"
import "net"

func main() {
   conn,_ := net.Dial("tcp", "", "127.0.0.1:7777")
   
}
