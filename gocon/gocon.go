package gocon

var GlobalRoutines = make(map[string] *Routine)
const (
  _ = 1 << iota
  MSG_NOTIF
  MSG_REQ
)
type Message struct {
    Key string
    Data map[string]interface{}
    flags uint64
    rchan chan *Message
}
func NewMessage() *Message {
    m := new(Message)
    m.Data = make (map[string]interface{})
    return m
}

type State struct {
    Name string
    Data map[string]string
}

type Routine struct {
    Name string
    State State
    Chan chan *Message
}
func (r *Routine) Init() {
    r.Chan = make(chan *Message)
}

func (r *Routine) Register () bool {
    GlobalRoutines[r.Name] = r
    return true
}

func (r *Routine) ReqMessage(m *Message) *Message {
    m.flags |= MSG_REQ
   
    r.Chan <- m
    response := <- m.rchan
    return response
}

func (r *Routine) NotifMessage(m *Message) {
    m.flags |= MSG_NOTIF
    r.Chan <- m
}

func (r *Routine) ReceiveMessage() (*Message, chan *Message) {
    var m *Message
    m = <-r.Chan
    if m.flags & MSG_REQ > 0 {
        return m,m.rchan
    }
    return m,nil
    
    
}

/*
func (r *Routine) Listen(resource string) *Connection  {
   listener := new(Listener)
   listener.ch = make(chan *Syn)
   listener.routine = r
   
   Resources[resource] = listener
   syn := <-listener.ch

   ack := new (Ack)
   ack.ch = make(chan interface{})
   
   syn.ch<- ack
   conn := new(Connection)
   conn.in = ack.ch
   conn.out = syn.ch
   conn.other = syn.origin
   return conn
}
func (r *Routine) Connect(resource string) *Connection {
  listener, ok := Resources[resource]
  if !ok {
     return nil
  }
  syn := new(Syn)
  syn.origin = r
  syn.ch = make(chan interface{})
  
  listener.ch <- syn
  ack := (<-syn.ch).(*Ack)

  Resources[resource] = nil

  conn := new(Connection)
  conn.out = ack.ch
  conn.in = syn.ch
  
  conn.other = listener.routine
  return conn
}

*/
