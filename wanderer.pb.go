// Code generated by protoc-gen-go from "wanderer.proto"
// DO NOT EDIT!

package main

import proto "net/proto2/go/proto"

// Reference proto import to suppress error if it's not otherwise used.
var _ = proto.GetString

type Server_Type int32
const (
	Server_ANSWERBOOL = 1
	Server_ANSWERJSON = 2
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
	Server_JSON = 13
	Server_OBJECTPUSH = 14
	Server_UPDATELOCATION = 15
)
var Server_Type_name = map[int32] string {
	1: "ANSWERBOOL",
	2: "ANSWERJSON",
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
	13: "JSON",
	14: "OBJECTPUSH",
	15: "UPDATELOCATION",
}
var Server_Type_value = map[string] int32 {
	"ANSWERBOOL": 1,
	"ANSWERJSON": 2,
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
	"JSON": 13,
	"OBJECTPUSH": 14,
	"UPDATELOCATION": 15,
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
	Client_ANSWERJSON = 10
	Client_JSON = 11
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
	10: "ANSWERJSON",
	11: "JSON",
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
	"ANSWERJSON": 10,
	"JSON": 11,
}
func NewClient_Type(x int32) *Client_Type {
	e := Client_Type(x)
	return &e
}

type GObject_Type int32
const (
	GObject_PLAYER = 1
	GObject_NPC = 2
	GObject_BEAST = 3
	GObject_WALL = 4
)
var GObject_Type_name = map[int32] string {
	1: "PLAYER",
	2: "NPC",
	3: "BEAST",
	4: "WALL",
}
var GObject_Type_value = map[string] int32 {
	"PLAYER": 1,
	"NPC": 2,
	"BEAST": 3,
	"WALL": 4,
}
func NewGObject_Type(x int32) *GObject_Type {
	e := GObject_Type(x)
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

type Coordinate struct {
	X	*int32	"PB(varint,1,req,name=x)"
	Y	*int32	"PB(varint,2,req,name=y)"
	XXX_unrecognized	[]byte
}
func (this *Coordinate) Reset() {
	*this = Coordinate{}
}
func NewCoordinate() *Coordinate {
	return new(Coordinate)
}

type GObject struct {
	Id	*int32	"PB(varint,1,req,name=id)"
	Char	*int32	"PB(varint,2,req,name=char)"
	Description	*string	"PB(bytes,3,req,name=description)"
	Type	*GObject_Type	"PB(varint,4,req,name=type,enum=main.GObject_Type)"
	Coord	*Coordinate	"PB(bytes,5,req,name=coord)"
	Color	*int32	"PB(varint,6,opt,name=color)"
	Disappear	*bool	"PB(varint,7,opt,name=disappear)"
	Player	*GObjectPlayer	"PB(bytes,50,opt,name=player)"
	Npc	*GObjectNPC	"PB(bytes,51,opt,name=npc)"
	XXX_unrecognized	[]byte
}
func (this *GObject) Reset() {
	*this = GObject{}
}
func NewGObject() *GObject {
	return new(GObject)
}

type UpdatePlayerCoord struct {
	Coord	*Coordinate	"PB(bytes,1,req,name=coord)"
	XXX_unrecognized	[]byte
}
func (this *UpdatePlayerCoord) Reset() {
	*this = UpdatePlayerCoord{}
}
func NewUpdatePlayerCoord() *UpdatePlayerCoord {
	return new(UpdatePlayerCoord)
}

type GObjectPlayer struct {
	Name	*string	"PB(bytes,1,req,name=name)"
	Idle	*bool	"PB(varint,2,req,name=idle)"
	Sex	*bool	"PB(varint,3,req,name=sex)"
	XXX_unrecognized	[]byte
}
func (this *GObjectPlayer) Reset() {
	*this = GObjectPlayer{}
}
func NewGObjectPlayer() *GObjectPlayer {
	return new(GObjectPlayer)
}

type GObjectNPC struct {
	Name	*string	"PB(bytes,1,req,name=name)"
	Function	*string	"PB(bytes,2,req,name=function)"
	XXX_unrecognized	[]byte
}
func (this *GObjectNPC) Reset() {
	*this = GObjectNPC{}
}
func NewGObjectNPC() *GObjectNPC {
	return new(GObjectNPC)
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
	Direction	*ClientWalk_Direction	"PB(varint,1,req,name=direction,enum=main.ClientWalk_Direction)"
	XXX_unrecognized	[]byte
}
func (this *ClientWalk) Reset() {
	*this = ClientWalk{}
}
func NewClientWalk() *ClientWalk {
	return new(ClientWalk)
}

func init() {
	proto.RegisterEnum("main.Server_Type", Server_Type_name, Server_Type_value)
	proto.RegisterEnum("main.Client_Type", Client_Type_name, Client_Type_value)
	proto.RegisterEnum("main.GObject_Type", GObject_Type_name, GObject_Type_value)
	proto.RegisterEnum("main.ClientWalk_Direction", ClientWalk_Direction_name, ClientWalk_Direction_value)
}
