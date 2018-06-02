// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/scraper.proto

package scraper

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

type ScrapeItem struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScrapeItem) Reset()         { *m = ScrapeItem{} }
func (m *ScrapeItem) String() string { return proto.CompactTextString(m) }
func (*ScrapeItem) ProtoMessage()    {}
func (*ScrapeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{0}
}
func (m *ScrapeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScrapeItem.Unmarshal(m, b)
}
func (m *ScrapeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScrapeItem.Marshal(b, m, deterministic)
}
func (dst *ScrapeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScrapeItem.Merge(dst, src)
}
func (m *ScrapeItem) XXX_Size() int {
	return xxx_messageInfo_ScrapeItem.Size(m)
}
func (m *ScrapeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ScrapeItem.DiscardUnknown(m)
}

var xxx_messageInfo_ScrapeItem proto.InternalMessageInfo

func (m *ScrapeItem) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ScrapeHistoryItem struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScrapeHistoryItem) Reset()         { *m = ScrapeHistoryItem{} }
func (m *ScrapeHistoryItem) String() string { return proto.CompactTextString(m) }
func (*ScrapeHistoryItem) ProtoMessage()    {}
func (*ScrapeHistoryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{1}
}
func (m *ScrapeHistoryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScrapeHistoryItem.Unmarshal(m, b)
}
func (m *ScrapeHistoryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScrapeHistoryItem.Marshal(b, m, deterministic)
}
func (dst *ScrapeHistoryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScrapeHistoryItem.Merge(dst, src)
}
func (m *ScrapeHistoryItem) XXX_Size() int {
	return xxx_messageInfo_ScrapeHistoryItem.Size(m)
}
func (m *ScrapeHistoryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ScrapeHistoryItem.DiscardUnknown(m)
}

var xxx_messageInfo_ScrapeHistoryItem proto.InternalMessageInfo

func (m *ScrapeHistoryItem) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ScrapeHistoryItem) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetWorkQueueRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkQueueRequest) Reset()         { *m = GetWorkQueueRequest{} }
func (m *GetWorkQueueRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkQueueRequest) ProtoMessage()    {}
func (*GetWorkQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{2}
}
func (m *GetWorkQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkQueueRequest.Unmarshal(m, b)
}
func (m *GetWorkQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkQueueRequest.Marshal(b, m, deterministic)
}
func (dst *GetWorkQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkQueueRequest.Merge(dst, src)
}
func (m *GetWorkQueueRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkQueueRequest.Size(m)
}
func (m *GetWorkQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkQueueRequest proto.InternalMessageInfo

type GetWorkQueueResponse struct {
	Items                []*ScrapeItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetWorkQueueResponse) Reset()         { *m = GetWorkQueueResponse{} }
func (m *GetWorkQueueResponse) String() string { return proto.CompactTextString(m) }
func (*GetWorkQueueResponse) ProtoMessage()    {}
func (*GetWorkQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{3}
}
func (m *GetWorkQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkQueueResponse.Unmarshal(m, b)
}
func (m *GetWorkQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkQueueResponse.Marshal(b, m, deterministic)
}
func (dst *GetWorkQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkQueueResponse.Merge(dst, src)
}
func (m *GetWorkQueueResponse) XXX_Size() int {
	return xxx_messageInfo_GetWorkQueueResponse.Size(m)
}
func (m *GetWorkQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkQueueResponse proto.InternalMessageInfo

func (m *GetWorkQueueResponse) GetItems() []*ScrapeItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatusRequest) Reset()         { *m = GetStatusRequest{} }
func (m *GetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatusRequest) ProtoMessage()    {}
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{4}
}
func (m *GetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusRequest.Unmarshal(m, b)
}
func (m *GetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusRequest.Merge(dst, src)
}
func (m *GetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatusRequest.Size(m)
}
func (m *GetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusRequest proto.InternalMessageInfo

type GetStatusResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatusResponse) Reset()         { *m = GetStatusResponse{} }
func (m *GetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatusResponse) ProtoMessage()    {}
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{5}
}
func (m *GetStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusResponse.Unmarshal(m, b)
}
func (m *GetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusResponse.Marshal(b, m, deterministic)
}
func (dst *GetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResponse.Merge(dst, src)
}
func (m *GetStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatusResponse.Size(m)
}
func (m *GetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResponse proto.InternalMessageInfo

func (m *GetStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetWorkHistoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkHistoryRequest) Reset()         { *m = GetWorkHistoryRequest{} }
func (m *GetWorkHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkHistoryRequest) ProtoMessage()    {}
func (*GetWorkHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{6}
}
func (m *GetWorkHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkHistoryRequest.Unmarshal(m, b)
}
func (m *GetWorkHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkHistoryRequest.Marshal(b, m, deterministic)
}
func (dst *GetWorkHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkHistoryRequest.Merge(dst, src)
}
func (m *GetWorkHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkHistoryRequest.Size(m)
}
func (m *GetWorkHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkHistoryRequest proto.InternalMessageInfo

type GetWorkHistoryResponse struct {
	Items                []*ScrapeHistoryItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetWorkHistoryResponse) Reset()         { *m = GetWorkHistoryResponse{} }
func (m *GetWorkHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetWorkHistoryResponse) ProtoMessage()    {}
func (*GetWorkHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scraper_9df590cbc21b8ea9, []int{7}
}
func (m *GetWorkHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkHistoryResponse.Unmarshal(m, b)
}
func (m *GetWorkHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkHistoryResponse.Marshal(b, m, deterministic)
}
func (dst *GetWorkHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkHistoryResponse.Merge(dst, src)
}
func (m *GetWorkHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetWorkHistoryResponse.Size(m)
}
func (m *GetWorkHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkHistoryResponse proto.InternalMessageInfo

func (m *GetWorkHistoryResponse) GetItems() []*ScrapeHistoryItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ScrapeItem)(nil), "scraper.ScrapeItem")
	proto.RegisterType((*ScrapeHistoryItem)(nil), "scraper.ScrapeHistoryItem")
	proto.RegisterType((*GetWorkQueueRequest)(nil), "scraper.GetWorkQueueRequest")
	proto.RegisterType((*GetWorkQueueResponse)(nil), "scraper.GetWorkQueueResponse")
	proto.RegisterType((*GetStatusRequest)(nil), "scraper.GetStatusRequest")
	proto.RegisterType((*GetStatusResponse)(nil), "scraper.GetStatusResponse")
	proto.RegisterType((*GetWorkHistoryRequest)(nil), "scraper.GetWorkHistoryRequest")
	proto.RegisterType((*GetWorkHistoryResponse)(nil), "scraper.GetWorkHistoryResponse")
}

