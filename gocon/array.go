package gocon

type Array struct {
    A []interface{}
}
func (a *Array) Swap(index_a, index_b int) {
    m := a.A[index_a]
    a.A[index_a] = a.A[index_b]
    a.A[index_b] = m
}
func (a *Array) DeleteIndex(index int) {
    a.Swap(index, len(a.A)-1)
    a.A = a.A[0:len(a.A)-1]
}

func copyElements(dst,src []interface{}) {
    for i,e := range src {
        dst[i] = e
    }
}

func (a *Array) Append(i interface{}) {
    if len(a.A) == cap(a.A) {
        //Make array larger
        oldcap := cap(a.A)
        new_A := make([]interface{}, cap(a.A)*2)
        copyElements(new_A, a.A)
        a.A = new_A[0:oldcap]
    }
    a.A = a.A[0:len(a.A)+1]
    a.A[len(a.A)-1] = i  
}

func (a *Array) Len() int {
    return len(a.A)
}



func NewArray() *Array {
    a := new(Array)
    a.A = make([]interface{}, 16)
    a.A = a.A[0:0]
    return a
}
