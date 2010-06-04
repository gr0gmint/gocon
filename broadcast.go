package main


type Broadcast struct {
    Data map[string]interface{}
}
func NewBroadcast () {
    b := new(Broadcast)
    b.Data = make(map[string]interface{})
}

type BroadcastServer struct {
    Routine
    Filters []*Filter 
}
func (r *BroadcastServer) Main() {
    r.NumFilters = 0
    r.Filters = make(*Filter, 1000)[0:1]
    m := <- r.Chan 
    switch  {
        case m.MsgType == "broadcast":
            //Handle the broadcast
        case m.MsgType == "hot":
            shared := make(map[string]interface{})
            shared["b"] = r
            h := m.Data["hot"].(*Hot)
            h.Unpack(shared)
    }
}
type HotBroadcastServerAddFilter struct {
       Filter *Filter
}
type (*h HotBroadCastServerAddFilter) Unpack(shared map[string]interface{}) int {
    b := shared["b"].(*BroadcastServer)
    numfilters := len(b.Filters)
    b.Filters = b.Filters[0:numfilters+1]
    b.Filters[numfilters] = h.Filter
    return 0
}
type Filter interface {
    ParseBroadcast(b *Broadcast)
}

type FilterPlayer struct {
    PlFilterMap map[*Player]([]*Filter)
}
func NewFilterPlayer() *FilterPlayer struct {
    f := new(FilterPlayer)
    f.PlFilterMap = make(map[*Player]([]*Filter))
}
func (f *FilterPlayer) AddFilter(player *Player, filter *Filter) {
    m := &f.PlFilterMap
    if filterlist,ok := m[player]; !ok {
        filterlist := make(*Filter,1000)[0:1]
        m[player] = filterlist
    } else {
        m[player] = m[player][0:len(m)+1] //expand the slice by one
    }
    m[player][len(m[player])] = filter
    
}
func (f *FilterPlayer) ParseBroadcast(b *Broadcast) {
    player := b.Data["player"].(*Player)
    if filter, ok := f.PlFilterMap[player]; ok {
        filter.ParseBroadcast(b)
    }
}


type FilterDistance struct { 
    //Some serious time<->memory tradeoff is needed here
    //There is shared data among goroutines in this filter, but it should be safe.
    Sections [][][][]*Cords
}

func NewBroadcastServer() *BroadcastServer {
    b := new(BroadcastServer)
    b.Init()
    return b
}
type LogicalConstruct struct {
    
}
type FineFilter struct {
    
}
