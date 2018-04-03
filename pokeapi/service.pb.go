// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokeapi/service.proto

/*
Package pokeapi is a generated protocol buffer package.

It is generated from these files:
	pokeapi/service.proto

It has these top-level messages:
	GetRequest
	GetReply
*/
package pokeapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the Pokemon unique identifier.
type GetRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the Pokemon information.
type GetReply struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Weight string `protobuf:"bytes,3,opt,name=weight" json:"weight,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetReply) GetWeight() string {
	if m != nil {
		return m.Weight
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "pokeapi.GetRequest")
	proto.RegisterType((*GetReply)(nil), "pokeapi.GetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pokeapi service

type PokeapiClient interface {
	// GetPokemon get information about a pokemon.
	GetPokemon(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type pokeapiClient struct {
	cc *grpc.ClientConn
}

func NewPokeapiClient(cc *grpc.ClientConn) PokeapiClient {
	return &pokeapiClient{cc}
}

func (c *pokeapiClient) GetPokemon(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/pokeapi.Pokeapi/GetPokemon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pokeapi service

type PokeapiServer interface {
	// GetPokemon get information about a pokemon.
	GetPokemon(context.Context, *GetRequest) (*GetReply, error)
}

func RegisterPokeapiServer(s *grpc.Server, srv PokeapiServer) {
	s.RegisterService(&_Pokeapi_serviceDesc, srv)
}

func _Pokeapi_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokeapiServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokeapi.Pokeapi/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokeapiServer).GetPokemon(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pokeapi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokeapi.Pokeapi",
	HandlerType: (*PokeapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPokemon",
			Handler:    _Pokeapi_GetPokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokeapi/service.proto",
}

func init() { proto.RegisterFile("pokeapi/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3d, 0x4b, 0xc6, 0x30,
	0x14, 0x85, 0x7d, 0xeb, 0x4b, 0xab, 0x77, 0x10, 0x8c, 0x28, 0x45, 0x1c, 0xa4, 0x93, 0x4b, 0x93,
	0x7e, 0x88, 0xb8, 0xea, 0xa0, 0xab, 0x74, 0x74, 0xeb, 0xc7, 0xa5, 0x0d, 0x36, 0x4d, 0x6c, 0x6f,
	0x94, 0xfa, 0xeb, 0xa5, 0x31, 0xa0, 0x6e, 0x39, 0xcf, 0x43, 0x38, 0xe7, 0xc2, 0xb9, 0xd1, 0x6f,
	0x58, 0x1b, 0x29, 0x16, 0x9c, 0x3f, 0x64, 0x8b, 0xdc, 0xcc, 0x9a, 0x34, 0x8b, 0x3c, 0x4e, 0xae,
	0x00, 0x9e, 0x91, 0x2a, 0x7c, 0xb7, 0xb8, 0x10, 0x3b, 0x81, 0x40, 0x76, 0xf1, 0xee, 0x7a, 0x77,
	0xb3, 0xaf, 0x02, 0xd9, 0x25, 0x4f, 0x70, 0xe4, 0xac, 0x19, 0xd7, 0x3f, 0xee, 0x78, 0x73, 0x8c,
	0xc1, 0x7e, 0xaa, 0x15, 0xc6, 0x81, 0x23, 0xee, 0xcd, 0x2e, 0x20, 0xfc, 0x44, 0xd9, 0x0f, 0x14,
	0x1f, 0x3a, 0xea, 0x53, 0xf1, 0x00, 0xd1, 0xcb, 0x4f, 0x21, 0xbb, 0x73, 0x85, 0x5b, 0x52, 0x7a,
	0x62, 0x67, 0xdc, 0x0f, 0xe1, 0xbf, 0x2b, 0x2e, 0x4f, 0xff, 0x43, 0x33, 0xae, 0xc9, 0xc1, 0x63,
	0xf9, 0x9a, 0xf7, 0x92, 0x06, 0xdb, 0xf0, 0x56, 0x2b, 0xb1, 0x10, 0xe2, 0xf4, 0xa5, 0x2d, 0x09,
	0x85, 0x48, 0xd6, 0xa4, 0x45, 0x96, 0xdf, 0x67, 0xb7, 0x59, 0x99, 0xf6, 0xb3, 0x69, 0x85, 0xff,
	0xdf, 0x84, 0xee, 0xda, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x20, 0xa6, 0xb6, 0xc5, 0x06, 0x01,
	0x00, 0x00,
}
