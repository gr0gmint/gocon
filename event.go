package main


const (
   E_DEFER = iota
   E_DIRECT
   
)
type Event interface {
    Type() string
    Description() string
    Unpack(map[string]interface{}) int
}

type EventPlayerJoin struct {
    Player *Player
}
func (e *EventPlayerJoin) Type() string { return "PlayerJoin" }
func (e *EventPlayerJoin) Description() string {    return "Player "+e.Player.Name+ " is trying to join the world" }
func (e *EventPlayerJoin) Unpack(commune map[string]interface{}) whatdo int {
    world := commune["world"].(*World)
    rchan := commune["rchan"].(chan *Message)
    randomcoord := world.GetCoord(rand.Int() % WORLD_SIZE_X, rand.Int() % WORLD_SIZE_Y)
   
    m := NewMessage()
    if world.PlacePlayer(e.Player, randomcoord) {

        m.Key = "accepted"
        
    } else {

        m.Key = "declined"
    }
    rchan <- m
    return E_DIRECT
}


const (
    DIRECTION_UP = iota
    DIRECTION_DOWN
    DIRECTION_LEFT
    DIRECTION_RIGHT
type EventWorldPlayerMove struct {
    Player *Player
    direction int
}
func (e *EventWorldPlayerMove) Type() { return "PlayerMovement" }
func (e *EventWorldPlayerMove) Description() { 
    var direction string
    switch {
        case e.direction == DIRECTION_UP:
            direction = "up"
        case e.direction == DIRECTION_DOWN:
            direction = "down"
        case e.direction == DIRECTION_LEFT:
            direction = "left"
        case e.direction == DIRECTION_RIGHT:
            direction = "right"
    }
    return "Player wants to move "+direction
func (e *EventWorldPlayerMove) Unpack(commune map[string]interface{}) int {
    world := commune["world"].(*World)
    rchan := commune["rchan"].(*World)
    var int dirx,diry = 0
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
    currentcord := world.PlayerIndex[e.Player.Name]

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
