// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/campaign_feed_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [CampaignFeedService.GetCampaignFeed][google.ads.googleads.v0.services.CampaignFeedService.GetCampaignFeed].
type GetCampaignFeedRequest struct {
	// The resource name of the campaign feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignFeedRequest) Reset()         { *m = GetCampaignFeedRequest{} }
func (m *GetCampaignFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignFeedRequest) ProtoMessage()    {}
func (*GetCampaignFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c, []int{0}
}
func (m *GetCampaignFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignFeedRequest.Unmarshal(m, b)
}
func (m *GetCampaignFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignFeedRequest.Marshal(b, m, deterministic)
}
func (dst *GetCampaignFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignFeedRequest.Merge(dst, src)
}
func (m *GetCampaignFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignFeedRequest.Size(m)
}
func (m *GetCampaignFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignFeedRequest proto.InternalMessageInfo

func (m *GetCampaignFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CampaignFeedService.MutateCampaignFeeds][google.ads.googleads.v0.services.CampaignFeedService.MutateCampaignFeeds].
type MutateCampaignFeedsRequest struct {
	// The ID of the customer whose campaign feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign feeds.
	Operations []*CampaignFeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignFeedsRequest) Reset()         { *m = MutateCampaignFeedsRequest{} }
func (m *MutateCampaignFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedsRequest) ProtoMessage()    {}
func (*MutateCampaignFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c, []int{1}
}
func (m *MutateCampaignFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedsRequest.Merge(dst, src)
}
func (m *MutateCampaignFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Size(m)
}
func (m *MutateCampaignFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedsRequest proto.InternalMessageInfo

func (m *MutateCampaignFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignFeedsRequest) GetOperations() []*CampaignFeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign feed.
type CampaignFeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignFeedOperation_Create
	//	*CampaignFeedOperation_Update
	//	*CampaignFeedOperation_Remove
	Operation            isCampaignFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CampaignFeedOperation) Reset()         { *m = CampaignFeedOperation{} }
func (m *CampaignFeedOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignFeedOperation) ProtoMessage()    {}
func (*CampaignFeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c, []int{2}
}
func (m *CampaignFeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeedOperation.Unmarshal(m, b)
}
func (m *CampaignFeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeedOperation.Marshal(b, m, deterministic)
}
func (dst *CampaignFeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeedOperation.Merge(dst, src)
}
func (m *CampaignFeedOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignFeedOperation.Size(m)
}
func (m *CampaignFeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeedOperation proto.InternalMessageInfo

func (m *CampaignFeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignFeedOperation_Operation interface {
	isCampaignFeedOperation_Operation()
}

type CampaignFeedOperation_Create struct {
	Create *resources.CampaignFeed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignFeedOperation_Update struct {
	Update *resources.CampaignFeed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignFeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignFeedOperation_Create) isCampaignFeedOperation_Operation() {}

func (*CampaignFeedOperation_Update) isCampaignFeedOperation_Operation() {}

func (*CampaignFeedOperation_Remove) isCampaignFeedOperation_Operation() {}

func (m *CampaignFeedOperation) GetOperation() isCampaignFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignFeedOperation) GetCreate() *resources.CampaignFeed {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignFeedOperation) GetUpdate() *resources.CampaignFeed {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignFeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CampaignFeedOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CampaignFeedOperation_OneofMarshaler, _CampaignFeedOperation_OneofUnmarshaler, _CampaignFeedOperation_OneofSizer, []interface{}{
		(*CampaignFeedOperation_Create)(nil),
		(*CampaignFeedOperation_Update)(nil),
		(*CampaignFeedOperation_Remove)(nil),
	}
}

func _CampaignFeedOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CampaignFeedOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CampaignFeedOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *CampaignFeedOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *CampaignFeedOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("CampaignFeedOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _CampaignFeedOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CampaignFeedOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CampaignFeed)
		err := b.DecodeMessage(msg)
		m.Operation = &CampaignFeedOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CampaignFeed)
		err := b.DecodeMessage(msg)
		m.Operation = &CampaignFeedOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &CampaignFeedOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _CampaignFeedOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CampaignFeedOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CampaignFeedOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignFeedOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignFeedOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a campaign feed mutate.
type MutateCampaignFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateCampaignFeedsResponse) Reset()         { *m = MutateCampaignFeedsResponse{} }
func (m *MutateCampaignFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedsResponse) ProtoMessage()    {}
func (*MutateCampaignFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c, []int{3}
}
func (m *MutateCampaignFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedsResponse.Merge(dst, src)
}
func (m *MutateCampaignFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Size(m)
}
func (m *MutateCampaignFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedsResponse proto.InternalMessageInfo

func (m *MutateCampaignFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignFeedsResponse) GetResults() []*MutateCampaignFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign feed mutate.
type MutateCampaignFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignFeedResult) Reset()         { *m = MutateCampaignFeedResult{} }
func (m *MutateCampaignFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedResult) ProtoMessage()    {}
func (*MutateCampaignFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c, []int{4}
}
func (m *MutateCampaignFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedResult.Unmarshal(m, b)
}
func (m *MutateCampaignFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedResult.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedResult.Merge(dst, src)
}
func (m *MutateCampaignFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedResult.Size(m)
}
func (m *MutateCampaignFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedResult proto.InternalMessageInfo

func (m *MutateCampaignFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignFeedRequest)(nil), "google.ads.googleads.v0.services.GetCampaignFeedRequest")
	proto.RegisterType((*MutateCampaignFeedsRequest)(nil), "google.ads.googleads.v0.services.MutateCampaignFeedsRequest")
	proto.RegisterType((*CampaignFeedOperation)(nil), "google.ads.googleads.v0.services.CampaignFeedOperation")
	proto.RegisterType((*MutateCampaignFeedsResponse)(nil), "google.ads.googleads.v0.services.MutateCampaignFeedsResponse")
	proto.RegisterType((*MutateCampaignFeedResult)(nil), "google.ads.googleads.v0.services.MutateCampaignFeedResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignFeedServiceClient is the client API for CampaignFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignFeedServiceClient interface {
	// Returns the requested campaign feed in full detail.
	GetCampaignFeed(ctx context.Context, in *GetCampaignFeedRequest, opts ...grpc.CallOption) (*resources.CampaignFeed, error)
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error)
}

type campaignFeedServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignFeedServiceClient(cc *grpc.ClientConn) CampaignFeedServiceClient {
	return &campaignFeedServiceClient{cc}
}

func (c *campaignFeedServiceClient) GetCampaignFeed(ctx context.Context, in *GetCampaignFeedRequest, opts ...grpc.CallOption) (*resources.CampaignFeed, error) {
	out := new(resources.CampaignFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CampaignFeedService/GetCampaignFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignFeedServiceClient) MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error) {
	out := new(MutateCampaignFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CampaignFeedService/MutateCampaignFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignFeedServiceServer is the server API for CampaignFeedService service.
type CampaignFeedServiceServer interface {
	// Returns the requested campaign feed in full detail.
	GetCampaignFeed(context.Context, *GetCampaignFeedRequest) (*resources.CampaignFeed, error)
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	MutateCampaignFeeds(context.Context, *MutateCampaignFeedsRequest) (*MutateCampaignFeedsResponse, error)
}

func RegisterCampaignFeedServiceServer(s *grpc.Server, srv CampaignFeedServiceServer) {
	s.RegisterService(&_CampaignFeedService_serviceDesc, srv)
}

func _CampaignFeedService_GetCampaignFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignFeedServiceServer).GetCampaignFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CampaignFeedService/GetCampaignFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignFeedServiceServer).GetCampaignFeed(ctx, req.(*GetCampaignFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignFeedService_MutateCampaignFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CampaignFeedService/MutateCampaignFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, req.(*MutateCampaignFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CampaignFeedService",
	HandlerType: (*CampaignFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignFeed",
			Handler:    _CampaignFeedService_GetCampaignFeed_Handler,
		},
		{
			MethodName: "MutateCampaignFeeds",
			Handler:    _CampaignFeedService_MutateCampaignFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/campaign_feed_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/campaign_feed_service.proto", fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c)
}

var fileDescriptor_campaign_feed_service_a5d2ffb834d3ed7c = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x0e, 0x2a, 0x74, 0x5d, 0xa8, 0xb4, 0x55, 0xc1, 0x0a, 0x08, 0x22, 0x53, 0x89, 0x2a,
	0x07, 0x3b, 0x04, 0xa1, 0x22, 0xb7, 0x11, 0x4a, 0x11, 0x69, 0x7b, 0x28, 0xad, 0x5c, 0x54, 0x24,
	0x14, 0xc9, 0xda, 0xda, 0x1b, 0xcb, 0xaa, 0xed, 0x35, 0xbb, 0xeb, 0xa0, 0xaa, 0xea, 0x85, 0x3f,
	0x40, 0x7c, 0x01, 0x1c, 0xb9, 0x73, 0xe2, 0x0f, 0xb8, 0x21, 0x7e, 0xa1, 0x27, 0x7e, 0x02, 0x64,
	0xaf, 0x37, 0x24, 0x6d, 0xa2, 0x40, 0x6e, 0xe3, 0xd9, 0x37, 0x6f, 0x66, 0xf6, 0xcd, 0x8e, 0xc1,
	0x46, 0x40, 0x48, 0x10, 0x61, 0x0b, 0xf9, 0xcc, 0x12, 0x66, 0x6e, 0xf5, 0x1b, 0x16, 0xc3, 0xb4,
	0x1f, 0x7a, 0x98, 0x59, 0x1e, 0x8a, 0x53, 0x14, 0x06, 0x89, 0xdb, 0xc3, 0xd8, 0x77, 0x4b, 0xb7,
	0x99, 0x52, 0xc2, 0x09, 0xac, 0x89, 0x10, 0x13, 0xf9, 0xcc, 0x1c, 0x44, 0x9b, 0xfd, 0x86, 0x29,
	0xa3, 0xab, 0x4f, 0x26, 0xf1, 0x53, 0xcc, 0x48, 0x46, 0x2f, 0x25, 0x10, 0xc4, 0xd5, 0xbb, 0x32,
	0x2c, 0x0d, 0x2d, 0x94, 0x24, 0x84, 0x23, 0x1e, 0x92, 0x84, 0x95, 0xa7, 0x65, 0x5a, 0xab, 0xf8,
	0x3a, 0xca, 0x7a, 0x56, 0x2f, 0xc4, 0x91, 0xef, 0xc6, 0x88, 0x1d, 0x97, 0x88, 0x7b, 0x17, 0x11,
	0xef, 0x28, 0x4a, 0x53, 0x4c, 0x25, 0xc3, 0xed, 0xf2, 0x9c, 0xa6, 0x9e, 0xc5, 0x38, 0xe2, 0x59,
	0x79, 0x60, 0xb4, 0xc0, 0xad, 0x2d, 0xcc, 0x9f, 0x97, 0x25, 0x75, 0x30, 0xf6, 0x1d, 0xfc, 0x36,
	0xc3, 0x8c, 0xc3, 0x07, 0xe0, 0x86, 0xac, 0xd9, 0x4d, 0x50, 0x8c, 0x75, 0xa5, 0xa6, 0xac, 0xce,
	0x3b, 0x0b, 0xd2, 0xf9, 0x12, 0xc5, 0xd8, 0x38, 0x57, 0x40, 0x75, 0x37, 0xe3, 0x88, 0xe3, 0x61,
	0x0a, 0x26, 0x39, 0xee, 0x03, 0xcd, 0xcb, 0x18, 0x27, 0x31, 0xa6, 0x6e, 0xe8, 0x97, 0x0c, 0x40,
	0xba, 0x76, 0x7c, 0xf8, 0x1a, 0x00, 0x92, 0x62, 0x2a, 0xba, 0xd5, 0xd5, 0x5a, 0x65, 0x55, 0x6b,
	0xae, 0x99, 0xd3, 0x6e, 0xd9, 0x1c, 0x4e, 0xb6, 0x27, 0xe3, 0x9d, 0x21, 0x2a, 0xf8, 0x10, 0x2c,
	0xa6, 0x88, 0xf2, 0x10, 0x45, 0x6e, 0x0f, 0x85, 0x51, 0x46, 0xb1, 0x5e, 0xa9, 0x29, 0xab, 0xd7,
	0x9d, 0x9b, 0xa5, 0xbb, 0x23, 0xbc, 0x79, 0x9b, 0x7d, 0x14, 0x85, 0x3e, 0xe2, 0xd8, 0x25, 0x49,
	0x74, 0xa2, 0x5f, 0x2d, 0x60, 0x0b, 0xd2, 0xb9, 0x97, 0x44, 0x27, 0xc6, 0x07, 0x15, 0x2c, 0x8f,
	0xcd, 0x09, 0xd7, 0x81, 0x96, 0xa5, 0x45, 0x70, 0xae, 0x46, 0x11, 0xac, 0x35, 0xab, 0xb2, 0x03,
	0x29, 0x87, 0xd9, 0xc9, 0x05, 0xdb, 0x45, 0xec, 0xd8, 0x01, 0x02, 0x9e, 0xdb, 0x70, 0x07, 0xcc,
	0x79, 0x14, 0x23, 0x2e, 0xee, 0x56, 0x6b, 0x5a, 0x13, 0x3b, 0x1f, 0x4c, 0xcf, 0x48, 0xeb, 0xdb,
	0x57, 0x9c, 0x92, 0x20, 0xa7, 0x12, 0xc4, 0xba, 0x3a, 0x33, 0x95, 0x20, 0x80, 0x3a, 0x98, 0xa3,
	0x38, 0x26, 0x7d, 0x71, 0x63, 0xf3, 0xf9, 0x89, 0xf8, 0xde, 0xd4, 0xc0, 0xfc, 0xe0, 0x8a, 0x8d,
	0x6f, 0x0a, 0xb8, 0x33, 0x56, 0x7a, 0x96, 0x92, 0x84, 0x61, 0xd8, 0x01, 0xcb, 0x17, 0x14, 0x70,
	0x31, 0xa5, 0x84, 0x16, 0xac, 0x5a, 0x13, 0xca, 0x02, 0x69, 0xea, 0x99, 0x07, 0xc5, 0x48, 0x3a,
	0x4b, 0xa3, 0xda, 0xbc, 0xc8, 0xe1, 0xf0, 0x15, 0xb8, 0x46, 0x31, 0xcb, 0x22, 0x2e, 0xe7, 0xc3,
	0x9e, 0x3e, 0x1f, 0x97, 0xeb, 0x72, 0x0a, 0x0a, 0x47, 0x52, 0x19, 0xcf, 0x80, 0x3e, 0x09, 0xf4,
	0x4f, 0x93, 0xdf, 0xfc, 0x54, 0x01, 0x4b, 0xc3, 0xb1, 0x07, 0x22, 0x37, 0xfc, 0xaa, 0x80, 0xc5,
	0x0b, 0x2f, 0x0a, 0x3e, 0x9d, 0x5e, 0xf1, 0xf8, 0x47, 0x58, 0xfd, 0x5f, 0x19, 0x8d, 0xb5, 0xf7,
	0x3f, 0xcf, 0x3f, 0xaa, 0x8f, 0xa0, 0x95, 0xef, 0x9c, 0xd3, 0x91, 0x36, 0x5a, 0xf2, 0xdd, 0x31,
	0xab, 0x3e, 0x58, 0x42, 0x85, 0x66, 0x56, 0xfd, 0x0c, 0xfe, 0x50, 0xc0, 0xd2, 0x18, 0x39, 0xe1,
	0xc6, 0x2c, 0xb7, 0x2d, 0x17, 0x40, 0xb5, 0x35, 0x63, 0xb4, 0x98, 0x21, 0xa3, 0x55, 0x74, 0xb3,
	0x66, 0x34, 0xf3, 0x6e, 0xfe, 0x96, 0x7f, 0x3a, 0xb4, 0x54, 0x5a, 0xf5, 0xb3, 0xd1, 0x66, 0xec,
	0xb8, 0x20, 0xb4, 0x95, 0xfa, 0xe6, 0x6f, 0x05, 0xac, 0x78, 0x24, 0x9e, 0x5a, 0xc3, 0xa6, 0x3e,
	0x46, 0xc9, 0xfd, 0xfc, 0xed, 0xee, 0x2b, 0x6f, 0xb6, 0xcb, 0xe8, 0x80, 0x44, 0x28, 0x09, 0x4c,
	0x42, 0x03, 0x2b, 0xc0, 0x49, 0xf1, 0xb2, 0xe5, 0x86, 0x4f, 0x43, 0x36, 0xf9, 0x87, 0xb2, 0x2e,
	0x8d, 0xcf, 0x6a, 0x65, 0xab, 0xdd, 0xfe, 0xa2, 0xd6, 0xb6, 0x04, 0x61, 0xdb, 0x67, 0xa6, 0x30,
	0x73, 0xeb, 0xb0, 0x61, 0x96, 0x89, 0xd9, 0x77, 0x09, 0xe9, 0xb6, 0x7d, 0xd6, 0x1d, 0x40, 0xba,
	0x87, 0x8d, 0xae, 0x84, 0xfc, 0x52, 0x57, 0x84, 0xdf, 0xb6, 0xdb, 0x3e, 0xb3, 0xed, 0x01, 0xc8,
	0xb6, 0x0f, 0x1b, 0xb6, 0x2d, 0x61, 0x47, 0x73, 0x45, 0x9d, 0x8f, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xcb, 0x9f, 0x04, 0xf7, 0x06, 0x00, 0x00,
}
