// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tlogpb/tlog.proto

package tlogpb

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

type Record struct {
	Labels               *Labels     `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	Location             *Location   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Message              *Message    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SpanStart            *SpanStart  `protobuf:"bytes,4,opt,name=span_start,json=spanStart,proto3" json:"span_start,omitempty"`
	SpanFinish           *SpanFinish `protobuf:"bytes,5,opt,name=span_finish,json=spanFinish,proto3" json:"span_finish,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetLabels() *Labels {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Record) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Record) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Record) GetSpanStart() *SpanStart {
	if m != nil {
		return m.SpanStart
	}
	return nil
}

func (m *Record) GetSpanFinish() *SpanFinish {
	if m != nil {
		return m.SpanFinish
	}
	return nil
}

type Labels struct {
	Label                []string `protobuf:"bytes,1,rep,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Labels) Reset()         { *m = Labels{} }
func (m *Labels) String() string { return proto.CompactTextString(m) }
func (*Labels) ProtoMessage()    {}
func (*Labels) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{1}
}

func (m *Labels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Labels.Unmarshal(m, b)
}
func (m *Labels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Labels.Marshal(b, m, deterministic)
}
func (m *Labels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Labels.Merge(m, src)
}
func (m *Labels) XXX_Size() int {
	return xxx_messageInfo_Labels.Size(m)
}
func (m *Labels) XXX_DiscardUnknown() {
	xxx_messageInfo_Labels.DiscardUnknown(m)
}

var xxx_messageInfo_Labels proto.InternalMessageInfo

func (m *Labels) GetLabel() []string {
	if m != nil {
		return m.Label
	}
	return nil
}

type Location struct {
	Pc                   int64    `protobuf:"varint,1,opt,name=pc,proto3" json:"pc,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	File                 string   `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Line                 int32    `protobuf:"varint,4,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{2}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetPc() int64 {
	if m != nil {
		return m.Pc
	}
	return 0
}

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Location) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

type Message struct {
	Span                 []byte   `protobuf:"bytes,1,opt,name=span,proto3" json:"span,omitempty"`
	Location             int64    `protobuf:"varint,2,opt,name=location,proto3" json:"location,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{3}
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

func (m *Message) GetSpan() []byte {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *Message) GetLocation() int64 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *Message) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SpanStart struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               []byte   `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Location             int64    `protobuf:"varint,3,opt,name=location,proto3" json:"location,omitempty"`
	Started              int64    `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanStart) Reset()         { *m = SpanStart{} }
func (m *SpanStart) String() string { return proto.CompactTextString(m) }
func (*SpanStart) ProtoMessage()    {}
func (*SpanStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{4}
}

func (m *SpanStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanStart.Unmarshal(m, b)
}
func (m *SpanStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanStart.Marshal(b, m, deterministic)
}
func (m *SpanStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanStart.Merge(m, src)
}
func (m *SpanStart) XXX_Size() int {
	return xxx_messageInfo_SpanStart.Size(m)
}
func (m *SpanStart) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanStart.DiscardUnknown(m)
}

var xxx_messageInfo_SpanStart proto.InternalMessageInfo

func (m *SpanStart) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SpanStart) GetParent() []byte {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *SpanStart) GetLocation() int64 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *SpanStart) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

type SpanFinish struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Elapsed              int64    `protobuf:"varint,2,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanFinish) Reset()         { *m = SpanFinish{} }
func (m *SpanFinish) String() string { return proto.CompactTextString(m) }
func (*SpanFinish) ProtoMessage()    {}
func (*SpanFinish) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{5}
}

func (m *SpanFinish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanFinish.Unmarshal(m, b)
}
func (m *SpanFinish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanFinish.Marshal(b, m, deterministic)
}
func (m *SpanFinish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanFinish.Merge(m, src)
}
func (m *SpanFinish) XXX_Size() int {
	return xxx_messageInfo_SpanFinish.Size(m)
}
func (m *SpanFinish) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanFinish.DiscardUnknown(m)
}

var xxx_messageInfo_SpanFinish proto.InternalMessageInfo

func (m *SpanFinish) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SpanFinish) GetElapsed() int64 {
	if m != nil {
		return m.Elapsed
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "tlogpb.Record")
	proto.RegisterType((*Labels)(nil), "tlogpb.Labels")
	proto.RegisterType((*Location)(nil), "tlogpb.Location")
	proto.RegisterType((*Message)(nil), "tlogpb.Message")
	proto.RegisterType((*SpanStart)(nil), "tlogpb.SpanStart")
	proto.RegisterType((*SpanFinish)(nil), "tlogpb.SpanFinish")
}

func init() { proto.RegisterFile("tlogpb/tlog.proto", fileDescriptor_ec43dc181f2a6e80) }

