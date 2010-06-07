package main


import . "container/vector"

type Interval struct {
    Start int
    Stop int
    Sibling *Interval
    Data map[string]interface{}
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
           if i ,ok := <-c; ok {
                interval := i.(*Interval)
                if interval.Start <=  point && interval.Stop >= point {
                    acc_chan <- interval
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
            node.Intervals.Pop()
            c.Tree.NodeMap[interval] = nil
            if node.Intervals.Len() == 0 && node.Left == nil && node.Right == nil { //remove references to node
                if node == c.Tree.Top {break;}
                if node == node.Parent.Left {
                    node.Parent.Left = nil
                } else {
                    node.Parent.Right = nil
                }

            }
            break
        }
    }
}

func (c *CenterNode) AddInterval(interval *Interval) {
    if interval.Start <= c.Point && interval.Stop >= c.Point {
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
func NewCenterTree(start,stop int) *CenterTree {
    btree := new(CenterTree)
    btree.Start = start
    btree.Stop = stop
    btree.Top= new(CenterNode)
    btree.Top.Point = (stop+start)/2 
    btree.NodeMap = make(map[*Interval]*CenterNode)
    
    return btree
}

type Filter interface {
    ParseBroadcast(b *Broadcast)
}

type FilterPlayer struct {
    PlFilterMap map[*Player]*Filter
}
func NewFilterPlayer() *FilterPlayer {
    f := new(FilterPlayer)
    f.PlFilterMap = make(map[*Player]*Filter)
    return f
}
func (f *FilterPlayer) SetFilter(player *Player, filter *Filter) {
    m := f.PlFilterMap
    m[player] = filter
    
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
    f.XTree = NewCenterTree(0, WORLD_SIZE_X)
    return f

}
func FindOverlapping(v *Vector, coord *Coord) *Vector { 
    c := v.Iter()
    overlapping := new(Vector)
    for {
        i, ok := <-c
        if !ok { break; }
        in := i.(*Interval)
        if in.Sibling.Start <= coord.X && in.Sibling.Stop >= coord.Y {
            overlapping.Push(in)
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
        i, ok := <-c
        if !ok { break; }
        in := i.(*Interval)
        filter := in.Data["filter"].(*Filter)
        go filter.ParseBroadcast(b)
    }
    
}
func (f *FilterDistanceFromPlayer) AddFilter(startX,stopX,startY,stopY int, filter *Filter) {
    intervalX := new(Interval)
    intervalY := new(Interval)
    intervalX.Sibling = intervalY
    intervalX.Start = startX
    intervalX.Stop = stopX
    intervalY.Start = startY
    intervalY.Stop = stopY
    
    intervalX.Data["filter"] = filter
    f.XTree.AddInterval(intervalX)  
}
