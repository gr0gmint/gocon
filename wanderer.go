package main

import . "gocon"
import "rand"
import "time"
import "net"
import "fmt"
//import "os"

/* TODO:
     save to db */

const (
 WORLD_SIZE_X = 1024 //For better use in CenterTree
 WORLD_SIZE_Y = 1024
)

const (
 PORT_INIT = 0
)


type WorldHandler struct {
    HotRoutine
    World *World
    hotlock bool
}

type Player struct {
    Name string
}

type Coord struct {
    X,Y int
    Data map[string]interface{}
}

type World struct {
    Coords map[int](map[int]*Coord)
    PlayerIndex map[*Player]*Coord
}


type Listener struct {
    Routine   
}


var BServer = NewBroadcastServer()



func (c *Coord) GetPlayer(name string) *Player {
    
    players := c.Data["players"].(map[string]*Player)
    return players[name]
}

func (c *Coord) SetPlayer(name string, value *Player) {
    players := c.Data["players"].(map[string]*Player)
    players[name] = value
    return
}


func NewWorld () *World {
    world := new(World)
    world.Coords = make(map[int](map[int]*Coord))
    world.PlayerIndex = make(map[*Player]*Coord)
    return world
} 

func (w *World) PlacePlayer(player *Player, c *Coord)  bool {
        c.SetPlayer(player.Name, nil)
        w.PlayerIndex[player] = nil
    c.SetPlayer(player.Name, player)
    w.PlayerIndex[player] = c
    return true
}
func (w *World) GetCoord(x,y int) *Coord {
    if !(x >= 0 && x <= WORLD_SIZE_X && y >= 0 && y <= WORLD_SIZE_Y) {
        return nil
    }
    _,ok := w.Coords[x]
    if !ok  { w.Coords[x] = make(map[int]*Coord) }
    _,ok = w.Coords[x][y]
    if !ok {
        fmt.Printf("New *Coord created\n")
        c := new(Coord)
        c.Data = make(map[string]interface{})
        c.X = x
        c.Y = y
        //Common fields
        c.Data["players"] = make(map[string]*Player)
        w.Coords[x][y] = c
    }
    return w.Coords[x][y]
}


func (this *WorldHandler) Main()  {
    this.Name =  "worldhandler"
    GlobalRoutines["worldhandler"] = this
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
func (this *WorldHandler) PlacePlayer(player *Player, coord *Coord) bool {
    h := NewHot(func(data map[string]interface{}) {
        self := data["self"].(*GenericHot)
        m := NewMessage()
        if this.World.PlacePlayer(player,coord) {
            fmt.Printf("Trying to broadcast now\n")
            b := NewBroadcast()
            b.Data["player"] = player
            b.Data["coord"] = coord
            BServer.Broadcast(b)
            
            m.Key = "accepted"
        } else {
            m.Key = "declined"
        }
        self.Answer<-m
    })
    this.QueryHot(h)
    answer := <-h.Answer
    if answer.Key == "declined" {return false}
    return true
}

func (this *WorldHandler) GetCurrentCoord (player *Player )*Coord  {
    return this.World.PlayerIndex[player]
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
                if this.PlacePlayer(player, newcoord) {
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
    laddr, err := net.ResolveTCPAddr("0.0.0.0:7777")
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    listener, err := net.ListenTCP("tcp", laddr)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    for {
        conn,err := listener.AcceptTCP()
        if err ==nil {
            proxy := NewProtoProxy(conn)
            inithandler := NewInitProtoHandler(proxy)
            go proxy.Main()
            go inithandler.Main()
        } else {
            fmt.Printf("%s\n", err)
            return
        }
    }   
}


func main() {
   rand.Seed(time.Nanoseconds())
    worldhandler := new(WorldHandler)
    go worldhandler.Main()
    BServer.Setup()
go BServer.Main()
   server := new(Listener)
   
   server.Main()  
}
