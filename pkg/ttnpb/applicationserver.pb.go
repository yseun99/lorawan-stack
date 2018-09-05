// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import context "context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SetAsApplicationAPIKeyRequest struct {
	ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,embedded=application_ids" json:"application_ids"`
	APIKey                 string   `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SetAsApplicationAPIKeyRequest) Reset()      { *m = SetAsApplicationAPIKeyRequest{} }
func (*SetAsApplicationAPIKeyRequest) ProtoMessage() {}
func (*SetAsApplicationAPIKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationserver_c20314c5a369681c, []int{0}
}
func (m *SetAsApplicationAPIKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetAsApplicationAPIKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetAsApplicationAPIKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetAsApplicationAPIKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAsApplicationAPIKeyRequest.Merge(dst, src)
}
func (m *SetAsApplicationAPIKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetAsApplicationAPIKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAsApplicationAPIKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAsApplicationAPIKeyRequest proto.InternalMessageInfo

func (m *SetAsApplicationAPIKeyRequest) GetAPIKey() string {
	if m != nil {
		return m.APIKey
	}
	return ""
}

func init() {
	proto.RegisterType((*SetAsApplicationAPIKeyRequest)(nil), "ttn.lorawan.v3.SetAsApplicationAPIKeyRequest")
	golang_proto.RegisterType((*SetAsApplicationAPIKeyRequest)(nil), "ttn.lorawan.v3.SetAsApplicationAPIKeyRequest")
}
func (this *SetAsApplicationAPIKeyRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetAsApplicationAPIKeyRequest)
	if !ok {
		that2, ok := that.(SetAsApplicationAPIKeyRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIdentifiers.Equal(&that1.ApplicationIdentifiers) {
		return false
	}
	if this.APIKey != that1.APIKey {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for As service

type AsClient interface {
	// Set the API key that the Application Server should use to link to the Network Server.
	SetApplicationAPIKey(ctx context.Context, in *SetAsApplicationAPIKeyRequest, opts ...grpc.CallOption) (*types.Empty, error)
	DeleteApplicationAPIKey(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	Subscribe(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (As_SubscribeClient, error)
}

type asClient struct {
	cc *grpc.ClientConn
}

func NewAsClient(cc *grpc.ClientConn) AsClient {
	return &asClient{cc}
}

func (c *asClient) SetApplicationAPIKey(ctx context.Context, in *SetAsApplicationAPIKeyRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.As/SetApplicationAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asClient) DeleteApplicationAPIKey(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.As/DeleteApplicationAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asClient) Subscribe(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (As_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_As_serviceDesc.Streams[0], "/ttn.lorawan.v3.As/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &asSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type As_SubscribeClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type asSubscribeClient struct {
	grpc.ClientStream
}

func (x *asSubscribeClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for As service

type AsServer interface {
	// Set the API key that the Application Server should use to link to the Network Server.
	SetApplicationAPIKey(context.Context, *SetAsApplicationAPIKeyRequest) (*types.Empty, error)
	DeleteApplicationAPIKey(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
	Subscribe(*ApplicationIdentifiers, As_SubscribeServer) error
}

func RegisterAsServer(s *grpc.Server, srv AsServer) {
	s.RegisterService(&_As_serviceDesc, srv)
}

func _As_SetApplicationAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAsApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsServer).SetApplicationAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.As/SetApplicationAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsServer).SetApplicationAPIKey(ctx, req.(*SetAsApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _As_DeleteApplicationAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsServer).DeleteApplicationAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.As/DeleteApplicationAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsServer).DeleteApplicationAPIKey(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _As_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationIdentifiers)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AsServer).Subscribe(m, &asSubscribeServer{stream})
}

type As_SubscribeServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type asSubscribeServer struct {
	grpc.ServerStream
}

func (x *asSubscribeServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

var _As_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.As",
	HandlerType: (*AsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetApplicationAPIKey",
			Handler:    _As_SetApplicationAPIKey_Handler,
		},
		{
			MethodName: "DeleteApplicationAPIKey",
			Handler:    _As_DeleteApplicationAPIKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _As_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver.proto",
}

func (m *SetAsApplicationAPIKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetAsApplicationAPIKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintApplicationserver(dAtA, i, uint64(m.ApplicationIdentifiers.Size()))
	n1, err := m.ApplicationIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.APIKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApplicationserver(dAtA, i, uint64(len(m.APIKey)))
		i += copy(dAtA[i:], m.APIKey)
	}
	return i, nil
}

func encodeVarintApplicationserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSetAsApplicationAPIKeyRequest(r randyApplicationserver, easy bool) *SetAsApplicationAPIKeyRequest {
	this := &SetAsApplicationAPIKeyRequest{}
	v1 := NewPopulatedApplicationIdentifiers(r, easy)
	this.ApplicationIdentifiers = *v1
	this.APIKey = randStringApplicationserver(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApplicationserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApplicationserver(r randyApplicationserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApplicationserver(r randyApplicationserver) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneApplicationserver(r)
	}
	return string(tmps)
}
func randUnrecognizedApplicationserver(r randyApplicationserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApplicationserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApplicationserver(dAtA []byte, r randyApplicationserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApplicationserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SetAsApplicationAPIKeyRequest) Size() (n int) {
	var l int
	_ = l
	l = m.ApplicationIdentifiers.Size()
	n += 1 + l + sovApplicationserver(uint64(l))
	l = len(m.APIKey)
	if l > 0 {
		n += 1 + l + sovApplicationserver(uint64(l))
	}
	return n
}

func sovApplicationserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplicationserver(x uint64) (n int) {
	return sovApplicationserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *SetAsApplicationAPIKeyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetAsApplicationAPIKeyRequest{`,
		`ApplicationIdentifiers:` + strings.Replace(strings.Replace(this.ApplicationIdentifiers.String(), "ApplicationIdentifiers", "ApplicationIdentifiers", 1), `&`, ``, 1) + `,`,
		`APIKey:` + fmt.Sprintf("%v", this.APIKey) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SetAsApplicationAPIKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplicationserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetAsApplicationAPIKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetAsApplicationAPIKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ApplicationIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplicationserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplicationserver
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
func skipApplicationserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplicationserver
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
					return 0, ErrIntOverflowApplicationserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplicationserver
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApplicationserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplicationserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApplicationserver(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApplicationserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplicationserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver.proto", fileDescriptor_applicationserver_c20314c5a369681c)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver.proto", fileDescriptor_applicationserver_c20314c5a369681c)
}

