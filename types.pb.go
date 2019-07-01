// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package icsgo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// channel tell
type ChannelTell struct {
	// channel name
	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// user who made the tell
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// message / body of the tell
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelTell) Reset()         { *m = ChannelTell{} }
func (m *ChannelTell) String() string { return proto.CompactTextString(m) }
func (*ChannelTell) ProtoMessage()    {}
func (*ChannelTell) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *ChannelTell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelTell.Unmarshal(m, b)
}
func (m *ChannelTell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelTell.Marshal(b, m, deterministic)
}
func (m *ChannelTell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelTell.Merge(m, src)
}
func (m *ChannelTell) XXX_Size() int {
	return xxx_messageInfo_ChannelTell.Size(m)
}
func (m *ChannelTell) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelTell.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelTell proto.InternalMessageInfo

func (m *ChannelTell) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChannelTell) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ChannelTell) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// private tell
type PrivateTell struct {
	// user who made the tell
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// message / body of the tell
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateTell) Reset()         { *m = PrivateTell{} }
func (m *PrivateTell) String() string { return proto.CompactTextString(m) }
func (*PrivateTell) ProtoMessage()    {}
func (*PrivateTell) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *PrivateTell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateTell.Unmarshal(m, b)
}
func (m *PrivateTell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateTell.Marshal(b, m, deterministic)
}
func (m *PrivateTell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateTell.Merge(m, src)
}
func (m *PrivateTell) XXX_Size() int {
	return xxx_messageInfo_PrivateTell.Size(m)
}
func (m *PrivateTell) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateTell.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateTell proto.InternalMessageInfo

