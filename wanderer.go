package main

import . "gocon"
import "rand"
import "time"
import "net"
//import "os"

/* TODO:
     save to db */

const (
 WORLD_SIZE_X = 1024 //For better use in CenterTree
 WORLD_SIZE_Y = 1024
)


type WorldHandler struct {
    HotRoutine
    World *World
    hotlock bool
}

type Player struct {
    Name string
    PRoutine *Routine
}

type Coord struct {
    X,Y int
    Data map[string]interface{}
}

type World struct {
    Coords [WORLD_SIZE_X][WORLD_SIZE_Y] Coord
    PlayerIndex map[*Player]*Coord
}


type Listener struct {
    Routine   
}


var BServer = NewBroadcastServer()



func (c *Coord) Players() map[string]*Player {
    return c.Data["players"].(map[string]*Player)
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
    if coord,ok := w.PlayerIndex[player]; ok {  
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
    return &w.Coords[x][y]
}


func (this *WorldHandler) Main()  {
    this.Name =  "worldhandler"
    this.Register()
    this.Init()
    this.World = NewWorld()
    go this.HotStart()
    
}
const (
    DIRECTION_UP = iota
    DIRECTION_DOWN
    DIRECTION_LEFT
    DIRECTION_RIGHT
)
func (this *WorldHandler) PlayerPlace(player *Player, coord *Coord) bool {
    h := NewHot(func(data map[string]interface{}) {
        self := data["self"].(*GenericHot)
        m := NewMessage()
        if this.World.PlacePlayer(player,coord) {
            m.Key = "accepted"
        } else {
            m.Key = "declined"
        }
        self.Answer<-m
    })
    answer := <-h.Answer
    if answer.Key == "declined" {return false}
    return true
}
func (this *WorldHandler) PlayerMove(player *Player, direction int) bool {
    h:=NewHot(func(data map[string]interface{}){
            self := data["self"].(*GenericHot)
            var dirx,diry = 0,0
            switch {
                case direction == DIRECTION_UP:
                    diry += 1
                case direction == DIRECTION_DOWN:
                    diry -= 1
                case direction == DIRECTION_LEFT:
                    dirx -= 1
                case direction == DIRECTION_RIGHT:
                    dirx += 1
            }
            
            currentcord := this.World.PlayerIndex[player]

            m:=NewMessage()
            if currentcord.X+dirx > WORLD_SIZE_X || currentcord.X+dirx < 0 {
                m.Key = "declined"
            } else if currentcord.Y + diry > WORLD_SIZE_Y || currentcord.Y+diry < 0 {
                m.Key = "declined"
            } else {    
                newcoord := this.World.GetCoord(currentcord.X+dirx, currentcord.Y+diry)
                if this.World.PlacePlayer(player, newcoord) {
                    m.Key = "accepted"
                } else {
                    m.Key = "declined"
                }
            }
            self.Answer<-m
    })

        this.QueryHot(h)

    m := <-h.Answer
    if m.Key == "accepted" { return true; }
    return false
    
}


func (r *Listener) Main() { 
    //var err *os.Error
    laddr, _ := net.ResolveTCPAddr("0.0.0.0:7777")
    listener, _ := net.ListenTCP("tcp", laddr)
    for {
        conn,err := listener.AcceptTCP()
        if err ==nil {
        
            proxy := NewProtoProxy(conn)
            inithandler := NewInitProtoHandler(proxy)

            proxy.SetDefault(inithandler)
            go proxy.Main()
        }
    }   
}


func main() {
   rand.Seed(time.Nanoseconds())
    worldhandler := new(WorldHandler)
    go worldhandler.Main()

   server := new(Listener)
   //server.Init()
   server.Main()  
}
