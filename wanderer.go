package wanderer

import . "gocon"
import "rand"
import "time"

//import "goprotobuf.googlecode.com/hg/proto"



const (
 WORLD_SIZE_X = 100
 WORLD_SIZE_Y = 100
)



type Coord struct {
    X,Y int
    Data map[string]interface{}
}
func (c *Coord) Players() map[string]*Player {
    return c.Data["players"].(map[string]*Player)
}

type Filter struct {
    
}

type BroadcastServer struct {
    *Routine
    Filters []Filter
}
func NewBroadcastServer() *BroadcastServer {
    b := new(BroadcastServer)
    b.Filters = make([1024]Filter)
}

var BServer = NewBroadcastServer()



type EventQueue struct {
    *Routine
}

type GameObject struct {
    x,y int
    Key string
}

type World struct {
    Coords [WORLD_SIZE_X*WORLD_SIZE_Y] Coord
    PlayerIndex map[*Player]*Coord
}
func NewWorld () *World {
    world := new(World)

    //Make Coord objects
    for i,c := range world.Coords {
        c.Data = make(map[string]interface{})
        c.x = i % WORLD_SIZE_X
        c.y = i / WORLD_SIZE_Y
        //Common fields
        c.Data["players"] = make(map[string]*Player)
        c.Data["gobjects"] = make([]*GameObject)
    }

    //Make Player:>Coord index
    world.PlayerIndex = make(map[*Player]*Coord
    
    return world
} 

func (w *World) PlacePlayer(player *Player, c *Coord) ok bool {
    
    
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
func (w *World) GetCoord(x,y int) {
    return w.Coords[x+y*WORLD_SIZE_X]
}



//This is basically the main routine of the server
type WandererServer struct {
    *Routine
    World *World
}
type Player struct {
    Name string
    PRoutine *Routine
}


func (r *WandererServer) Main()  {
    r.Name =  "wandererserver"
    r.Register()
    r.World = NewWorld()

    //setup some stuff

    //main loop
    var m  *Message
    var rchan chan *Message //Reponse channel, if needed
    for {
        m,rchan = r.ReceiveMessage()
        msgtype := m.Key
        switch {
            case msgtype == "event":
                e := m.Data["event"].(*Event)
                commune := make(map[string]interface{})
                commune["rchan"] = rchan
                commune["world"] = r
                whatdo := e.Unpack(commune)
                if whatdo == E_DIRECT {
                    r.parseCommune(commune)
                }
            case ??            
            
        }    
    }
}
func (r *WandererServer) parseCommune(commune map[string]interface{}) {
        //Check commune for some variables we might want
    if log, ok := commune["log"].(string); ok { //Event want us to log something
        //LOG FUNCTIONS
    }
} 


func main() {
   rand.Seed(time.Nanoseconds())
}