func (m *PrivateTell) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *PrivateTell) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// a game start message
type GameStart struct {
	// id of the game that was started
	GameId uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	// handle of player one
	PlayerOne string `protobuf:"bytes,2,opt,name=player_one,json=playerOne,proto3" json:"player_one,omitempty"`
	// handle of player two
	PlayerTwo            string   `protobuf:"bytes,3,opt,name=player_two,json=playerTwo,proto3" json:"player_two,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStart) Reset()         { *m = GameStart{} }
func (m *GameStart) String() string { return proto.CompactTextString(m) }
func (*GameStart) ProtoMessage()    {}
func (*GameStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *GameStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStart.Unmarshal(m, b)
}
func (m *GameStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStart.Marshal(b, m, deterministic)
}
func (m *GameStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStart.Merge(m, src)
}
func (m *GameStart) XXX_Size() int {
	return xxx_messageInfo_GameStart.Size(m)
}
func (m *GameStart) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStart.DiscardUnknown(m)
}

var xxx_messageInfo_GameStart proto.InternalMessageInfo

func (m *GameStart) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameStart) GetPlayerOne() string {
	if m != nil {
		return m.PlayerOne
	}
	return ""
}

func (m *GameStart) GetPlayerTwo() string {
	if m != nil {
		return m.PlayerTwo
	}
	return ""
}

// a game end message
type GameEnd struct {
	// id of the game that was ended
	GameId uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	// user who won the game, if it wasn't a draw
	Winner string `protobuf:"bytes,2,opt,name=winner,proto3" json:"winner,omitempty"`
	// handle of the user who lost the game, if it wasn't a draw
	Loser string `protobuf:"bytes,3,opt,name=loser,proto3" json:"loser,omitempty"`
	// reason for game end
	Reason uint32 `protobuf:"varint,4,opt,name=reason,proto3" json:"reason,omitempty"`
	// message associated with the game result
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameEnd) Reset()         { *m = GameEnd{} }
func (m *GameEnd) String() string { return proto.CompactTextString(m) }
func (*GameEnd) ProtoMessage()    {}
func (*GameEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *GameEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEnd.Unmarshal(m, b)
}
func (m *GameEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEnd.Marshal(b, m, deterministic)
}
func (m *GameEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEnd.Merge(m, src)
}
func (m *GameEnd) XXX_Size() int {
	return xxx_messageInfo_GameEnd.Size(m)
}
func (m *GameEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEnd.DiscardUnknown(m)
}

var xxx_messageInfo_GameEnd proto.InternalMessageInfo

func (m *GameEnd) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameEnd) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *GameEnd) GetLoser() string {
	if m != nil {
		return m.Loser
	}
	return ""
}

func (m *GameEnd) GetReason() uint32 {
	if m != nil {
		return m.Reason
	}
	return 0
}

func (m *GameEnd) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// a game move message
type GameMove struct {
	Fen                  string   `protobuf:"bytes,1,opt,name=fen,proto3" json:"fen,omitempty"`
	Turn                 string   `protobuf:"bytes,2,opt,name=turn,proto3" json:"turn,omitempty"`
	GameId               uint32   `protobuf:"varint,3,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	WhiteName            string   `protobuf:"bytes,4,opt,name=white_name,json=whiteName,proto3" json:"white_name,omitempty"`
	BlackName            string   `protobuf:"bytes,5,opt,name=black_name,json=blackName,proto3" json:"black_name,omitempty"`
	Role                 uint32   `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	Time                 uint32   `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	Inc                  uint32   `protobuf:"varint,8,opt,name=inc,proto3" json:"inc,omitempty"`
	WhiteTime            uint32   `protobuf:"varint,9,opt,name=white_time,json=whiteTime,proto3" json:"white_time,omitempty"`
	BlackTime            uint32   `protobuf:"varint,10,opt,name=black_time,json=blackTime,proto3" json:"black_time,omitempty"`
	Move                 string   `protobuf:"bytes,11,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameMove) Reset()         { *m = GameMove{} }
func (m *GameMove) String() string { return proto.CompactTextString(m) }
func (*GameMove) ProtoMessage()    {}
func (*GameMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *GameMove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameMove.Unmarshal(m, b)
}
func (m *GameMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameMove.Marshal(b, m, deterministic)
}
func (m *GameMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameMove.Merge(m, src)
}
func (m *GameMove) XXX_Size() int {
	return xxx_messageInfo_GameMove.Size(m)
}
func (m *GameMove) XXX_DiscardUnknown() {
	xxx_messageInfo_GameMove.DiscardUnknown(m)
}

var xxx_messageInfo_GameMove proto.InternalMessageInfo

func (m *GameMove) GetFen() string {
	if m != nil {
		return m.Fen
	}
	return ""
}

func (m *GameMove) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

func (m *GameMove) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameMove) GetWhiteName() string {
	if m != nil {
		return m.WhiteName
	}
	return ""
}

func (m *GameMove) GetBlackName() string {
	if m != nil {
		return m.BlackName
	}
	return ""
}

func (m *GameMove) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *GameMove) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GameMove) GetInc() uint32 {
	if m != nil {
		return m.Inc
	}
	return 0
}

func (m *GameMove) GetWhiteTime() uint32 {
	if m != nil {
		return m.WhiteTime
	}
	return 0
}

func (m *GameMove) GetBlackTime() uint32 {
	if m != nil {
		return m.BlackTime
	}
	return 0
}

func (m *GameMove) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

// a generic message from the server
type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelTell)(nil), "icsgo.ChannelTell")
	proto.RegisterType((*PrivateTell)(nil), "icsgo.PrivateTell")
	proto.RegisterType((*GameStart)(nil), "icsgo.GameStart")
	proto.RegisterType((*GameEnd)(nil), "icsgo.GameEnd")
	proto.RegisterType((*GameMove)(nil), "icsgo.GameMove")
	proto.RegisterType((*Message)(nil), "icsgo.Message")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0x56, 0xda, 0x34, 0x69, 0x2e, 0x7a, 0xa5, 0x57, 0x16, 0x02, 0x2f, 0x48, 0x28, 0x2c, 0x4c,
	0x2c, 0x8c, 0x8c, 0x08, 0x21, 0x86, 0x02, 0x2a, 0x65, 0xae, 0xdc, 0xf4, 0x68, 0x2d, 0x62, 0xbb,
	0x72, 0xdc, 0x46, 0x5d, 0xf9, 0x0b, 0xfc, 0x61, 0xe4, 0x73, 0x02, 0x29, 0x82, 0xed, 0xf9, 0xf0,
	0x7d, 0x3d, 0x32, 0xe4, 0x6e, 0xbf, 0xc1, 0xfa, 0x72, 0x63, 0x8d, 0x33, 0x6c, 0x24, 0xcb, 0x7a,
	0x65, 0x8a, 0x17, 0xc8, 0x6f, 0xd6, 0x42, 0x6b, 0xac, 0x66, 0x58, 0x55, 0x8c, 0x43, 0x5a, 0x06,
	0xca, 0xa3, 0xb3, 0xe8, 0x22, 0x9b, 0x76, 0x94, 0x31, 0x88, 0xb7, 0x35, 0x5a, 0x3e, 0x20, 0x99,
	0xb0, 0x7f, 0xad, 0xb0, 0xae, 0xc5, 0x0a, 0xf9, 0x30, 0xbc, 0x6e, 0x69, 0x71, 0x0d, 0xf9, 0x93,
	0x95, 0x3b, 0xe1, 0x90, 0xda, 0x76, 0xc5, 0xd1, 0xef, 0xc5, 0x83, 0xc3, 0xe2, 0x05, 0x64, 0x77,
	0x42, 0xe1, 0xb3, 0x13, 0xd6, 0xb1, 0x13, 0x48, 0x57, 0x42, 0xe1, 0x5c, 0x2e, 0xa9, 0xfa, 0xdf,
	0x34, 0xf1, 0xf4, 0x7e, 0xc9, 0x4e, 0x01, 0x36, 0x95, 0xd8, 0xa3, 0x9d, 0x1b, 0xdd, 0xb5, 0xc8,
	0x82, 0xf2, 0xa8, 0xb1, 0x67, 0xbb, 0xc6, 0xb4, 0xeb, 0xb5, 0xf6, 0xac, 0x31, 0xc5, 0x7b, 0x04,
	0xa9, 0x1f, 0x72, 0xab, 0x97, 0x7f, 0x8f, 0x38, 0x86, 0xa4, 0x91, 0x5a, 0x7f, 0x5d, 0xdd, 0x32,
	0x76, 0x04, 0xa3, 0xca, 0xf8, 0x7b, 0x42, 0xdb, 0x40, 0xfc, 0x6b, 0x8b, 0xa2, 0x36, 0x9a, 0xc7,
	0xa1, 0x4b, 0x60, 0xfd, 0x43, 0x47, 0x87, 0x87, 0x7e, 0x0c, 0x60, 0xec, 0x97, 0x98, 0x98, 0x1d,
	0xb2, 0xff, 0x30, 0x7c, 0x45, 0xdd, 0x46, 0xe4, 0xa1, 0x4f, 0xcd, 0x6d, 0xad, 0xee, 0x22, 0xf7,
	0xb8, 0xbf, 0xeb, 0xf0, 0x67, 0x1c, 0xcd, 0x5a, 0x3a, 0x9c, 0x6b, 0xa1, 0x90, 0x36, 0xc8, 0xa6,
	0x19, 0x29, 0x0f, 0x42, 0x51, 0x1c, 0x8b, 0x4a, 0x94, 0x6f, 0xc1, 0x0e, 0x7b, 0x64, 0xa4, 0x90,
	0xcd, 0x20, 0xb6, 0xa6, 0x42, 0x9e, 0x50, 0x4f, 0xc2, 0x34, 0x5e, 0x2a, 0xe4, 0x69, 0xd0, 0x3c,
	0xf6, 0x4b, 0x4a, 0x5d, 0xf2, 0x31, 0x49, 0x1e, 0x7e, 0xcf, 0xa5, 0xb7, 0x19, 0x19, 0x61, 0xee,
	0x4c, 0xf6, 0xe7, 0x92, 0x0d, 0xc1, 0x26, 0x85, 0x6c, 0x06, 0xb1, 0x32, 0x3b, 0xe4, 0x79, 0x38,
	0xd1, 0xe3, 0xe2, 0x1c, 0xd2, 0x49, 0x08, 0xa8, 0x1f, 0x5d, 0x74, 0x10, 0xdd, 0x22, 0xa1, 0x5f,
	0x7c, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x83, 0x36, 0x43, 0xd4, 0x02, 0x00, 0x00,
}