var fileDescriptor_applicationserver_c20314c5a369681c = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4f, 0x1b, 0x4d,
	0x10, 0xc5, 0x77, 0x5c, 0xf8, 0xfb, 0xb8, 0x48, 0x44, 0x3a, 0x45, 0x09, 0x72, 0xc4, 0xd8, 0x0a,
	0x52, 0x44, 0xa4, 0x78, 0x2f, 0x82, 0xbf, 0xc0, 0x56, 0x52, 0x20, 0x9a, 0x08, 0x92, 0x82, 0x48,
	0x11, 0xda, 0x33, 0xc3, 0x79, 0x65, 0xfb, 0x76, 0x73, 0xbb, 0x06, 0xb9, 0xa3, 0xa4, 0x4c, 0x91,
	0x22, 0x5d, 0xa2, 0x54, 0x94, 0x94, 0x94, 0x94, 0x2e, 0x5d, 0x52, 0x59, 0xdc, 0x5e, 0x43, 0x49,
	0x49, 0x19, 0x71, 0x5c, 0x64, 0x03, 0xc2, 0xa2, 0xdb, 0xd1, 0xfe, 0xf6, 0xbd, 0x37, 0x4f, 0xeb,
	0xbd, 0xe9, 0xaa, 0x44, 0xec, 0x8b, 0xb8, 0x6e, 0xac, 0x68, 0x75, 0x02, 0xa1, 0x65, 0x20, 0xb4,
	0xee, 0xca, 0x96, 0xb0, 0x52, 0xc5, 0x86, 0x92, 0x3d, 0x4a, 0xb8, 0x4e, 0x94, 0x55, 0xfe, 0xbc,
	0xb5, 0x31, 0x2f, 0x70, 0xbe, 0xb7, 0x5a, 0xa9, 0x47, 0xd2, 0xb6, 0xfb, 0x21, 0x6f, 0xa9, 0x5e,
	0x10, 0xa9, 0x48, 0x05, 0x39, 0x16, 0xf6, 0x77, 0xf3, 0x29, 0x1f, 0xf2, 0xd3, 0xcd, 0xf3, 0xca,
	0xcb, 0x48, 0xa9, 0xa8, 0x4b, 0x13, 0x8a, 0x7a, 0xda, 0x0e, 0x8a, 0xcb, 0xa5, 0x99, 0x31, 0x1e,
	0x86, 0xe4, 0x0e, 0xc5, 0x56, 0xee, 0x4a, 0x4a, 0x4c, 0x01, 0xd5, 0xee, 0x43, 0x3d, 0x32, 0x46,
	0x44, 0x54, 0x10, 0xaf, 0x7e, 0x81, 0xb7, 0xb8, 0x49, 0xb6, 0x61, 0x1a, 0x13, 0x87, 0xc6, 0xc7,
	0xb5, 0x75, 0x1a, 0x6c, 0xd0, 0xb7, 0x3e, 0x19, 0xeb, 0x6f, 0x79, 0x4f, 0xa7, 0xdc, 0xb7, 0xe5,
	0x8e, 0x59, 0x80, 0x1a, 0x2c, 0x3f, 0x59, 0x79, 0xcd, 0x6f, 0x77, 0xc0, 0xa7, 0x24, 0xd6, 0x26,
	0x51, 0x9a, 0xff, 0x0f, 0xc7, 0x55, 0x36, 0x1a, 0x57, 0x61, 0x63, 0x5e, 0x4c, 0x13, 0xc6, 0x5f,
	0xf2, 0xfe, 0x13, 0x5a, 0x6e, 0x77, 0x68, 0xb0, 0x50, 0xaa, 0xc1, 0xf2, 0x5c, 0xd3, 0x73, 0xe3,
	0x6a, 0xb9, 0xb0, 0x2f, 0x0b, 0x2d, 0xd7, 0x69, 0xb0, 0xf2, 0xa3, 0xe4, 0x95, 0x1a, 0xc6, 0xff,
	0xea, 0x3d, 0xbb, 0xce, 0x79, 0x37, 0xa5, 0x5f, 0xbf, 0x9b, 0x62, 0xe6, 0x36, 0x95, 0xe7, 0xfc,
	0xa6, 0x79, 0xfe, 0xaf, 0x79, 0xfe, 0xe1, 0xba, 0x79, 0x7f, 0xcb, 0x7b, 0xf1, 0x9e, 0xba, 0x64,
	0xe9, 0xbe, 0xc3, 0x23, 0xf7, 0x7c, 0x50, 0xfa, 0x93, 0x37, 0xb7, 0xd9, 0x0f, 0x4d, 0x2b, 0x91,
	0x21, 0x3d, 0x5a, 0x6c, 0x71, 0x06, 0xf7, 0x59, 0xbf, 0x83, 0xe6, 0x1f, 0x18, 0xa6, 0x08, 0xa3,
	0x14, 0xe1, 0x2c, 0x45, 0x76, 0x9e, 0x22, 0xbb, 0x48, 0x91, 0x5d, 0xa6, 0xc8, 0xae, 0x52, 0x84,
	0x03, 0x87, 0x70, 0xe8, 0x90, 0x1d, 0x39, 0x84, 0x63, 0x87, 0xec, 0xc4, 0x21, 0x3b, 0x75, 0xc8,
	0x86, 0x0e, 0x61, 0xe4, 0x10, 0xce, 0x1c, 0xb2, 0x73, 0x87, 0x70, 0xe1, 0x90, 0x5d, 0x3a, 0x84,
	0x2b, 0x87, 0xec, 0x20, 0x43, 0x76, 0x98, 0x21, 0x7c, 0xcf, 0x90, 0xfd, 0xcc, 0x10, 0x7e, 0x67,
	0xc8, 0x8e, 0x32, 0x64, 0xc7, 0x19, 0xc2, 0x49, 0x86, 0x70, 0x9a, 0x21, 0x7c, 0x79, 0x1b, 0x29,
	0x6e, 0xdb, 0x64, 0xdb, 0x32, 0x8e, 0x0c, 0x8f, 0xc9, 0xee, 0xab, 0xa4, 0x13, 0xdc, 0xfe, 0x66,
	0xba, 0x13, 0x05, 0xd6, 0xc6, 0x3a, 0x0c, 0xcb, 0x79, 0x15, 0xab, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x50, 0x2b, 0x56, 0x23, 0x59, 0x03, 0x00, 0x00,
}
