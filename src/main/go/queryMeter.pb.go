// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queryMeter.proto

package org_demo_queryMeter

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

type QueryResponse_InitResponseEnum int32

const (
	QueryResponse_CONTINUE  QueryResponse_InitResponseEnum = 0
	QueryResponse_RECONNECT QueryResponse_InitResponseEnum = 1
	QueryResponse_MOVE_TO   QueryResponse_InitResponseEnum = 2
)

var QueryResponse_InitResponseEnum_name = map[int32]string{
	0: "CONTINUE",
	1: "RECONNECT",
	2: "MOVE_TO",
}

var QueryResponse_InitResponseEnum_value = map[string]int32{
	"CONTINUE":  0,
	"RECONNECT": 1,
	"MOVE_TO":   2,
}

func (x QueryResponse_InitResponseEnum) String() string {
	return proto.EnumName(QueryResponse_InitResponseEnum_name, int32(x))
}

func (QueryResponse_InitResponseEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{1, 0}
}

type GetDataResponse_InitResponseEnum int32

const (
	GetDataResponse_CONTINUE  GetDataResponse_InitResponseEnum = 0
	GetDataResponse_RECONNECT GetDataResponse_InitResponseEnum = 1
)

var GetDataResponse_InitResponseEnum_name = map[int32]string{
	0: "CONTINUE",
	1: "RECONNECT",
}

var GetDataResponse_InitResponseEnum_value = map[string]int32{
	"CONTINUE":  0,
	"RECONNECT": 1,
}

func (x GetDataResponse_InitResponseEnum) String() string {
	return proto.EnumName(GetDataResponse_InitResponseEnum_name, int32(x))
}

func (GetDataResponse_InitResponseEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{5, 0}
}

