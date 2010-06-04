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

func NewBroadcastServer() *BroadcastServer {
    b := new(BroadcastServer)
    b.Init()
    return b
}
