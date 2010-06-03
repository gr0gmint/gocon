package main

import "./gocon"
import "fmt"
import "time"
//KRIG spil  (kortspillet)






type KrigServer struct {
 *gocon.Routine
 connections [2]*gocon.Connection	
}


func (self *KrigServer) main () {
  self.waitForPlayers()
  
}

func (self *KrigServer) waitForPlayers() {

   var conns [2]*gocon.Connection
   for i:=0; i< 2; i++ {
   	conns[i] = self.Listen("krigserver")
   }
   self.connections = conns  
   
}



type KrigClient struct {
  *gocon.Routine
  connection *gocon.Connection
}





func (self *KrigClient) main() {
   var conn *gocon.Connection
   self.connection = self.Connect("krigserver")
   if self.connection == nil {
     return
     }
   fmt.Printf("%s",conn.Receive())
}





func main() {
   joe := new(KrigServer)
   go joe.main()
   time.Sleep(1000000000)
   client := new(KrigClient)
   client.main()
   
}
