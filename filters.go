package main


import . "container/vector"
import "fmt"

type Interval struct {
    Start int
    Stop int
    Node *CenterNode
    Sibling *Interval
    CBack func(*Broadcast)
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
    Leap int //for easy calculation
    NextTree *CenterTree //For 2D search
    
}
type CenterTree struct {
    Start,Stop int
    Top *CenterNode
    //NodeMap map[*Interval]*CenterNode
}

type Filter interface {
    ParseBroadcast(b *Broadcast)
}

type FilterPlayer struct {
    PlFilterMap map[*Player]*Filter
}

type FilterDistanceFromPlayer struct { 
    XTree *CenterTree
    PlIntervalMap map[*Player]*Interval
}

func (c *CenterNode) searchNodeForPoint(point int, acc_chan chan *Interval) { 
    fmt.Printf("In CenterNode·seachNodeForPoint\n")
        if c.Intervals == nil { c.Intervals = new(Vector) }
        ch := c.Intervals.Iter()
        for {
           if i := <-ch; i != nil {
                interval := i.(*Interval)
                if interval.Start <=  point && interval.Stop >= point {
                    acc_chan <- interval
                }
           } else {  break }
            
        }

        if point <= c.Point && c.Left != nil  {
            go c.Left.searchNodeForPoint(point, acc_chan)
            return 
        }  
        if point >= c.Point && c.Right != nil {
            go c.Right.searchNodeForPoint(point, acc_chan)
            return
        }
    acc_chan <- nil
    
}
func (this *Interval) Detach() {
    this.Node.RemoveInterval(this)
}
func (c *CenterNode) RemoveInterval(interval *Interval) {
    //node := c.Tree.NodeMap[interval]
    node := interval.Node
    length := node.Intervals.Len()
    
    for i := 0; i < length; i++ {   
        if node.Intervals.At(i).(*Interval) == interval {
            node.Intervals.Swap(i,length-1)
            node.Intervals.Pop()
            //c.Tree.NodeMap[interval] = nil
            if node.Intervals.Len() == 0 && node.Left == nil && node.Right == nil { 
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
    fmt.Printf("CenterNode·AddInterval\n")
    if interval.Start <= c.Point && interval.Stop >= c.Point {
        if c.Intervals == nil {
            c.Intervals = new(Vector)
        }  
        interval.Node = c
        c.Intervals.Push(interval)
        //c.Tree.NodeMap[interval] = c
        return
    } else if interval.Stop < c.Point {
        if c.Left == nil { //Create left node if it does not exist
            c.Left = new(CenterNode)
            c.Left.Parent=c
            c.Left.Tree = c.Tree
            c.Left.Leap = c.Leap/2
            c.Left.Point = c.Point-(c.Left.Leap)
            fmt.Printf("c.left.Point = %d\n", c.Left.Point)
        }
        c.Left.AddInterval(interval)
    } else if interval.Start > c.Point {
        if c.Right == nil { 
            c.Right = new(CenterNode)
            c.Right.Parent=c
            c.Right.Tree = c.Tree
            c.Right.Leap= c.Leap/2
            c.Right.Point = c.Point + c.Right.Leap/2
            fmt.Printf("c.Right.Point = %d\n", c.Right.Point)
        }
        c.Right.AddInterval(interval)
    }
    
}
func (c *CenterTree) FindIntervals(point int) *Vector /* *Interval, *CenterNode */ { 
    fmt.Printf("CenterTree·FindIntervals\n")
    acc_chan := make(chan *Interval, 100)
    intervals := new(Vector)
    go c.Top.searchNodeForPoint(point, acc_chan)
    for {
        i := <-acc_chan
        if i == nil {
            break
        }
        fmt.Printf("Found an interval\n")
        intervals.Push(i)
    }
    return intervals
}
/*
func (c *CenterTree) FindIntervalsByInterval(interval *Interval) *Vector { 
    acc_chan := make(chan *Interval, 100)
    intervals := new(Vector)
    go c.Top.searchNodeForPoint(interval.Start, acc_chan)
    go c.Top.searchNodeForPoint(interval.Stop, acc_chan)
    for {
        i := <-acc_chan
        if i == nil {
            break
        }
        intervals.Push(i)
    }
    return intervals
}
*/
func (c *CenterTree) AddInterval(interval *Interval) {
    c.Top.AddInterval(interval)
}
func NewCenterTree(start,stop int) *CenterTree {
    btree := new(CenterTree)
    btree.Start = start
    btree.Stop = stop
    btree.Top= new(CenterNode)
    btree.Top.Point = (stop+start)/2 
    btree.Top.Tree = btree
    btree.Top.Leap = (stop+start)/2 
    //btree.NodeMap = make(map[*Interval]*CenterNode)
    
    return btree
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


func NewFilterDistanceFromPlayer() *FilterDistanceFromPlayer {
    f := new(FilterDistanceFromPlayer)
    f.XTree = NewCenterTree(0, WORLD_SIZE_X)
    f.PlIntervalMap  = make(map[*Player]*Interval)
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
    fmt.Printf("In FilterDistanceFromPlayer·ParseBroadcast\n")
    coord := b.Data["coord"].(*Coord)
    v_intervalsx := f.XTree.FindIntervals(coord.X)
    if v_intervalsx.Len() == 0 {return}
    c := v_intervalsx.Iter()
    for {
        i := <- c
        if i == nil {break}
        in := i.(*Interval)
        if !(in.Sibling.Start <= coord.Y && in.Sibling.Stop >= coord.Y) { continue; } 
        fmt.Printf("Calling a interval-callback\n")
        if in.CBack == nil {continue;}
        go in.CBack(b)
    }

    
}
func (f *FilterDistanceFromPlayer) RegisterInterval(startX,stopX,startY,stopY int) *Interval{
    fmt.Printf("In FilterDistanceFromPlayer·AddFilter\n")
    intervalX := new(Interval)
    intervalY := new(Interval)
    intervalX.Start = startX
    intervalX.Stop = stopX
    intervalX.Sibling = intervalY
    intervalY.Start = startY
    intervalY.Stop = stopY
    f.XTree.AddInterval(intervalX)
    return intervalX
    
}
