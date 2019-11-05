// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfer/transfer.proto

package transfer

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

// RangeGenMode 线程分配方式
type RangeGenMode int32

const (
	// RangeGenMode_Default 根据parallel平均生成
	RangeGenMode_Default RangeGenMode = 0
	// RangeGenMode_BlockSize 根据blockSize生成
	RangeGenMode_BlockSize RangeGenMode = 1
)

var RangeGenMode_name = map[int32]string{
	0: "Default",
	1: "BlockSize",
}

var RangeGenMode_value = map[string]int32{
	"Default":   0,
	"BlockSize": 1,
}

func (x RangeGenMode) String() string {
	return proto.EnumName(RangeGenMode_name, int32(x))
}

func (RangeGenMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44038b0c710d7f2f, []int{0}
}

//Range 请求范围
type Range struct {
	Begin                int64    `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_44038b0c710d7f2f, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// DownloadInstanceInfoExport 断点续传
type DownloadInstanceInfoExport struct {
	RangeGenMode         RangeGenMode `protobuf:"varint,1,opt,name=range_gen_mode,json=rangeGenMode,proto3,enum=transfer.RangeGenMode" json:"range_gen_mode,omitempty"`
	TotalSize            int64        `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	GenBegin             int64        `protobuf:"varint,3,opt,name=gen_begin,json=genBegin,proto3" json:"gen_begin,omitempty"`
	BlockSize            int64        `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Ranges               []*Range     `protobuf:"bytes,5,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DownloadInstanceInfoExport) Reset()         { *m = DownloadInstanceInfoExport{} }
func (m *DownloadInstanceInfoExport) String() string { return proto.CompactTextString(m) }
func (*DownloadInstanceInfoExport) ProtoMessage()    {}
func (*DownloadInstanceInfoExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_44038b0c710d7f2f, []int{1}
}

func (m *DownloadInstanceInfoExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadInstanceInfoExport.Unmarshal(m, b)
}
func (m *DownloadInstanceInfoExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadInstanceInfoExport.Marshal(b, m, deterministic)
}
func (m *DownloadInstanceInfoExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadInstanceInfoExport.Merge(m, src)
}
func (m *DownloadInstanceInfoExport) XXX_Size() int {
	return xxx_messageInfo_DownloadInstanceInfoExport.Size(m)
}
func (m *DownloadInstanceInfoExport) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadInstanceInfoExport.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadInstanceInfoExport proto.InternalMessageInfo

func (m *DownloadInstanceInfoExport) GetRangeGenMode() RangeGenMode {
	if m != nil {
		return m.RangeGenMode
	}
	return RangeGenMode_Default
}

func (m *DownloadInstanceInfoExport) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *DownloadInstanceInfoExport) GetGenBegin() int64 {
	if m != nil {
		return m.GenBegin
	}
	return 0
}

func (m *DownloadInstanceInfoExport) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *DownloadInstanceInfoExport) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func init() {
	proto.RegisterEnum("transfer.RangeGenMode", RangeGenMode_name, RangeGenMode_value)
	proto.RegisterType((*Range)(nil), "transfer.Range")
	proto.RegisterType((*DownloadInstanceInfoExport)(nil), "transfer.DownloadInstanceInfoExport")
}

func init() { proto.RegisterFile("transfer/transfer.proto", fileDescriptor_44038b0c710d7f2f) }

var fileDescriptor_44038b0c710d7f2f = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x31, 0xb5, 0x99, 0xd6, 0x1a, 0x16, 0xd1, 0xa0, 0x14, 0x4a, 0x2f, 0x96, 0x1e,
	0x5a, 0xa8, 0x57, 0x4f, 0xa5, 0x22, 0x3d, 0x78, 0x89, 0x3f, 0x20, 0x6c, 0x9a, 0x49, 0x08, 0xc6,
	0x99, 0xb2, 0x59, 0x51, 0xfa, 0x63, 0xfd, 0x2d, 0x92, 0x49, 0x23, 0xa1, 0xb7, 0x7d, 0x6f, 0x78,
	0xef, 0x7d, 0x2c, 0xdc, 0x59, 0xa3, 0xa9, 0xca, 0xd0, 0x2c, 0xdb, 0xc7, 0x62, 0x6f, 0xd8, 0xb2,
	0xea, 0xb7, 0x7a, 0xba, 0x04, 0x2f, 0xd2, 0x94, 0xa3, 0xba, 0x01, 0x2f, 0xc1, 0xbc, 0xa0, 0xd0,
	0x99, 0x38, 0x33, 0x37, 0x6a, 0x84, 0x0a, 0xc0, 0x45, 0x4a, 0xc3, 0x73, 0xf1, 0xea, 0xe7, 0xf4,
	0xd7, 0x81, 0xfb, 0x0d, 0x7f, 0x53, 0xc9, 0x3a, 0xdd, 0x52, 0x65, 0x35, 0xed, 0x70, 0x4b, 0x19,
	0xbf, 0xfc, 0xec, 0xd9, 0x58, 0xf5, 0x0c, 0x23, 0x53, 0xf7, 0xc5, 0x39, 0x52, 0xfc, 0xc9, 0x29,
	0x4a, 0xdf, 0x68, 0x75, 0xbb, 0xf8, 0x47, 0x90, 0xbd, 0x57, 0xa4, 0x37, 0x4e, 0x31, 0x1a, 0x9a,
	0x8e, 0x52, 0x63, 0x00, 0xcb, 0x56, 0x97, 0x71, 0x55, 0x1c, 0xf0, 0xb8, 0xea, 0x8b, 0xf3, 0x5e,
	0x1c, 0x50, 0x3d, 0x80, 0x5f, 0xd7, 0x36, 0x9c, 0xae, 0x5c, 0xfb, 0x39, 0xd2, 0x5a, 0x50, 0xc7,
	0x00, 0x49, 0xc9, 0xbb, 0x8f, 0x26, 0x7b, 0xd1, 0x64, 0xc5, 0x91, 0xec, 0x23, 0xf4, 0x64, 0xaa,
	0x0a, 0xbd, 0x89, 0x3b, 0x1b, 0xac, 0xae, 0x4f, 0x80, 0xa2, 0xe3, 0x79, 0x3e, 0x87, 0x61, 0x97,
	0x50, 0x0d, 0xe0, 0x72, 0x83, 0x99, 0xfe, 0x2a, 0x6d, 0x70, 0xa6, 0xae, 0xc0, 0x5f, 0xb7, 0x95,
	0x81, 0x93, 0xf4, 0xe4, 0x3b, 0x9f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x5c, 0x09, 0xcc,
	0x69, 0x01, 0x00, 0x00,
}