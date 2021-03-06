// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deals.proto

package filecoin_deals_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DealConfig struct {
	Miner                string   `protobuf:"bytes,1,opt,name=miner,proto3" json:"miner,omitempty"`
	EpochPrice           uint64   `protobuf:"varint,2,opt,name=epochPrice,proto3" json:"epochPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealConfig) Reset()         { *m = DealConfig{} }
func (m *DealConfig) String() string { return proto.CompactTextString(m) }
func (*DealConfig) ProtoMessage()    {}
func (*DealConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{0}
}

func (m *DealConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealConfig.Unmarshal(m, b)
}
func (m *DealConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealConfig.Marshal(b, m, deterministic)
}
func (m *DealConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealConfig.Merge(m, src)
}
func (m *DealConfig) XXX_Size() int {
	return xxx_messageInfo_DealConfig.Size(m)
}
func (m *DealConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DealConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DealConfig proto.InternalMessageInfo

func (m *DealConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *DealConfig) GetEpochPrice() uint64 {
	if m != nil {
		return m.EpochPrice
	}
	return 0
}

type DealInfo struct {
	ProposalCid          string   `protobuf:"bytes,1,opt,name=proposalCid,proto3" json:"proposalCid,omitempty"`
	StateID              uint64   `protobuf:"varint,2,opt,name=stateID,proto3" json:"stateID,omitempty"`
	StateName            string   `protobuf:"bytes,3,opt,name=stateName,proto3" json:"stateName,omitempty"`
	Miner                string   `protobuf:"bytes,4,opt,name=miner,proto3" json:"miner,omitempty"`
	PieceRef             []byte   `protobuf:"bytes,5,opt,name=pieceRef,proto3" json:"pieceRef,omitempty"`
	Size                 uint64   `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	PricePerEpoch        uint64   `protobuf:"varint,7,opt,name=pricePerEpoch,proto3" json:"pricePerEpoch,omitempty"`
	Duration             uint64   `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealInfo) Reset()         { *m = DealInfo{} }
func (m *DealInfo) String() string { return proto.CompactTextString(m) }
func (*DealInfo) ProtoMessage()    {}
func (*DealInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{1}
}

func (m *DealInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealInfo.Unmarshal(m, b)
}
func (m *DealInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealInfo.Marshal(b, m, deterministic)
}
func (m *DealInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealInfo.Merge(m, src)
}
func (m *DealInfo) XXX_Size() int {
	return xxx_messageInfo_DealInfo.Size(m)
}
func (m *DealInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DealInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DealInfo proto.InternalMessageInfo

func (m *DealInfo) GetProposalCid() string {
	if m != nil {
		return m.ProposalCid
	}
	return ""
}

func (m *DealInfo) GetStateID() uint64 {
	if m != nil {
		return m.StateID
	}
	return 0
}

func (m *DealInfo) GetStateName() string {
	if m != nil {
		return m.StateName
	}
	return ""
}

func (m *DealInfo) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *DealInfo) GetPieceRef() []byte {
	if m != nil {
		return m.PieceRef
	}
	return nil
}

func (m *DealInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DealInfo) GetPricePerEpoch() uint64 {
	if m != nil {
		return m.PricePerEpoch
	}
	return 0
}

func (m *DealInfo) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type StoreParams struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DealConfigs          []*DealConfig `protobuf:"bytes,2,rep,name=dealConfigs,proto3" json:"dealConfigs,omitempty"`
	Duration             uint64        `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreParams) Reset()         { *m = StoreParams{} }
func (m *StoreParams) String() string { return proto.CompactTextString(m) }
func (*StoreParams) ProtoMessage()    {}
func (*StoreParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{2}
}

func (m *StoreParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreParams.Unmarshal(m, b)
}
func (m *StoreParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreParams.Marshal(b, m, deterministic)
}
func (m *StoreParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreParams.Merge(m, src)
}
func (m *StoreParams) XXX_Size() int {
	return xxx_messageInfo_StoreParams.Size(m)
}
func (m *StoreParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreParams.DiscardUnknown(m)
}

var xxx_messageInfo_StoreParams proto.InternalMessageInfo