// tag::QueryRequest[]
type QueryRequest struct {
	MeterUuid            [][]byte `protobuf:"bytes,1,rep,name=meterUuid,proto3" json:"meterUuid,omitempty"`
	FromTime             int32    `protobuf:"varint,2,opt,name=fromTime,proto3" json:"fromTime,omitempty"`
	ToTime               int32    `protobuf:"varint,3,opt,name=toTime,proto3" json:"toTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetMeterUuid() [][]byte {
	if m != nil {
		return m.MeterUuid
	}
	return nil
}

func (m *QueryRequest) GetFromTime() int32 {
	if m != nil {
		return m.FromTime
	}
	return 0
}

func (m *QueryRequest) GetToTime() int32 {
	if m != nil {
		return m.ToTime
	}
	return 0
}

type QueryResponse struct {
	Action               QueryResponse_InitResponseEnum `protobuf:"varint,1,opt,name=action,proto3,enum=org.demo.queryMeter.QueryResponse_InitResponseEnum" json:"action,omitempty"`
	StreamUid            string                         `protobuf:"bytes,2,opt,name=streamUid,proto3" json:"streamUid,omitempty"`
	Url                  string                         `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{1}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetAction() QueryResponse_InitResponseEnum {
	if m != nil {
		return m.Action
	}
	return QueryResponse_CONTINUE
}

func (m *QueryResponse) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

func (m *QueryResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// tag::GetDataRequest[]
type GetDataRequest struct {
	StreamUid            string   `protobuf:"bytes,1,opt,name=streamUid,proto3" json:"streamUid,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	BackpressureCount    int32    `protobuf:"varint,3,opt,name=backpressureCount,proto3" json:"backpressureCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{2}
}

func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (m *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(m, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

func (m *GetDataRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetDataRequest) GetBackpressureCount() int32 {
	if m != nil {
		return m.BackpressureCount
	}
	return 0
}

type MeterMeasurement struct {
	Timestamp            int32    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Measurement          int32    `protobuf:"varint,2,opt,name=measurement,proto3" json:"measurement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeterMeasurement) Reset()         { *m = MeterMeasurement{} }
func (m *MeterMeasurement) String() string { return proto.CompactTextString(m) }
func (*MeterMeasurement) ProtoMessage()    {}
func (*MeterMeasurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{3}
}

func (m *MeterMeasurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeterMeasurement.Unmarshal(m, b)
}
func (m *MeterMeasurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeterMeasurement.Marshal(b, m, deterministic)
}
func (m *MeterMeasurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeterMeasurement.Merge(m, src)
}
func (m *MeterMeasurement) XXX_Size() int {
	return xxx_messageInfo_MeterMeasurement.Size(m)
}
func (m *MeterMeasurement) XXX_DiscardUnknown() {
	xxx_messageInfo_MeterMeasurement.DiscardUnknown(m)
}

var xxx_messageInfo_MeterMeasurement proto.InternalMessageInfo

func (m *MeterMeasurement) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MeterMeasurement) GetMeasurement() int32 {
	if m != nil {
		return m.Measurement
	}
	return 0
}

type MeterData struct {
	MeterUuid            []byte              `protobuf:"bytes,1,opt,name=meterUuid,proto3" json:"meterUuid,omitempty"`
	Data                 []*MeterMeasurement `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeterData) Reset()         { *m = MeterData{} }
func (m *MeterData) String() string { return proto.CompactTextString(m) }
func (*MeterData) ProtoMessage()    {}
func (*MeterData) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{4}
}

func (m *MeterData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeterData.Unmarshal(m, b)
}
func (m *MeterData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeterData.Marshal(b, m, deterministic)
}
func (m *MeterData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeterData.Merge(m, src)
}
func (m *MeterData) XXX_Size() int {
	return xxx_messageInfo_MeterData.Size(m)
}
func (m *MeterData) XXX_DiscardUnknown() {
	xxx_messageInfo_MeterData.DiscardUnknown(m)
}

var xxx_messageInfo_MeterData proto.InternalMessageInfo

func (m *MeterData) GetMeterUuid() []byte {
	if m != nil {
		return m.MeterUuid
	}
	return nil
}

func (m *MeterData) GetData() []*MeterMeasurement {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetDataResponse struct {
	Action               GetDataResponse_InitResponseEnum `protobuf:"varint,1,opt,name=action,proto3,enum=org.demo.queryMeter.GetDataResponse_InitResponseEnum" json:"action,omitempty"`
	StreamUid            string                           `protobuf:"bytes,2,opt,name=streamUid,proto3" json:"streamUid,omitempty"`
	Offset               int32                            `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Data                 []*MeterData                     `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{5}
}

func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetAction() GetDataResponse_InitResponseEnum {
	if m != nil {
		return m.Action
	}
	return GetDataResponse_CONTINUE
}

func (m *GetDataResponse) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

func (m *GetDataResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetDataResponse) GetData() []*MeterData {
	if m != nil {
		return m.Data
	}
	return nil
}

// tag::AckDataRequest[]
type AckDataRequest struct {
	CursorUid            string   `protobuf:"bytes,1,opt,name=cursorUid,proto3" json:"cursorUid,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	BackpressureDelta    int32    `protobuf:"varint,3,opt,name=backpressureDelta,proto3" json:"backpressureDelta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckDataRequest) Reset()         { *m = AckDataRequest{} }
func (m *AckDataRequest) String() string { return proto.CompactTextString(m) }
func (*AckDataRequest) ProtoMessage()    {}
func (*AckDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{6}
}

func (m *AckDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckDataRequest.Unmarshal(m, b)
}
func (m *AckDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckDataRequest.Marshal(b, m, deterministic)
}
func (m *AckDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckDataRequest.Merge(m, src)
}
func (m *AckDataRequest) XXX_Size() int {
	return xxx_messageInfo_AckDataRequest.Size(m)
}
func (m *AckDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AckDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AckDataRequest proto.InternalMessageInfo

func (m *AckDataRequest) GetCursorUid() string {
	if m != nil {
		return m.CursorUid
	}
	return ""
}

func (m *AckDataRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *AckDataRequest) GetBackpressureDelta() int32 {
	if m != nil {
		return m.BackpressureDelta
	}
	return 0
}

type AckDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckDataResponse) Reset()         { *m = AckDataResponse{} }
func (m *AckDataResponse) String() string { return proto.CompactTextString(m) }
func (*AckDataResponse) ProtoMessage()    {}
func (*AckDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{7}
}

func (m *AckDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckDataResponse.Unmarshal(m, b)
}
func (m *AckDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckDataResponse.Marshal(b, m, deterministic)
}
func (m *AckDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckDataResponse.Merge(m, src)
}
func (m *AckDataResponse) XXX_Size() int {
	return xxx_messageInfo_AckDataResponse.Size(m)
}
func (m *AckDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AckDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AckDataResponse proto.InternalMessageInfo

// tag::CompleteRequest[]
type CompleteRequest struct {
	MonitoringData       string   `protobuf:"bytes,1,opt,name=monitoringData,proto3" json:"monitoringData,omitempty"`
	StreamUid            string   `protobuf:"bytes,2,opt,name=streamUid,proto3" json:"streamUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteRequest) Reset()         { *m = CompleteRequest{} }
func (m *CompleteRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteRequest) ProtoMessage()    {}
func (*CompleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{8}
}

func (m *CompleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteRequest.Unmarshal(m, b)
}
func (m *CompleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteRequest.Marshal(b, m, deterministic)
}
func (m *CompleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteRequest.Merge(m, src)
}
func (m *CompleteRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteRequest.Size(m)
}
func (m *CompleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteRequest proto.InternalMessageInfo

func (m *CompleteRequest) GetMonitoringData() string {
	if m != nil {
		return m.MonitoringData
	}
	return ""
}

func (m *CompleteRequest) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

type CompleteResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteResponse) Reset()         { *m = CompleteResponse{} }
func (m *CompleteResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteResponse) ProtoMessage()    {}
func (*CompleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ae58230653d403, []int{9}
}

func (m *CompleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteResponse.Unmarshal(m, b)
}
func (m *CompleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteResponse.Marshal(b, m, deterministic)
}
func (m *CompleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteResponse.Merge(m, src)
}
func (m *CompleteResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteResponse.Size(m)
}
func (m *CompleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteResponse proto.InternalMessageInfo

func (m *CompleteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CompleteResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("org.demo.queryMeter.QueryResponse_InitResponseEnum", QueryResponse_InitResponseEnum_name, QueryResponse_InitResponseEnum_value)
	proto.RegisterEnum("org.demo.queryMeter.GetDataResponse_InitResponseEnum", GetDataResponse_InitResponseEnum_name, GetDataResponse_InitResponseEnum_value)
	proto.RegisterType((*QueryRequest)(nil), "org.demo.queryMeter.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "org.demo.queryMeter.QueryResponse")
	proto.RegisterType((*GetDataRequest)(nil), "org.demo.queryMeter.GetDataRequest")
	proto.RegisterType((*MeterMeasurement)(nil), "org.demo.queryMeter.MeterMeasurement")
	proto.RegisterType((*MeterData)(nil), "org.demo.queryMeter.MeterData")
	proto.RegisterType((*GetDataResponse)(nil), "org.demo.queryMeter.GetDataResponse")
	proto.RegisterType((*AckDataRequest)(nil), "org.demo.queryMeter.AckDataRequest")
	proto.RegisterType((*AckDataResponse)(nil), "org.demo.queryMeter.AckDataResponse")
	proto.RegisterType((*CompleteRequest)(nil), "org.demo.queryMeter.CompleteRequest")
	proto.RegisterType((*CompleteResponse)(nil), "org.demo.queryMeter.CompleteResponse")
}

func init() { proto.RegisterFile("queryMeter.proto", fileDescriptor_82ae58230653d403) }

var fileDescriptor_82ae58230653d403 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0xf6, 0xa0, 0xa5, 0x30, 0x6d, 0xe1, 0xba, 0x26, 0x84, 0x10, 0x63, 0xf0, 0xb4, 0x86, 0x07,
	0x73, 0x1a, 0x1a, 0x1f, 0x4c, 0x7c, 0x31, 0xf4, 0x62, 0x1a, 0x05, 0xd2, 0x13, 0xda, 0x47, 0xdd,
	0xde, 0x2d, 0xe4, 0x52, 0xf6, 0xf6, 0xba, 0xbb, 0x67, 0xe2, 0x5f, 0xf4, 0xe7, 0xf8, 0xea, 0x8b,
	0xd9, 0xbd, 0x85, 0x83, 0x13, 0x68, 0xf5, 0x8d, 0x99, 0xd9, 0xf9, 0xe6, 0x9b, 0xf9, 0xbe, 0x1c,
	0x60, 0xdf, 0xa5, 0x84, 0xff, 0x18, 0x10, 0x49, 0xb8, 0x9b, 0x70, 0x26, 0x19, 0x7a, 0xcc, 0xf8,
	0xcc, 0x0d, 0x09, 0x65, 0x6e, 0x5e, 0x72, 0xbe, 0xc1, 0xd1, 0xa5, 0x8a, 0x7c, 0x72, 0x97, 0x12,
	0x21, 0xd1, 0x13, 0xa8, 0x51, 0x55, 0x98, 0xa4, 0x51, 0xd8, 0xb2, 0x3a, 0xe5, 0xee, 0x91, 0x9f,
	0x27, 0x50, 0x1b, 0xaa, 0x53, 0xce, 0xe8, 0x38, 0xa2, 0xa4, 0x55, 0xea, 0x58, 0xdd, 0x7d, 0x7f,
	0x19, 0xa3, 0x26, 0x54, 0x24, 0xd3, 0x95, 0xb2, 0xae, 0x98, 0xc8, 0xf9, 0x69, 0xc1, 0xb1, 0x19,
	0x21, 0x12, 0x16, 0x0b, 0x82, 0x3e, 0x41, 0x05, 0x07, 0x32, 0x62, 0x71, 0xcb, 0xea, 0x58, 0xdd,
	0x7a, 0xef, 0xcc, 0xdd, 0xc0, 0xcc, 0x5d, 0xeb, 0x71, 0x2f, 0xe2, 0x48, 0x2e, 0x02, 0x2f, 0x4e,
	0xa9, 0x6f, 0x20, 0x14, 0x61, 0x21, 0x39, 0xc1, 0x74, 0x12, 0x85, 0x9a, 0x53, 0xcd, 0xcf, 0x13,
	0xc8, 0x86, 0x72, 0xca, 0xe7, 0x9a, 0x51, 0xcd, 0x57, 0x3f, 0x9d, 0xf7, 0x60, 0x17, 0xb1, 0xd0,
	0x11, 0x54, 0xfb, 0xa3, 0xe1, 0xf8, 0x62, 0x38, 0xf1, 0xec, 0x47, 0xe8, 0x18, 0x6a, 0xbe, 0xd7,
	0x1f, 0x0d, 0x87, 0x5e, 0x7f, 0x6c, 0x5b, 0xe8, 0x10, 0x0e, 0x06, 0xa3, 0x2b, 0xef, 0xeb, 0x78,
	0x64, 0x97, 0x1c, 0x09, 0xf5, 0x8f, 0x44, 0x9e, 0x63, 0x89, 0x57, 0x0e, 0x96, 0xcf, 0xb7, 0x8a,
	0xf3, 0x9b, 0x50, 0x61, 0xd3, 0xa9, 0x20, 0xd2, 0x9c, 0xcb, 0x44, 0xe8, 0x15, 0x9c, 0xdc, 0xe0,
	0xe0, 0x36, 0xe1, 0x44, 0x88, 0x94, 0x93, 0x3e, 0x4b, 0x63, 0x69, 0xee, 0xf6, 0x77, 0xc1, 0xf1,
	0xc1, 0xd6, 0x37, 0x19, 0x10, 0xac, 0x92, 0x94, 0xc4, 0x7a, 0xae, 0x8c, 0x28, 0x11, 0x12, 0xd3,
	0x44, 0xcf, 0xdd, 0xf7, 0xf3, 0x04, 0xea, 0xc0, 0x21, 0xcd, 0x1f, 0x9b, 0xe1, 0xab, 0x29, 0x27,
	0x84, 0x9a, 0xc6, 0x54, 0xbb, 0x14, 0x55, 0xb7, 0xd6, 0x55, 0x7f, 0x07, 0x7b, 0x21, 0x96, 0xb8,
	0x55, 0xea, 0x94, 0xbb, 0x87, 0xbd, 0xd3, 0x8d, 0x6a, 0x15, 0xf9, 0xf9, 0xba, 0xc5, 0xf9, 0x6d,
	0x41, 0x63, 0x79, 0x30, 0x23, 0xff, 0xa0, 0x20, 0xff, 0xdb, 0x8d, 0x80, 0x85, 0xae, 0xff, 0x35,
	0x40, 0x2e, 0x40, 0x79, 0x4d, 0x80, 0x9e, 0xd9, 0x69, 0x4f, 0xef, 0xf4, 0x74, 0xfb, 0x4e, 0x9a,
	0x44, 0xb6, 0xcc, 0xeb, 0x7f, 0xb4, 0x8e, 0x72, 0xcb, 0x87, 0xe0, 0xb6, 0xe0, 0x96, 0x20, 0xe5,
	0x82, 0xf1, 0x15, 0xb7, 0x2c, 0x13, 0x0f, 0x75, 0xcb, 0x39, 0x99, 0x4b, 0xbc, 0xc9, 0x2d, 0xba,
	0xe0, 0x9c, 0x40, 0x63, 0x39, 0x35, 0x63, 0xea, 0x5c, 0x43, 0xa3, 0xcf, 0x68, 0x32, 0x27, 0x92,
	0x2c, 0x98, 0xbc, 0x84, 0x3a, 0x65, 0x71, 0x24, 0x19, 0x8f, 0xe2, 0x99, 0x7a, 0x6c, 0xe8, 0x14,
	0xb2, 0xbb, 0xcf, 0xeb, 0x7c, 0x06, 0x3b, 0x07, 0x36, 0xfa, 0x36, 0xa1, 0xc2, 0x89, 0x48, 0xe7,
	0x52, 0x23, 0x56, 0x7d, 0x13, 0x29, 0x4f, 0x86, 0x44, 0x04, 0x3c, 0x4a, 0xb4, 0xf8, 0x19, 0xd6,
	0x6a, 0xaa, 0xf7, 0xab, 0x04, 0x6d, 0x8f, 0x06, 0xf8, 0x72, 0xa9, 0x81, 0x50, 0x1c, 0xbe, 0x10,
	0xfe, 0x3d, 0x0a, 0x94, 0x71, 0xf6, 0xd4, 0xfd, 0xd1, 0xb3, 0x5d, 0xdf, 0x0b, 0xbd, 0x5d, 0xdb,
	0xb9, 0xff, 0x93, 0x82, 0xae, 0xe0, 0x60, 0x96, 0x99, 0x0c, 0x3d, 0xdf, 0x6d, 0xc1, 0x0c, 0xf3,
	0xc5, 0x43, 0x7c, 0xfa, 0xc6, 0x42, 0x63, 0x38, 0xc0, 0xd9, 0xfd, 0xb7, 0xe0, 0xae, 0x7b, 0x62,
	0x0b, 0x6e, 0x41, 0x42, 0x74, 0x0d, 0xd5, 0xc5, 0xa5, 0xd1, 0xe6, 0x8e, 0x82, 0xc2, 0xed, 0xd3,
	0x7b, 0x5e, 0x65, 0xc0, 0x37, 0x15, 0xfd, 0xef, 0x70, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0x30, 0xac, 0xe6, 0x31, 0x06, 0x00, 0x00,
}