func init() { proto.RegisterFile("proto/scraper.proto", fileDescriptor_scraper_9df590cbc21b8ea9) }

var fileDescriptor_scraper_9df590cbc21b8ea9 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x8d, 0xc1, 0x4a, 0x46, 0x29, 0xed, 0xc6, 0xd6, 0x1a, 0x6a, 0x2d, 0x7b, 0xaa, 0x08,
	0x55, 0xea, 0x13, 0x88, 0x42, 0x55, 0xe8, 0xc1, 0xe4, 0xe0, 0x39, 0x96, 0x39, 0x04, 0x4d, 0x37,
	0xee, 0x4e, 0x04, 0x5f, 0xdd, 0x93, 0x74, 0xb3, 0xf9, 0xb3, 0xa1, 0xbd, 0xed, 0xcc, 0x6f, 0xf2,
	0xf1, 0xcd, 0x10, 0xf0, 0x33, 0x29, 0x48, 0xdc, 0xaa, 0xb5, 0x8c, 0x33, 0x94, 0x73, 0x5d, 0xb1,
	0x63, 0x53, 0xf2, 0x09, 0x40, 0xa4, 0x9f, 0x2f, 0x84, 0x29, 0xeb, 0x81, 0x9b, 0xcb, 0xaf, 0x91,
	0x33, 0x75, 0x66, 0x5e, 0xb8, 0x7d, 0xf2, 0x47, 0xe8, 0x17, 0xf9, 0x73, 0xa2, 0x48, 0xc8, 0x5f,
	0x3d, 0x36, 0x06, 0x8f, 0x92, 0x14, 0x15, 0xc5, 0x69, 0xa6, 0x87, 0xdd, 0xb0, 0x6e, 0x94, 0x90,
	0xc3, 0x1a, 0x32, 0x00, 0x7f, 0x89, 0xf4, 0x2e, 0xe4, 0xe7, 0x5b, 0x8e, 0x39, 0x86, 0xf8, 0x9d,
	0xa3, 0x22, 0xfe, 0x00, 0x67, 0x76, 0x5b, 0x65, 0x62, 0xa3, 0x90, 0x5d, 0xc3, 0x51, 0x42, 0x98,
	0xaa, 0x91, 0x33, 0x75, 0x67, 0x27, 0x0b, 0x7f, 0x5e, 0xba, 0xd7, 0xa6, 0x61, 0x31, 0xc1, 0x19,
	0xf4, 0x96, 0x48, 0x11, 0xc5, 0x94, 0xab, 0x12, 0x7b, 0x03, 0xfd, 0x46, 0xcf, 0x30, 0x87, 0xd0,
	0x51, 0xba, 0x63, 0x96, 0x33, 0x15, 0x3f, 0x87, 0x81, 0x71, 0x30, 0x0b, 0x96, 0x94, 0x57, 0x18,
	0xb6, 0x03, 0x83, 0xba, 0xb3, 0xf5, 0x82, 0x96, 0x5e, 0xe3, 0x50, 0xc6, 0x72, 0xf1, 0xe7, 0x40,
	0x77, 0x25, 0x36, 0x09, 0x09, 0x19, 0xa1, 0xfc, 0x49, 0xd6, 0xc8, 0x56, 0x70, 0xda, 0xdc, 0x9d,
	0x8d, 0x2b, 0xca, 0x8e, 0x4b, 0x05, 0x97, 0x7b, 0xd2, 0xc2, 0x88, 0x1f, 0xb0, 0x08, 0xba, 0xb6,
	0x2d, 0x9b, 0xb4, 0x3f, 0xb1, 0xf7, 0x0b, 0xae, 0xf6, 0xe6, 0x15, 0xf4, 0x09, 0xbc, 0xea, 0x90,
	0xec, 0xa2, 0x39, 0x6f, 0x1d, 0x3c, 0x08, 0x76, 0x45, 0x25, 0xe5, 0xa3, 0xa3, 0xff, 0xb8, 0xfb,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0xa3, 0xe4, 0x53, 0x88, 0x02, 0x00, 0x00,
}
