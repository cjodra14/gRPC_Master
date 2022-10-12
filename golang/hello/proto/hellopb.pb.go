// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hellopb.proto

package proto

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

type Hello struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c6d3de28842fa, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Hello) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type HelloRequest struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c6d3de28842fa, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type HelloResponse struct {
	CustomHello          string   `protobuf:"bytes,1,opt,name=custom_hello,json=customHello,proto3" json:"custom_hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c6d3de28842fa, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetCustomHello() string {
	if m != nil {
		return m.CustomHello
	}
	return ""
}

type HelloManyLanguagesRequest struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloManyLanguagesRequest) Reset()         { *m = HelloManyLanguagesRequest{} }
func (m *HelloManyLanguagesRequest) String() string { return proto.CompactTextString(m) }
func (*HelloManyLanguagesRequest) ProtoMessage()    {}
func (*HelloManyLanguagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c6d3de28842fa, []int{3}
}

func (m *HelloManyLanguagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloManyLanguagesRequest.Unmarshal(m, b)
}
func (m *HelloManyLanguagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloManyLanguagesRequest.Marshal(b, m, deterministic)
}
func (m *HelloManyLanguagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloManyLanguagesRequest.Merge(m, src)
}
func (m *HelloManyLanguagesRequest) XXX_Size() int {
	return xxx_messageInfo_HelloManyLanguagesRequest.Size(m)
}
func (m *HelloManyLanguagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloManyLanguagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloManyLanguagesRequest proto.InternalMessageInfo

func (m *HelloManyLanguagesRequest) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type HelloManyLanguagesResponse struct {
	HelloLanguage        string   `protobuf:"bytes,1,opt,name=hello_language,json=helloLanguage,proto3" json:"hello_language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloManyLanguagesResponse) Reset()         { *m = HelloManyLanguagesResponse{} }
func (m *HelloManyLanguagesResponse) String() string { return proto.CompactTextString(m) }
func (*HelloManyLanguagesResponse) ProtoMessage()    {}
func (*HelloManyLanguagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c6d3de28842fa, []int{4}
}

func (m *HelloManyLanguagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloManyLanguagesResponse.Unmarshal(m, b)
}
func (m *HelloManyLanguagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloManyLanguagesResponse.Marshal(b, m, deterministic)
}
func (m *HelloManyLanguagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloManyLanguagesResponse.Merge(m, src)
}
func (m *HelloManyLanguagesResponse) XXX_Size() int {
	return xxx_messageInfo_HelloManyLanguagesResponse.Size(m)
}
func (m *HelloManyLanguagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloManyLanguagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloManyLanguagesResponse proto.InternalMessageInfo

func (m *HelloManyLanguagesResponse) GetHelloLanguage() string {
	if m != nil {
		return m.HelloLanguage
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "hello.Hello")
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.HelloResponse")
	proto.RegisterType((*HelloManyLanguagesRequest)(nil), "hello.HelloManyLanguagesRequest")
	proto.RegisterType((*HelloManyLanguagesResponse)(nil), "hello.HelloManyLanguagesResponse")
}

func init() { proto.RegisterFile("proto/hellopb.proto", fileDescriptor_457c6d3de28842fa) }

var fileDescriptor_457c6d3de28842fa = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0x48, 0xd2, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x5c,
	0x25, 0x3b, 0x2e, 0x56, 0x0f, 0x10, 0x43, 0x48, 0x96, 0x8b, 0x2b, 0x2d, 0xb3, 0xa8, 0xb8, 0x24,
	0x3e, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x13, 0x2c, 0xe2, 0x97,
	0x98, 0x9b, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x50, 0x94, 0x9a, 0x96, 0x59, 0x21, 0xc1, 0x04, 0x96,
	0x82, 0xf2, 0x94, 0x8c, 0xb8, 0x78, 0xc0, 0xfa, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x94, 0xb8, 0x20, 0x06, 0x83, 0x4d, 0xe0, 0x36, 0xe2, 0xd1, 0x03, 0xf3, 0xf4, 0x20, 0x6a, 0xa0,
	0x76, 0x1a, 0x71, 0xf1, 0x42, 0xf5, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29, 0x72, 0xf1,
	0x24, 0x97, 0x16, 0x97, 0xe4, 0xe7, 0xc6, 0x23, 0xf4, 0x72, 0x06, 0x71, 0x43, 0xc4, 0xc0, 0x4a,
	0x95, 0xec, 0xb9, 0x24, 0xc1, 0x0c, 0xdf, 0xc4, 0xbc, 0x4a, 0x9f, 0xc4, 0xbc, 0xf4, 0xd2, 0xc4,
	0xf4, 0xd4, 0x62, 0x52, 0x2c, 0x75, 0xe6, 0x92, 0xc2, 0x66, 0x00, 0xd4, 0x05, 0xaa, 0x5c, 0x7c,
	0x60, 0x65, 0xf1, 0x39, 0x50, 0x29, 0xa8, 0x1b, 0x78, 0xc1, 0xa2, 0x30, 0xf5, 0x46, 0x0b, 0x19,
	0xa1, 0xde, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0x82, 0x05, 0x9f, 0x30, 0x8a,
	0x9d, 0x10, 0x77, 0x49, 0x89, 0xa0, 0x0a, 0x42, 0xed, 0x8a, 0xe5, 0x12, 0xc2, 0x74, 0x89, 0x90,
	0x02, 0xb2, 0x5a, 0x6c, 0xbe, 0x94, 0x52, 0xc4, 0xa3, 0x02, 0x62, 0xb4, 0x12, 0x83, 0x01, 0xa3,
	0x13, 0x7b, 0x14, 0x2b, 0x38, 0x86, 0x93, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0xb4, 0x9b, 0xc2, 0xff, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	//Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	//Server streaming
	//This service will return  hello/greeting in different languages
	HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/hello.HelloService/HelloManyLanguages", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloManyLanguagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloManyLanguagesClient interface {
	Recv() (*HelloManyLanguagesResponse, error)
	grpc.ClientStream
}

type helloServiceHelloManyLanguagesClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloManyLanguagesClient) Recv() (*HelloManyLanguagesResponse, error) {
	m := new(HelloManyLanguagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	//Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	//Server streaming
	//This service will return  hello/greeting in different languages
	HelloManyLanguages(*HelloManyLanguagesRequest, HelloService_HelloManyLanguagesServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) HelloManyLanguages(req *HelloManyLanguagesRequest, srv HelloService_HelloManyLanguagesServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloManyLanguages not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloManyLanguages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloManyLanguagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloManyLanguages(m, &helloServiceHelloManyLanguagesServer{stream})
}

type HelloService_HelloManyLanguagesServer interface {
	Send(*HelloManyLanguagesResponse) error
	grpc.ServerStream
}

type helloServiceHelloManyLanguagesServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloManyLanguagesServer) Send(m *HelloManyLanguagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloManyLanguages",
			Handler:       _HelloService_HelloManyLanguages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/hellopb.proto",
}