func (m *StoreParams) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StoreParams) GetDealConfigs() []*DealConfig {
	if m != nil {
		return m.DealConfigs
	}
	return nil
}

func (m *StoreParams) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type StoreRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*StoreRequest_StoreParams
	//	*StoreRequest_Chunk
	Payload              isStoreRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{3}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

type isStoreRequest_Payload interface {
	isStoreRequest_Payload()
}

type StoreRequest_StoreParams struct {
	StoreParams *StoreParams `protobuf:"bytes,1,opt,name=storeParams,proto3,oneof"`
}

type StoreRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*StoreRequest_StoreParams) isStoreRequest_Payload() {}

func (*StoreRequest_Chunk) isStoreRequest_Payload() {}

func (m *StoreRequest) GetPayload() isStoreRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StoreRequest) GetStoreParams() *StoreParams {
	if x, ok := m.GetPayload().(*StoreRequest_StoreParams); ok {
		return x.StoreParams
	}
	return nil
}

func (m *StoreRequest) GetChunk() []byte {
	if x, ok := m.GetPayload().(*StoreRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StoreRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StoreRequest_StoreParams)(nil),
		(*StoreRequest_Chunk)(nil),
	}
}

type StoreReply struct {
	Cids                 []string      `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	FailedDeals          []*DealConfig `protobuf:"bytes,2,rep,name=failedDeals,proto3" json:"failedDeals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreReply) Reset()         { *m = StoreReply{} }
func (m *StoreReply) String() string { return proto.CompactTextString(m) }
func (*StoreReply) ProtoMessage()    {}
func (*StoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{4}
}

func (m *StoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReply.Unmarshal(m, b)
}
func (m *StoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReply.Marshal(b, m, deterministic)
}
func (m *StoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReply.Merge(m, src)
}
func (m *StoreReply) XXX_Size() int {
	return xxx_messageInfo_StoreReply.Size(m)
}
func (m *StoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReply proto.InternalMessageInfo

func (m *StoreReply) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

func (m *StoreReply) GetFailedDeals() []*DealConfig {
	if m != nil {
		return m.FailedDeals
	}
	return nil
}

type WatchRequest struct {
	Proposals            []string `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{5}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetProposals() []string {
	if m != nil {
		return m.Proposals
	}
	return nil
}

type WatchReply struct {
	DealInfo             *DealInfo `protobuf:"bytes,1,opt,name=dealInfo,proto3" json:"dealInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WatchReply) Reset()         { *m = WatchReply{} }
func (m *WatchReply) String() string { return proto.CompactTextString(m) }
func (*WatchReply) ProtoMessage()    {}
func (*WatchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71783c876a92172d, []int{6}
}

func (m *WatchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchReply.Unmarshal(m, b)
}
func (m *WatchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchReply.Marshal(b, m, deterministic)
}
func (m *WatchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchReply.Merge(m, src)
}
func (m *WatchReply) XXX_Size() int {
	return xxx_messageInfo_WatchReply.Size(m)
}
func (m *WatchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchReply.DiscardUnknown(m)
}

var xxx_messageInfo_WatchReply proto.InternalMessageInfo

func (m *WatchReply) GetDealInfo() *DealInfo {
	if m != nil {
		return m.DealInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*DealConfig)(nil), "filecoin.deals.pb.DealConfig")
	proto.RegisterType((*DealInfo)(nil), "filecoin.deals.pb.DealInfo")
	proto.RegisterType((*StoreParams)(nil), "filecoin.deals.pb.StoreParams")
	proto.RegisterType((*StoreRequest)(nil), "filecoin.deals.pb.StoreRequest")
	proto.RegisterType((*StoreReply)(nil), "filecoin.deals.pb.StoreReply")
	proto.RegisterType((*WatchRequest)(nil), "filecoin.deals.pb.WatchRequest")
	proto.RegisterType((*WatchReply)(nil), "filecoin.deals.pb.WatchReply")
}

func init() { proto.RegisterFile("deals.proto", fileDescriptor_71783c876a92172d) }

var fileDescriptor_71783c876a92172d = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xed, 0x6c, 0xda, 0x6d, 0x7b, 0xd3, 0x7d, 0x70, 0x10, 0x09, 0x6b, 0x77, 0x2d, 0xc1, 0x87,
	0x3c, 0x48, 0x90, 0x8a, 0xf8, 0x28, 0x66, 0xb7, 0xb2, 0x45, 0x90, 0x30, 0x2e, 0xf8, 0x3c, 0x9b,
	0xdc, 0xd8, 0xc1, 0x34, 0x13, 0x33, 0x29, 0x58, 0x9f, 0xfd, 0x0f, 0xc1, 0xef, 0xf3, 0x23, 0x64,
	0xa6, 0x93, 0x26, 0xeb, 0xb6, 0xe0, 0xdb, 0xdc, 0x93, 0x7b, 0xcf, 0x3d, 0xe7, 0xcc, 0x04, 0xdc,
	0x14, 0x79, 0xae, 0xc2, 0xb2, 0x92, 0xb5, 0xa4, 0x8f, 0x32, 0x91, 0x63, 0x22, 0x45, 0x11, 0x5a,
	0xf4, 0xce, 0x8f, 0x00, 0xae, 0x91, 0xe7, 0x57, 0xb2, 0xc8, 0xc4, 0x17, 0xfa, 0x18, 0x06, 0x6b,
	0x51, 0x60, 0xe5, 0x91, 0x19, 0x09, 0xc6, 0x6c, 0x57, 0xd0, 0x4b, 0x00, 0x2c, 0x65, 0xb2, 0x8a,
	0x2b, 0x91, 0xa0, 0x77, 0x32, 0x23, 0x41, 0x9f, 0x75, 0x10, 0xff, 0x0f, 0x81, 0x91, 0x26, 0x59,
	0x16, 0x99, 0xa4, 0x33, 0x70, 0xcb, 0x4a, 0x96, 0x52, 0xf1, 0xfc, 0x4a, 0xa4, 0x96, 0xa8, 0x0b,
	0x51, 0x0f, 0x86, 0xaa, 0xe6, 0x35, 0x2e, 0xaf, 0x2d, 0x57, 0x53, 0xd2, 0x29, 0x8c, 0xcd, 0xf1,
	0x23, 0x5f, 0xa3, 0xe7, 0x98, 0xc9, 0x16, 0x68, 0xc5, 0xf5, 0xbb, 0xe2, 0xce, 0x61, 0x54, 0x0a,
	0x4c, 0x90, 0x61, 0xe6, 0x0d, 0x66, 0x24, 0x98, 0xb0, 0x7d, 0x4d, 0x29, 0xf4, 0x95, 0xf8, 0x81,
	0xde, 0xa9, 0x59, 0x63, 0xce, 0xf4, 0x39, 0x9c, 0x95, 0x5a, 0x75, 0x8c, 0xd5, 0x42, 0x5b, 0xf0,
	0x86, 0xe6, 0xe3, 0x7d, 0x50, 0xb3, 0xa6, 0x9b, 0x8a, 0xd7, 0x42, 0x16, 0xde, 0xc8, 0x34, 0xec,
	0x6b, 0xff, 0x27, 0x01, 0xf7, 0x53, 0x2d, 0x2b, 0x8c, 0x79, 0xc5, 0xd7, 0x4a, 0xfb, 0xe1, 0x69,
	0x5a, 0xa1, 0x52, 0xd6, 0x6d, 0x53, 0xd2, 0xb7, 0xbb, 0xf8, 0x77, 0xe1, 0x2a, 0xef, 0x64, 0xe6,
	0x04, 0xee, 0xfc, 0x22, 0x7c, 0x70, 0x0b, 0x61, 0x7b, 0x05, 0xac, 0x3b, 0x71, 0x4f, 0x86, 0xf3,
	0x8f, 0x8c, 0x0d, 0x4c, 0x8c, 0x0a, 0x86, 0xdf, 0x36, 0xa8, 0x6a, 0x1a, 0x81, 0xab, 0x5a, 0x55,
	0x46, 0x8a, 0x3b, 0xbf, 0x3c, 0xb0, 0xac, 0xa3, 0xfd, 0xa6, 0xc7, 0xba, 0x43, 0xf4, 0x09, 0x0c,
	0x92, 0xd5, 0xa6, 0xf8, 0x6a, 0x2e, 0x66, 0x72, 0xd3, 0x63, 0xbb, 0x32, 0x1a, 0xc3, 0xb0, 0xe4,
	0xdb, 0x5c, 0xf2, 0xd4, 0xe7, 0x00, 0x76, 0x6d, 0x99, 0x6f, 0x75, 0xc2, 0x89, 0x48, 0xf5, 0x36,
	0x27, 0x18, 0x33, 0x73, 0xd6, 0xae, 0x33, 0x2e, 0x72, 0x4c, 0xb5, 0xab, 0xff, 0x75, 0xdd, 0x99,
	0xf0, 0x5f, 0xc0, 0xe4, 0x33, 0xaf, 0x93, 0x55, 0xe3, 0x6c, 0x0a, 0xe3, 0xe6, 0xfd, 0x34, 0x9b,
	0x5a, 0xc0, 0x5f, 0x00, 0xd8, 0x6e, 0x2d, 0xe8, 0x0d, 0x8c, 0x52, 0xfb, 0x14, 0x6d, 0x04, 0x4f,
	0x8f, 0x6c, 0xd6, 0x2d, 0x6c, 0xdf, 0x3c, 0xff, 0x45, 0xc0, 0x79, 0x17, 0x2f, 0xe9, 0x07, 0x18,
	0x18, 0x7f, 0xf4, 0xd9, 0xb1, 0xe8, 0xac, 0xac, 0xf3, 0x8b, 0xe3, 0x0d, 0x65, 0xbe, 0xf5, 0x7b,
	0x01, 0xd1, 0x64, 0x46, 0xdb, 0x41, 0xb2, 0xae, 0xc7, 0x83, 0x64, 0xad, 0x2d, 0xbf, 0xf7, 0x92,
	0x44, 0xaf, 0x61, 0x2a, 0x64, 0x58, 0xe3, 0xf7, 0x5a, 0xe4, 0xf8, 0xb0, 0x3d, 0x3a, 0x7b, 0x6f,
	0x21, 0x93, 0x62, 0x4c, 0x7e, 0x9f, 0x38, 0xb7, 0xb7, 0x8b, 0xbb, 0x53, 0xf3, 0xef, 0xbf, 0xfa,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x93, 0xa5, 0x8d, 0x6a, 0x0a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Store(ctx context.Context, opts ...grpc.CallOption) (API_StoreClient, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (API_WatchClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Store(ctx context.Context, opts ...grpc.CallOption) (API_StoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/filecoin.deals.pb.API/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIStoreClient{stream}
	return x, nil
}

type API_StoreClient interface {
	Send(*StoreRequest) error
	CloseAndRecv() (*StoreReply, error)
	grpc.ClientStream
}

type aPIStoreClient struct {
	grpc.ClientStream
}

func (x *aPIStoreClient) Send(m *StoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIStoreClient) CloseAndRecv() (*StoreReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (API_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/filecoin.deals.pb.API/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_WatchClient interface {
	Recv() (*WatchReply, error)
	grpc.ClientStream
}

type aPIWatchClient struct {
	grpc.ClientStream
}

func (x *aPIWatchClient) Recv() (*WatchReply, error) {
	m := new(WatchReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Store(API_StoreServer) error
	Watch(*WatchRequest, API_WatchServer) error
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Store(srv API_StoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedAPIServer) Watch(req *WatchRequest, srv API_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).Store(&aPIStoreServer{stream})
}

type API_StoreServer interface {
	SendAndClose(*StoreReply) error
	Recv() (*StoreRequest, error)
	grpc.ServerStream
}

type aPIStoreServer struct {
	grpc.ServerStream
}

func (x *aPIStoreServer) SendAndClose(m *StoreReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIStoreServer) Recv() (*StoreRequest, error) {
	m := new(StoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Watch(m, &aPIWatchServer{stream})
}

type API_WatchServer interface {
	Send(*WatchReply) error
	grpc.ServerStream
}

type aPIWatchServer struct {
	grpc.ServerStream
}

func (x *aPIWatchServer) Send(m *WatchReply) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filecoin.deals.pb.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _API_Store_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _API_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deals.proto",
}
