// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/ante/v2/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_859fdaca3f2040a6, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_859fdaca3f2040a6, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryMinimumCommissionRequest is the request type for the Query/MinimumCommission RPC method.
type QueryMinimumCommissionRequest struct {
}

func (m *QueryMinimumCommissionRequest) Reset()         { *m = QueryMinimumCommissionRequest{} }
func (m *QueryMinimumCommissionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMinimumCommissionRequest) ProtoMessage()    {}
func (*QueryMinimumCommissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_859fdaca3f2040a6, []int{2}
}
func (m *QueryMinimumCommissionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinimumCommissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinimumCommissionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinimumCommissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinimumCommissionRequest.Merge(m, src)
}
func (m *QueryMinimumCommissionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinimumCommissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinimumCommissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinimumCommissionRequest proto.InternalMessageInfo

// QueryMinimumCommissionResponse is the response type for the Query/MinimumCommission RPC method.
type QueryMinimumCommissionResponse struct {
	// minimum commission enforced to all validators
	MinimumCommission github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=minimum_commission,json=minimumCommission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minimum_commission" yaml:"minimum_commission"`
}

func (m *QueryMinimumCommissionResponse) Reset()         { *m = QueryMinimumCommissionResponse{} }
func (m *QueryMinimumCommissionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMinimumCommissionResponse) ProtoMessage()    {}
func (*QueryMinimumCommissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_859fdaca3f2040a6, []int{3}
}
func (m *QueryMinimumCommissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinimumCommissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinimumCommissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinimumCommissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinimumCommissionResponse.Merge(m, src)
}
func (m *QueryMinimumCommissionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinimumCommissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinimumCommissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinimumCommissionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "terra.ante.v2.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "terra.ante.v2.QueryParamsResponse")
	proto.RegisterType((*QueryMinimumCommissionRequest)(nil), "terra.ante.v2.QueryMinimumCommissionRequest")
	proto.RegisterType((*QueryMinimumCommissionResponse)(nil), "terra.ante.v2.QueryMinimumCommissionResponse")
}

func init() { proto.RegisterFile("terra/ante/v2/query.proto", fileDescriptor_859fdaca3f2040a6) }

var fileDescriptor_859fdaca3f2040a6 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x4d, 0x8a, 0x16, 0x1c, 0xf1, 0xd0, 0xb1, 0x45, 0x1b, 0x6c, 0x62, 0x23, 0x88, 0x8a, 0xcd,
	0x40, 0x7a, 0xf3, 0x58, 0x7b, 0x52, 0x04, 0xed, 0xd1, 0x8b, 0x4c, 0xe3, 0x10, 0x83, 0x9d, 0xf9,
	0xd2, 0xcc, 0xa4, 0x10, 0x8f, 0xfe, 0x82, 0x85, 0x65, 0x4f, 0xbb, 0x3f, 0xa8, 0xc7, 0xc2, 0x5e,
	0x96, 0x3d, 0x94, 0xa5, 0xdd, 0x5f, 0xb0, 0xbf, 0x60, 0xc9, 0x24, 0xbb, 0x6c, 0x9b, 0x76, 0xd9,
	0x53, 0xc2, 0xbc, 0x37, 0xef, 0xbd, 0xef, 0x7d, 0x83, 0xda, 0x8a, 0x25, 0x09, 0x25, 0x54, 0x28,
	0x46, 0x66, 0x3e, 0x99, 0xa6, 0x2c, 0xc9, 0xbc, 0x38, 0x01, 0x05, 0xf8, 0x99, 0x86, 0xbc, 0x1c,
	0xf2, 0x66, 0xbe, 0xd5, 0x0c, 0x21, 0x04, 0x8d, 0x90, 0xfc, 0xaf, 0x20, 0x59, 0xaf, 0x42, 0x80,
	0x70, 0xc2, 0x08, 0x8d, 0x23, 0x42, 0x85, 0x00, 0x45, 0x55, 0x04, 0x42, 0x96, 0xa8, 0xb5, 0xa9,
	0x1e, 0xd3, 0x84, 0xf2, 0x12, 0x73, 0x9b, 0x08, 0xff, 0xc8, 0xdd, 0xbe, 0xeb, 0xc3, 0x11, 0x9b,
	0xa6, 0x4c, 0x2a, 0xf7, 0x0b, 0x7a, 0xbe, 0x71, 0x2a, 0x63, 0x10, 0x92, 0xe1, 0x3e, 0xaa, 0x17,
	0x97, 0x5f, 0x9a, 0xaf, 0xcd, 0x77, 0x4f, 0xfd, 0x96, 0xb7, 0x11, 0xce, 0x2b, 0xe8, 0x83, 0x47,
	0xf3, 0xa5, 0x63, 0x8c, 0x4a, 0xaa, 0xeb, 0xa0, 0x8e, 0xd6, 0xfa, 0x16, 0x89, 0x88, 0xa7, 0xfc,
	0x33, 0x70, 0x1e, 0x49, 0x19, 0x81, 0xb8, 0x31, 0x3b, 0x31, 0x91, 0xbd, 0x8f, 0x51, 0x1a, 0xff,
	0x43, 0x98, 0x17, 0xe0, 0xaf, 0xe0, 0x16, 0xd5, 0x21, 0x9e, 0x0c, 0xbe, 0xe6, 0x6e, 0xe7, 0x4b,
	0xe7, 0x6d, 0x18, 0xa9, 0x3f, 0xe9, 0xd8, 0x0b, 0x80, 0x93, 0x00, 0x24, 0x07, 0x59, 0x7e, 0x7a,
	0xf2, 0xf7, 0x5f, 0xa2, 0xb2, 0x98, 0x49, 0x6f, 0xc8, 0x82, 0xab, 0xa5, 0xd3, 0xce, 0x28, 0x9f,
	0x7c, 0x72, 0xab, 0x8a, 0xee, 0xa8, 0xc1, 0xb7, 0x33, 0xf8, 0x47, 0x35, 0xf4, 0x58, 0xc7, 0xc3,
	0x02, 0xd5, 0x8b, 0x09, 0x71, 0x77, 0x6b, 0xf0, 0x6a, 0x85, 0x96, 0x7b, 0x1f, 0xa5, 0x18, 0xcb,
	0xed, 0xfc, 0x3f, 0xbd, 0x3c, 0xac, 0xbd, 0xc0, 0x2d, 0xb2, 0x6b, 0x43, 0xf8, 0xd8, 0x44, 0x8d,
	0x4a, 0x27, 0xf8, 0xe3, 0x2e, 0xe1, 0x7d, 0xe5, 0x5a, 0xbd, 0x07, 0xb2, 0xcb, 0x44, 0xef, 0x75,
	0xa2, 0x37, 0xb8, 0xbb, 0x95, 0xa8, 0xda, 0xd5, 0x60, 0x38, 0x5f, 0xd9, 0xe6, 0x62, 0x65, 0x9b,
	0x17, 0x2b, 0xdb, 0x3c, 0x58, 0xdb, 0xc6, 0x62, 0x6d, 0x1b, 0x67, 0x6b, 0xdb, 0xf8, 0xf9, 0xe1,
	0xce, 0x26, 0xb4, 0x4c, 0x8f, 0x83, 0x60, 0x19, 0x09, 0x20, 0xc9, 0x9f, 0x68, 0x5c, 0xa8, 0xea,
	0x8d, 0x8c, 0xeb, 0xfa, 0x19, 0xf6, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x5c, 0x60, 0xc0,
	0x02, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries params of the ante.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// MinimumCommission queries minimum commission rate for all validators.
	MinimumCommission(ctx context.Context, in *QueryMinimumCommissionRequest, opts ...grpc.CallOption) (*QueryMinimumCommissionResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/terra.ante.v2.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MinimumCommission(ctx context.Context, in *QueryMinimumCommissionRequest, opts ...grpc.CallOption) (*QueryMinimumCommissionResponse, error) {
	out := new(QueryMinimumCommissionResponse)
	err := c.cc.Invoke(ctx, "/terra.ante.v2.Query/MinimumCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries params of the ante.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// MinimumCommission queries minimum commission rate for all validators.
	MinimumCommission(context.Context, *QueryMinimumCommissionRequest) (*QueryMinimumCommissionResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) MinimumCommission(ctx context.Context, req *QueryMinimumCommissionRequest) (*QueryMinimumCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinimumCommission not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.ante.v2.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MinimumCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMinimumCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinimumCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.ante.v2.Query/MinimumCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinimumCommission(ctx, req.(*QueryMinimumCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.ante.v2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "MinimumCommission",
			Handler:    _Query_MinimumCommission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/ante/v2/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryMinimumCommissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinimumCommissionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinimumCommissionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMinimumCommissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinimumCommissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinimumCommissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinimumCommission.Size()
		i -= size
		if _, err := m.MinimumCommission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryMinimumCommissionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMinimumCommissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinimumCommission.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryMinimumCommissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryMinimumCommissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinimumCommissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryMinimumCommissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryMinimumCommissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinimumCommissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumCommission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumCommission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
