package main

import . "gocon"
import "rand"
import "time"
import "net"
import "os"
import "reflect"




const (
 WORLD_SIZE_X = 1024 //For better use in CenterTree
 WORLD_SIZE_Y = 1024
)



type Coord struct {
    X,Y int
    Data map[string]interface{}
}

func (c *Coord) Players() map[string]*Player {
    return c.Data["players"].(map[string]*Player)
}





var BServer = NewBroadcastServer()


type World struct {
    Coords [WORLD_SIZE_X][WORLD_SIZE_Y] Coord
    PlayerIndex map[*Player]*Coord
}
func NewWorld () *World {
    world := new(World)

    //Make Coord objects
    for i := range world.Coords {
        for j,c := range world.Coords[i] {
            c.Data = make(map[string]interface{})
            c.X = i
            c.Y = j
            //Common fields
            c.Data["players"] = make(map[string]*Player)
            //c.Data["gobjects"] = make([]*GameObject, 100)
        }
    }

    //Make Player:>Coord index
    world.PlayerIndex = make(map[*Player]*Coord)
    
    return world
} 

func (w *World) PlacePlayer(player *Player, c *Coord)  bool {
    
    
    if coord,ok := w.PlayerIndex[player]; ok {  //Delete current entry of player in world
        if _,ok = coord.Players()[player.Name]; ok {
            coord.Players()[player.Name] = nil
        }
        w.PlayerIndex[player] = nil
    }
    c.Players()[player.Name] = player
    w.PlayerIndex[player] = c
    return true
}
func (w *World) GetCoord(x,y int) *Coord {
    return &w.Coords[x+y*WORLD_SIZE_X]
}

//This is basically the main routine of the server
type WorldHandler struct {
    Routine
    World *World
}

type Player struct {
    Name string
    PRoutine *Routine
}

func (r *WorldHandler) Main()  {
    r.Name =  "worldhandler"
    r.Init()
    r.World = NewWorld()

    //setup some stuff

    //main loop
    var m  *Message
    var rchan chan *Message //Reponse channel, if needed
    for {
        m,rchan = r.ReceiveMessage()
        msgtype := m.Key
        switch {
            case msgtype == "hot":
                e := m.Data["hotcode"].(*Hot)
                shared := make(map[string]interface{})
                shared["rchan"] = rchan
                shared["world"] = r
                /*whatdo := */ e.Unpack(shared)
            case msgtype == "rpc":
                //We use the awesome reflect package here
                rpc := m.Data["rpc"].(*RPC)                
        }    
    }
}

const (
    DIRECTION_UP = iota
    DIRECTION_DOWN
    DIRECTION_LEFT
    DIRECTION_RIGHT
)
type HotWorldPlayerMove struct {
    Player *Player
    direction int
}
func (e *HotWorldPlayerMove) Type() string { return "PlayerMovement" }

func (e *HotWorldPlayerMove) Unpack(shared map[string]interface{}) int {
    world := shared["world"].(*World)
    rchan := shared["rchan"].(chan *Message)
    var dirx,diry = 0,0
    switch {
        case e.direction == DIRECTION_UP:
            diry += 1
        case e.direction == DIRECTION_DOWN:
            diry -= 1
        case e.direction == DIRECTION_LEFT:
            dirx -= 1
        case e.direction == DIRECTION_RIGHT:
            dirx += 1
    }
    currentcord := world.PlayerIndex[e.Player]

    m:=NewMessage()
    if currentcord.X+dirx > WORLD_SIZE_X || currentcord.X+dirx < 0 {
        m.Key = "declined"
    } else if currentcord.Y + diry > WORLD_SIZE_Y || currentcord.Y+diry < 0 {
        m.Key = "declined"
    } else {    
        newcoord := world.GetCoord(currentcord.X+dirx, currentcord.Y+diry)
        if world.PlacePlayer(e.Player, newcoord) {
            m.Key = "accepted"
        } else {
            m.Key = "declined"
        }
    }
    rchan <- m
    return E_DIRECT
}


func (r *WorldHandler) parseShared(shared map[string]interface{}) {
    //if log, ok := shared["log"].(string); ok { //Event want us to log something
        //LOG FUNCTIONS
    //}
} 



type Server struct {
    Routine   
}

func (r *Server) Main() { 
    laddr := new (net.TCPAddr)
    laddr.IP = net.ResolveTCPAddr("0.0.0.0")
    laddr.Port = 7777
    listener := net.ListenTCP("tcp", laddr)
    for {
        conn,err := listener.AcceptTCP()
        if !err {
            NewProtoHandler(conn)
            go NewProtoHandler.Main()
        }
    }   
}


func main() {
   rand.Seed(time.Nanoseconds())
    worldhandler := new(WorldHandler)
    worldhandler.Register()
    go worldhandler.Main()

   server := new(Server)
   server.Init()
   server.Main()  
}
