package main

import "goprotobuf.googlecode.com/hg/proto"


type ProtoHandler struct {
    Routine
    Conn *net.Conn
}
func NewProtoHandler(c *Conn) *ProtoHandler {
    p := new(ProtoHandler)
    p.Conn = c
    return p
}
func (p *ProtoHandler) Main() {
    buffer := make([]byte, 100000)

    headersize, err := p.Conn.Read(buffer)
    if err { return }
    
        


