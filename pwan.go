// Code generated by protoc-gen-go from "wanderer.proto"
// DO NOT EDIT!

package pwan

import "goprotobuf.googlecode.com/hg/proto"

type Server_Type int32
const (
	Server_ANSWERBOOL = 1
	Server_ANSWER = 2
	Server_TEXTMESSAGE = 3
	Server_LOCATION = 4
	Server_PLAYERMOVE = 5
	Server_OBJECTMOVE = 6
	Server_OBJECT = 7
	Server_INVENTORY = 8
	Server_ATTACK = 9
	Server_PLAYERATTR = 10
	Server_SOCIAL = 11
	Server_QUESTLOG = 12
)
var Server_Type_name = map[int32] string {
	1: "ANSWERBOOL",
	2: "ANSWER",
	3: "TEXTMESSAGE",
	4: "LOCATION",
	5: "PLAYERMOVE",
	6: "OBJECTMOVE",
	7: "OBJECT",
	8: "INVENTORY",
	9: "ATTACK",
	10: "PLAYERATTR",
	11: "SOCIAL",
	12: "QUESTLOG",
}
var Server_Type_value = map[string] int32 {
	"ANSWERBOOL": 1,
	"ANSWER": 2,
	"TEXTMESSAGE": 3,
	"LOCATION": 4,
	"PLAYERMOVE": 5,
	"OBJECTMOVE": 6,
	"OBJECT": 7,
	"INVENTORY": 8,
	"ATTACK": 9,
	"PLAYERATTR": 10,
	"SOCIAL": 11,
	"QUESTLOG": 12,
}
func NewServer_Type(x int32) *Server_Type {
	e := Server_Type(x)
	return &e
}

type Client_Type int32
const (
	Client_JOIN = 1
	Client_WALK = 2
	Client_INTERACT = 3
	Client_INVENTORY = 4
	Client_ATTACK = 5
	Client_PLAYERATTR = 6
	Client_QUESTLOG = 7
	Client_SOCIAL = 8
	Client_ANSWERBOOL = 9
	Client_ANSWER = 10
)
var Client_Type_name = map[int32] string {
	1: "JOIN",
	2: "WALK",
	3: "INTERACT",
	4: "INVENTORY",
	5: "ATTACK",
	6: "PLAYERATTR",
	7: "QUESTLOG",
	8: "SOCIAL",
	9: "ANSWERBOOL",
	10: "ANSWER",
}
var Client_Type_value = map[string] int32 {
	"JOIN": 1,
	"WALK": 2,
	"INTERACT": 3,
	"INVENTORY": 4,
	"ATTACK": 5,
	"PLAYERATTR": 6,
	"QUESTLOG": 7,
	"SOCIAL": 8,
	"ANSWERBOOL": 9,
	"ANSWER": 10,
}
func NewClient_Type(x int32) *Client_Type {
	e := Client_Type(x)
	return &e
}

type ClientWalk_Direction int32
const (
	ClientWalk_DIRECTION_UP = 1
	ClientWalk_DIRECTION_DOWN = 2
	ClientWalk_DIRECTION_LEFT = 3
	ClientWalk_DIRECTION_RIGHT = 4
)
var ClientWalk_Direction_name = map[int32] string {
	1: "DIRECTION_UP",
	2: "DIRECTION_DOWN",
	3: "DIRECTION_LEFT",
	4: "DIRECTION_RIGHT",
}
var ClientWalk_Direction_value = map[string] int32 {
	"DIRECTION_UP": 1,
	"DIRECTION_DOWN": 2,
	"DIRECTION_LEFT": 3,
	"DIRECTION_RIGHT": 4,
}
func NewClientWalk_Direction(x int32) *ClientWalk_Direction {
	e := ClientWalk_Direction(x)
	return &e
}

type Server struct {
	XXX_unrecognized	[]byte
}
func (this *Server) Reset() {
	*this = Server{}
}
func NewServer() *Server {
	return new(Server)
}

type Client struct {
	XXX_unrecognized	[]byte
}
func (this *Client) Reset() {
	*this = Client{}
}
func NewClient() *Client {
	return new(Client)
}

type Header struct {
	Size	*int32	"PB(varint,1,req,name=size)"
	Type	*int32	"PB(varint,2,req,name=type)"
	Handler	*string	"PB(bytes,3,opt,name=handler)"
	XXX_unrecognized	[]byte
}
func (this *Header) Reset() {
	*this = Header{}
}
func NewHeader() *Header {
	return new(Header)
}

type ClientJoin struct {
	Playername	*string	"PB(bytes,1,req,name=playername)"
	XXX_unrecognized	[]byte
}
func (this *ClientJoin) Reset() {
	*this = ClientJoin{}
}
func NewClientJoin() *ClientJoin {
	return new(ClientJoin)
}

type ClientWalk struct {
	Direction	*ClientWalk_Direction	"PB(varint,1,req,name=direction,enum=pwan.ClientWalk_Direction)"
	XXX_unrecognized	[]byte
}
func (this *ClientWalk) Reset() {
	*this = ClientWalk{}
}
func NewClientWalk() *ClientWalk {
	return new(ClientWalk)
}

func init() {
	proto.RegisterEnum("pwan.Server_Type", Server_Type_name, Server_Type_value)
	proto.RegisterEnum("pwan.Client_Type", Client_Type_name, Client_Type_value)
	proto.RegisterEnum("pwan.ClientWalk_Direction", ClientWalk_Direction_name, ClientWalk_Direction_value)
}