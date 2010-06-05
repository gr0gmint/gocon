package main


import "math"
import . "container/vector"

type Interval struct {
    Śtart, Stop int
    Sibling *Interval
    Data interface{}
}




type CenterNode struct {
    Tree *CenterTree
    Point int
    Parent *CenterNode
    Left *CenterNode
    Right *CenterNode
    Intervals *Vector
    MaxIntervalStart, MaxIntervalStop int
}
type CenterTree struct {
    Start,Stop int
    Top *CenterNode
    NodeMap map[*Interval]*CenterNode
}

func (c *CenterNode) searchNodeForPoint(point int, acc_chan chan *Interval) { 
    if point >= c.MaxIntervalStart && point <= c.MaxIntervalStop {
        c := c.Intervals.Iter()
        for {
           if i.(*Interval) ,ok := <-c; ok {
                if i.Start <=  point && i.Stop >= point {
                    acc_chan <- i
                }
           } else {  break }
            
        }

    } else { //search the left and right nodes
        if point <= c.Point && c.Left != nil  {
            go c.Left.searchNodeForPoint(point, acc_chan)
            return 
        } 
        if point >= c.Point && c.Right != nil {
            go c.Right.searchNodeForPoint(point, acc_chan)
            return
        }
    }
    acc_chan <- nil
    
}
func (c *CenterNode) RemoveInterval(interval *Interval) {
    node := c.Tree.NodeMap[interval]
    length := node.Intervals.Len()
    for i := 0; i < length; i++ {   
        if node.Intervals.At(i).(*Interval) == interval {
            node.Intervals.Swap(i,length-1)
            node.Pop()
            break
        }
    }
}

func (c *CenterNode) AddInterval(interval *Interval) {
    if start <= c.Point && stop >= c.Point {
        if c.Intervals == nil {
            c.Intervals = new(Vector)
        }  
        c.Intervals.Push(interval)
        c.Tree.NodeMap[interval] = c
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
        if c.Right == nil { 
            c.Right = new(CenterNode)
            c.Right.Parent=c
            c.Right.Tree = c.Tree
            c.Right.Point = (c.Tree.Stop+c.Point)/2
        }
        c.Right.AddInterval(interval)
    }
    
}
func (c *CenterTree) FindIntervals(point int) *Vector /* *Interval */ { 
    acc_chan := make(chan *Interval, 100)
    intervals := new(Vector)
    go c.Top.searchNodeForPoint(point, acc_chan)
    for {
        i := <-acc_chan
        if i == nil {
            break
        }
        intervals.Push(i)
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
    btree.NodeMap = make(map[*Interval]*CenterNode)
    
    
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
        m[player] = m[player][0:len(m)+1] 
    }
    m[player][len(m[player])] = filter
    
}
func (f *FilterPlayer) ParseBroadcast(b *Broadcast) {
    player := b.Data["player"].(*Player)
    if filter, ok := f.PlFilterMap[player]; ok {
        filter.ParseBroadcast(b)
    }
}

type FilterDistanceFromPlayer struct { 

    XTree *CenterTree
}
func NewFilterDistanceFromPlayer() *FilterDistanceFromPlayer {
    f := new(FilterDistanceFromPlayer)
    f.XTree = NewCenterTree(0, WORLDSIZE_X)
    return f
}

}
func FindOverlapping(v *Vector, coord *Coord) *Vector { 
    c := v.Iter()
    overlapping := new(Vector)
    for {
        if i.(*Interval), ok := <-c; !ok {
            break
        }
        if i.Sibling.Start <= coord.X && i.Sibling.Stop >= coord.Y {
            overlapping.Push(i)
        }
    }
    return overlapping
}

func (f *FilterDistanceFromPlayer) ParseBroadcast(b *Broadcast) {
    coord := b.Data["coord"].(*Coord)
    v_intervalsx := f.XTree.FindIntervals(coord.X)
    overlapping := FindOverlapping(v_intervalsx, coord)   
    c := overlapping.Iter()
    for {
        if i.(*Interval), ok := <-c; !ok { break; }
        filter := i.Data["filter"].(*Filter)
        go filter.ParseBroadcast(b)
    }
    
}
func (f *FilterDistanceFromPlayer) AddFilter(startX,stopX,startY,stopY int, f *Filter) {
    intervalX := new(Interval)
    intervalY := new(Interval)
    intervalX.Sibling = intervalY
    intervalX.Start = startX
    intervalX.Stop = stopX
    intervalY.Start = startY
    intervalY.Stop = stopY
    
    intervalX.Data["filter"] = f
    f.XTree.AddInterval(intervalX)  
}
