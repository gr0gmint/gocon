// Code generated by protoc-gen-go from "wanderer.proto"
// DO NOT EDIT!

package wanderer

import "goprotobuf.googlecode.com/hg/proto"

type InitialGroup_MsgType int32
const (
	InitialGroup_JOIN = 1
	InitialGroup_STAT = 2
)
var InitialGroup_MsgType_name = map[int32] string {
	1: "JOIN",
	2: "STAT",
}
var InitialGroup_MsgType_value = map[string] int32 {
	"JOIN": 1,
	"STAT": 2,
}
func NewInitialGroup_MsgType(x int32) *InitialGroup_MsgType {
	e := InitialGroup_MsgType(x)
	return &e
}

type InWorldGroup_MsgType int32
const (
	InWorldGroup_WALK = 1
)
var InWorldGroup_MsgType_name = map[int32] string {
	1: "WALK",
}
var InWorldGroup_MsgType_value = map[string] int32 {
	"WALK": 1,
}
func NewInWorldGroup_MsgType(x int32) *InWorldGroup_MsgType {
	e := InWorldGroup_MsgType(x)
	return &e
}

type Walk_Direction int32
const (
	Walk_DIRECTION_UP = 1
	Walk_DIRECTION_DOWN = 2
	Walk_DIRECTION_LEFT = 3
	Walk_DIRECTION_RIGHT = 4
)
var Walk_Direction_name = map[int32] string {
	1: "DIRECTION_UP",
	2: "DIRECTION_DOWN",
	3: "DIRECTION_LEFT",
	4: "DIRECTION_RIGHT",
}
var Walk_Direction_value = map[string] int32 {
	"DIRECTION_UP": 1,
	"DIRECTION_DOWN": 2,
	"DIRECTION_LEFT": 3,
	"DIRECTION_RIGHT": 4,
}
func NewWalk_Direction(x int32) *Walk_Direction {
	e := Walk_Direction(x)
	return &e
}

type InitialGroup struct {
	Join	*Join	"PB(bytes,1,opt,name=join)"
	Stat	*Stat	"PB(bytes,2,opt,name=stat)"
	XXX_unrecognized	[]byte
}
func (this *InitialGroup) Reset() {
	*this = InitialGroup{}
}
func NewInitialGroup() *InitialGroup {
	return new(InitialGroup)
}

type Join struct {
	Playername	*string	"PB(bytes,1,req,name=playername)"
	XXX_unrecognized	[]byte
}
func (this *Join) Reset() {
	*this = Join{}
}
func NewJoin() *Join {
	return new(Join)
}

type Stat struct {
	Key	*string	"PB(bytes,1,req,name=key)"
	XXX_unrecognized	[]byte
}
func (this *Stat) Reset() {
	*this = Stat{}
}
func NewStat() *Stat {
	return new(Stat)
}

type InWorldGroup struct {
	Walk	*Walk	"PB(bytes,1,opt,name=walk)"
	XXX_unrecognized	[]byte
}
func (this *InWorldGroup) Reset() {
	*this = InWorldGroup{}
}
func NewInWorldGroup() *InWorldGroup {
	return new(InWorldGroup)
}

type Walk struct {
	Direction	*Walk_Direction	"PB(varint,1,req,name=direction,enum=wanderer.Walk_Direction)"
	XXX_unrecognized	[]byte
}
func (this *Walk) Reset() {
	*this = Walk{}
}
func NewWalk() *Walk {
	return new(Walk)
}

func init() {
	proto.RegisterEnum("wanderer.InitialGroup_MsgType", InitialGroup_MsgType_name, InitialGroup_MsgType_value)
	proto.RegisterEnum("wanderer.InWorldGroup_MsgType", InWorldGroup_MsgType_name, InWorldGroup_MsgType_value)
	proto.RegisterEnum("wanderer.Walk_Direction", Walk_Direction_name, Walk_Direction_value)
}