var fileDescriptor_ec43dc181f2a6e80 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0xc5, 0x76, 0x62, 0xc7, 0x93, 0x90, 0xdd, 0x0c, 0xcb, 0x22, 0xf6, 0xb0, 0x04, 0x1f, 0x4a,
	0x0a, 0x25, 0x2d, 0x0d, 0xf4, 0x13, 0x7a, 0x6a, 0x2f, 0x13, 0xe8, 0xb5, 0x28, 0xb6, 0x92, 0x0a,
	0x1c, 0x5b, 0x58, 0x3e, 0xf4, 0xb3, 0xfb, 0x09, 0x45, 0x23, 0x2b, 0x4d, 0xd3, 0x53, 0xde, 0xbc,
	0x79, 0xd2, 0xd3, 0x7b, 0x31, 0x2c, 0xfa, 0xba, 0x3d, 0x98, 0xdd, 0xad, 0xfb, 0x59, 0x9b, 0xae,
	0xed, 0x5b, 0x4c, 0x3d, 0x55, 0x7c, 0x44, 0x90, 0x92, 0x2a, 0xdb, 0xae, 0xc2, 0x2b, 0x48, 0x6b,
	0xb9, 0x53, 0xb5, 0x15, 0xd1, 0x32, 0x5a, 0x4d, 0xef, 0xe7, 0x6b, 0xaf, 0x59, 0x3f, 0x31, 0x4b,
	0xc3, 0x16, 0x6f, 0x60, 0x52, 0xb7, 0xa5, 0xec, 0x75, 0xdb, 0x88, 0x98, 0x95, 0xbf, 0x4f, 0xca,
	0x81, 0xa7, 0x93, 0x02, 0xaf, 0x21, 0x3b, 0x2a, 0x6b, 0xe5, 0x41, 0x89, 0x84, 0xc5, 0xbf, 0x82,
	0xf8, 0xd9, 0xd3, 0x14, 0xf6, 0x78, 0x07, 0x60, 0x8d, 0x6c, 0x5e, 0x6d, 0x2f, 0xbb, 0x5e, 0x8c,
	0x58, 0xbd, 0x08, 0xea, 0xad, 0x91, 0xcd, 0xd6, 0x2d, 0x28, 0xb7, 0x01, 0xe2, 0x06, 0xa6, 0x7c,
	0x62, 0xaf, 0x1b, 0x6d, 0xdf, 0xc4, 0x98, 0x8f, 0xe0, 0xf9, 0x91, 0x47, 0xde, 0x10, 0x5f, 0xec,
	0x71, 0xf1, 0x1f, 0x52, 0x9f, 0x08, 0xff, 0xc0, 0x98, 0x33, 0x89, 0x68, 0x99, 0xac, 0x72, 0xf2,
	0x43, 0xf1, 0x02, 0x93, 0x90, 0x03, 0xe7, 0x10, 0x9b, 0x92, 0xfb, 0x48, 0x28, 0x36, 0x25, 0x22,
	0x8c, 0x1a, 0x79, 0x54, 0x9c, 0x3b, 0x27, 0xc6, 0x8e, 0xdb, 0xeb, 0xda, 0xc7, 0xcb, 0x89, 0xb1,
	0xe3, 0x6a, 0xdd, 0x28, 0x0e, 0x31, 0x26, 0xc6, 0x85, 0x84, 0x6c, 0x88, 0xec, 0xd6, 0xee, 0x41,
	0x7c, 0xf1, 0x8c, 0x18, 0xe3, 0xbf, 0x8b, 0x5a, 0x93, 0xb3, 0x12, 0x11, 0x46, 0xbd, 0x3e, 0x7a,
	0x8b, 0x84, 0x18, 0x33, 0xa7, 0xde, 0x7d, 0x4f, 0x39, 0x31, 0x2e, 0x34, 0xe4, 0xa7, 0x9e, 0xdc,
	0xdb, 0x75, 0x35, 0x58, 0xc4, 0xba, 0xc2, 0xbf, 0x90, 0x1a, 0xd9, 0xa9, 0xa6, 0xe7, 0xeb, 0x67,
	0x34, 0x4c, 0xdf, 0x8c, 0x93, 0x0b, 0x63, 0x01, 0x19, 0xff, 0x1b, 0xaa, 0x62, 0x9f, 0x84, 0xc2,
	0x58, 0x3c, 0x00, 0x7c, 0xf5, 0xfb, 0xc3, 0x4b, 0x40, 0xa6, 0x6a, 0x69, 0xac, 0xaa, 0x86, 0x2c,
	0x61, 0xdc, 0xa5, 0xfc, 0xfd, 0x6d, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x65, 0x75, 0x1e, 0x72,
	0x94, 0x02, 0x00, 0x00,
}
