package main


import "math"

type Interval struct {
    Śtart, Stop int
    Data interface{}
}

//We use this for interval searching, for the FilterDistanceFromPlayer

type CenterNode struct {
    Tree *CenterTree
    Point int
    Parent *CenterNode
    Left *CenterNode
    Right *CenterNode
    Intervals []Interval
    MaxIntervalStart, MaxIntervalStop int //Max interval widespan colliding with this centerline. Extra "optimization"
}
type CenterTree struct {
    Start,Stop int
    Top *CenterNode
}

func (c *CenterNode) searchNodeForPoint(point int, acc_chan chan *Interval) { 
    if point >= c.MaxIntervalStart && point <= c.MaxIntervalStop {
        for i, interval := range c.Intervals {
            if interval.Start <=  point && interval.Stop >= point {
                acc_chan <- interval
            }
        }
    } else { //search the left and right nodes
        if point <= c.Point && c.Left != nil  {
            go c.Left.searchNodeForPoint(point, acc_chan)
            return //So we don't hit the acc_chan<-nil
        } 
        if point >= c.Point && c.Right != nil {
            go c.Right.searchNodeForPoint(point, acc_chan)
            return //So we don't hit the acc_chan<-nil
        }
    }
    acc_chan <- nil
    
}
func (c *CenterNode) AddInterval(interval *Interval) {
    if start <= c.Point && stop >= c.Point {
        if c.Intervals == nil {
            c.Intervals = make(Interval, 100)[0:1]
            c.Intervals[0] = interval
        } else {
            l := len(c.Intervals)
            c.Intervals = c.Intervals[0:l+1]
        }
        return
    } else if interval.Stop < c.Point {
        if c.Left == nil { //Create left node if it does not exist
            c.Left = new(CenterNode)
            c.Left.Parent=c
            c.Left.Tree = c.Tree
            c.Left.Point = (c.Tree.Start+c.Point)/2
        }
        c.Left.AddInterval(interval)
    } else if interval.Start > c.Point {
        if c.Right == nil { //Create left node if it does not exist
            c.Right = new(CenterNode)
            c.Right.Parent=c
            c.Right.Tree = c.Tree
            c.Right.Point = (c.Tree.Stop+c.Point)/2
        }
        c.Right.AddInterval(interval)
    }
    
}
func (c *CenterTree) FindIntervals(point int) []Interval { //The chan is for asynchronous searching
    acc_chan := make(chan *Interval, 100)
    intervals := make(*Interval, 100)
    go c.Top.searchNodeForPoint(point, acc_chan)
    num_intervals := 0
    for i := 0; i <100; i++ {
        intervals[i] = <- acc_chan
        if intervals[i] == nil {
            break
        }
    }
    return intervals
}
func (c *CenterTree) AddInterval(interval *Interval) {
    c.Top.AddInterval(interval)
}
func NewCenterTree(start,stop int) *BinaryTree {
    btree := new(CenterTree)
    btree.Start = start
    btree.Stop = stop
    btree.Top= new(CenterNode)
    btree.Top.Point = (stop+start)/2 
    
    
    order := (int)(math.Log10(base)/math.Log10(0.5))

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

type distance struct {
    X,Y int
}
type FilterDistanceFromPlayer struct { 
    //Some time<->memory tradeoff could be needed here. Solved with Interval tree
    XTree *CenterTree
    YTree *CenterTree
}
func NewFilterDistanceFromPlayer() *FilterDistanceFromPlayer {
    f := new(FilterDistanceFromPlayer)
    f.XTree = NewCenterTree(0, WORLDSIZE_X)
    f.YTree = NewCenterTree(0,WORLDSIZE_Y)
    return f
}
func (f *FilterDistanceFromPlayer) ParseBroadcast(b *Broadcast) {
    coord := b.Data["coord"].(*Coord)
    
}
func (f *FilterDistanceFromPlayer) AddFilter(f *Filter) {
    
}
