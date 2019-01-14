// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fridge.proto

package fridge

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetItemsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsRequest) Reset()         { *m = GetItemsRequest{} }
func (m *GetItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemsRequest) ProtoMessage()    {}
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_188e2283b7eee495, []int{0}
}

func (m *GetItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsRequest.Unmarshal(m, b)
}
func (m *GetItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsRequest.Merge(m, src)
}
func (m *GetItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemsRequest.Size(m)
}
func (m *GetItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsRequest proto.InternalMessageInfo

type ItemsResponse struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemsResponse) Reset()         { *m = ItemsResponse{} }
func (m *ItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ItemsResponse) ProtoMessage()    {}
func (*ItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_188e2283b7eee495, []int{1}
}

func (m *ItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemsResponse.Unmarshal(m, b)
}
func (m *ItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemsResponse.Marshal(b, m, deterministic)
}
func (m *ItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemsResponse.Merge(m, src)
}
func (m *ItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ItemsResponse.Size(m)
}
func (m *ItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemsResponse proto.InternalMessageInfo

func (m *ItemsResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_188e2283b7eee495, []int{2}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetItemsRequest)(nil), "GetItemsRequest")
	proto.RegisterType((*ItemsResponse)(nil), "ItemsResponse")
	proto.RegisterType((*Item)(nil), "Item")
}

func init() { proto.RegisterFile("fridge.proto", fileDescriptor_188e2283b7eee495) }

var fileDescriptor_188e2283b7eee495 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2b, 0xca, 0x4c,
	0x49, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe4, 0xe2, 0x77, 0x4f, 0x2d, 0xf1,
	0x2c, 0x49, 0xcd, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xd2, 0xe1, 0xe2, 0x85,
	0xf2, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb9, 0x58, 0x33, 0x41, 0x02, 0x12, 0x8c,
	0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xac, 0x7a, 0x20, 0xe9, 0x20, 0x88, 0x98, 0x92, 0x14, 0x17, 0x0b,
	0x88, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x1b, 0xc5, 0x73, 0xf1, 0xba, 0x81, 0x2d, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x15, 0xd2, 0xe1, 0xe2, 0x80, 0xd9, 0x26, 0x24, 0xa0, 0x87, 0x66, 0xb1, 0x14, 0x9f, 0x1e, 0xaa,
	0xbd, 0x0a, 0x5c, 0xec, 0x8e, 0x29, 0x29, 0x60, 0xd3, 0x21, 0x76, 0xa2, 0xab, 0x48, 0x62, 0x03,
	0x7b, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x45, 0x5a, 0xce, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FridgeServiceClient is the client API for FridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FridgeServiceClient interface {
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error)
	AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemsResponse, error)
}

type fridgeServiceClient struct {
	cc *grpc.ClientConn
}

func NewFridgeServiceClient(cc *grpc.ClientConn) FridgeServiceClient {
	return &fridgeServiceClient{cc}
}

func (c *fridgeServiceClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/FridgeService/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fridgeServiceClient) AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/FridgeService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FridgeServiceServer is the server API for FridgeService service.
type FridgeServiceServer interface {
	GetItems(context.Context, *GetItemsRequest) (*ItemsResponse, error)
	AddItem(context.Context, *Item) (*ItemsResponse, error)
}

func RegisterFridgeServiceServer(s *grpc.Server, srv FridgeServiceServer) {
	s.RegisterService(&_FridgeService_serviceDesc, srv)
}

func _FridgeService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FridgeServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FridgeService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FridgeServiceServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FridgeService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FridgeServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FridgeService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FridgeServiceServer).AddItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

var _FridgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FridgeService",
	HandlerType: (*FridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItems",
			Handler:    _FridgeService_GetItems_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _FridgeService_AddItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fridge.proto",
}
