// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p2beta1/geometry.proto

package vision // import "google.golang.org/genproto/googleapis/cloud/vision/v1p2beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A vertex represents a 2D point in the image.
// NOTE: the vertex coordinates are in the same scale as the original image.
type Vertex struct {
	// X coordinate.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_e749cb92138e5a14, []int{0}
}
func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A vertex represents a 2D point in the image.
// NOTE: the normalized vertex coordinates are relative to the original image
// and range from 0 to 1.
type NormalizedVertex struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalizedVertex) Reset()         { *m = NormalizedVertex{} }
func (m *NormalizedVertex) String() string { return proto.CompactTextString(m) }
func (*NormalizedVertex) ProtoMessage()    {}
func (*NormalizedVertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_e749cb92138e5a14, []int{1}
}
func (m *NormalizedVertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedVertex.Unmarshal(m, b)
}
func (m *NormalizedVertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedVertex.Marshal(b, m, deterministic)
}
func (m *NormalizedVertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedVertex.Merge(m, src)
}
func (m *NormalizedVertex) XXX_Size() int {
	return xxx_messageInfo_NormalizedVertex.Size(m)
}
func (m *NormalizedVertex) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedVertex.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedVertex proto.InternalMessageInfo

func (m *NormalizedVertex) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NormalizedVertex) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon for the detected image annotation.
type BoundingPoly struct {
	// The bounding polygon vertices.
	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	// The bounding polygon normalized vertices.
	NormalizedVertices   []*NormalizedVertex `protobuf:"bytes,2,rep,name=normalized_vertices,json=normalizedVertices,proto3" json:"normalized_vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BoundingPoly) Reset()         { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()    {}
func (*BoundingPoly) Descriptor() ([]byte, []int) {
	return fileDescriptor_e749cb92138e5a14, []int{2}
}
func (m *BoundingPoly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPoly.Unmarshal(m, b)
}
func (m *BoundingPoly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPoly.Marshal(b, m, deterministic)
}
func (m *BoundingPoly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPoly.Merge(m, src)
}
func (m *BoundingPoly) XXX_Size() int {
	return xxx_messageInfo_BoundingPoly.Size(m)
}
func (m *BoundingPoly) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPoly.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPoly proto.InternalMessageInfo

func (m *BoundingPoly) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *BoundingPoly) GetNormalizedVertices() []*NormalizedVertex {
	if m != nil {
		return m.NormalizedVertices
	}
	return nil
}

// A 3D position in the image, used primarily for Face detection landmarks.
// A valid Position must have both x and y coordinates.
// The position coordinates are in the same scale as the original image.
type Position struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	// Z coordinate (or depth).
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_e749cb92138e5a14, []int{3}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vertex)(nil), "google.cloud.vision.v1p2beta1.Vertex")
	proto.RegisterType((*NormalizedVertex)(nil), "google.cloud.vision.v1p2beta1.NormalizedVertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.vision.v1p2beta1.BoundingPoly")
	proto.RegisterType((*Position)(nil), "google.cloud.vision.v1p2beta1.Position")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p2beta1/geometry.proto", fileDescriptor_e749cb92138e5a14)
}

var fileDescriptor_e749cb92138e5a14 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x49, 0x87, 0x63, 0xc4, 0x09, 0x52, 0x2f, 0xbd, 0x08, 0xb3, 0x28, 0xec, 0x20, 0x09,
	0x9b, 0xde, 0x3c, 0x59, 0x0f, 0xde, 0xa4, 0xf4, 0xe0, 0xc1, 0x8b, 0x76, 0xed, 0x23, 0x04, 0xda,
	0xbc, 0x92, 0x66, 0x65, 0x2d, 0xfe, 0x57, 0xfe, 0x73, 0x1e, 0xa5, 0xc9, 0x28, 0x6c, 0x60, 0x77,
	0xfc, 0x5e, 0x7e, 0xef, 0x7b, 0x5f, 0xf8, 0xe8, 0xbd, 0x40, 0x14, 0x05, 0xf0, 0xac, 0xc0, 0x6d,
	0xce, 0x1b, 0x59, 0x4b, 0x54, 0xbc, 0x59, 0x55, 0xeb, 0x0d, 0x98, 0x74, 0xc5, 0x05, 0x60, 0x09,
	0x46, 0xb7, 0xac, 0xd2, 0x68, 0xd0, 0xbf, 0x76, 0x34, 0xb3, 0x34, 0x73, 0x34, 0x1b, 0xe8, 0xf0,
	0x96, 0x4e, 0xdf, 0x41, 0x1b, 0xd8, 0xf9, 0x73, 0x4a, 0x76, 0x01, 0x59, 0x90, 0xe5, 0x59, 0x42,
	0xac, 0x6a, 0x03, 0xcf, 0xa9, 0x36, 0x64, 0xf4, 0xf2, 0x0d, 0x75, 0x99, 0x16, 0xb2, 0x83, 0xfc,
	0x98, 0xf7, 0x0e, 0x78, 0xaf, 0xe7, 0x7f, 0x08, 0x9d, 0x47, 0xb8, 0x55, 0xb9, 0x54, 0x22, 0xc6,
	0xa2, 0xf5, 0x9f, 0xe9, 0xac, 0x01, 0x6d, 0x64, 0x06, 0x75, 0x40, 0x16, 0x93, 0xe5, 0xf9, 0xfa,
	0x8e, 0x8d, 0x06, 0x63, 0xee, 0x4a, 0x32, 0xac, 0xf9, 0x5f, 0xf4, 0x4a, 0x0d, 0x19, 0x3e, 0x07,
	0x37, 0xcf, 0xba, 0xf1, 0x13, 0x6e, 0xc7, 0xe9, 0x13, 0x5f, 0x1d, 0x4c, 0x7a, 0xab, 0xf0, 0x91,
	0xce, 0x62, 0xac, 0xa5, 0x91, 0xa8, 0xc6, 0x7e, 0xd7, 0xab, 0x2e, 0x98, 0x38, 0xd5, 0x45, 0xdf,
	0xf4, 0x26, 0xc3, 0x72, 0xfc, 0x7e, 0x74, 0xf1, 0xba, 0x6f, 0x25, 0xee, 0x4b, 0x89, 0xc9, 0xc7,
	0xcb, 0x9e, 0x17, 0x58, 0xa4, 0x4a, 0x30, 0xd4, 0x82, 0x0b, 0x50, 0xb6, 0x32, 0xee, 0x9e, 0xd2,
	0x4a, 0xd6, 0xff, 0x74, 0xfc, 0xe4, 0x06, 0xbf, 0x84, 0x6c, 0xa6, 0x76, 0xe5, 0xe1, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0xe4, 0x63, 0xcf, 0x15, 0x02, 0x00, 0x00,
}
