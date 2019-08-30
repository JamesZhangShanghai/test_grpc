// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wwwin-github.cisco.com/cloud-cmts/cmts-hm-common/telemetry/src/pb/cmmgmt/cmmgmt.proto

/*
Package cmmgmt is a generated protocol buffer package.
It is generated from these files:
	wwwin-github.cisco.com/cloud-cmts/cmts-hm-common/telemetry/src/pb/cmmgmt/cmmgmt.proto
It has these top-level messages:
	GetcmRequest
	SetcmRequest
	CmStateTmMsg
	AllCmStatesTmMsg
*/
package cmmgmt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import comm_util_pb "../comm_util_pb"

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

type CmMgmtOper int32

const (
	CmMgmtOper_GET_BY_MAC         CmMgmtOper = 0
	CmMgmtOper_GET_BY_SID         CmMgmtOper = 1
	CmMgmtOper_GET_BY_MAC_FROM_DB CmMgmtOper = 2
	CmMgmtOper_GET_BY_SID_FROM_DB CmMgmtOper = 3
	CmMgmtOper_CREATE             CmMgmtOper = 100
	CmMgmtOper_UPDATE             CmMgmtOper = 101
	CmMgmtOper_DELETE             CmMgmtOper = 102
	CmMgmtOper_UPDATE_STATE_ONLY  CmMgmtOper = 103
	CmMgmtOper_UPDATE_RNG_STATE   CmMgmtOper = 104
	CmMgmtOper_UPDATE_NET_STATE   CmMgmtOper = 105
)

var CmMgmtOper_name = map[int32]string{
	0:   "GET_BY_MAC",
	1:   "GET_BY_SID",
	2:   "GET_BY_MAC_FROM_DB",
	3:   "GET_BY_SID_FROM_DB",
	100: "CREATE",
	101: "UPDATE",
	102: "DELETE",
	103: "UPDATE_STATE_ONLY",
	104: "UPDATE_RNG_STATE",
	105: "UPDATE_NET_STATE",
}
var CmMgmtOper_value = map[string]int32{
	"GET_BY_MAC":         0,
	"GET_BY_SID":         1,
	"GET_BY_MAC_FROM_DB": 2,
	"GET_BY_SID_FROM_DB": 3,
	"CREATE":             100,
	"UPDATE":             101,
	"DELETE":             102,
	"UPDATE_STATE_ONLY":  103,
	"UPDATE_RNG_STATE":   104,
	"UPDATE_NET_STATE":   105,
}

func (x CmMgmtOper) String() string {
	return proto.EnumName(CmMgmtOper_name, int32(x))
}
func (CmMgmtOper) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CMMgmt service to request Cable Modem infomation
type GetcmRequest struct {
	// Client who initiate this request
	Clientname string `protobuf:"bytes,1,opt,name=clientname" json:"clientname"`
	// Transaction ID
	Transaction uint32 `protobuf:"varint,2,opt,name=transaction" json:"transaction"`
	// CM Management Operation Enumeration
	Getoper CmMgmtOper `protobuf:"varint,3,opt,name=getoper,enum=cmmgmt.CmMgmtOper" json:"getoper"`
	// Parameters applicable to [GET_BY_MAC/GET_BY_MAC_FROM_DB] operation:
	//   Mandatory: MAC Address
	//
	// Parameters applicable to [GET_BY_SID/GET_BY_SID_FROM_DB] operation:
	//   Mandatory: Service Group, MAC Domain, Primary SID
	ReqParam *comm_util_pb.ModemStateMsg `protobuf:"bytes,4,opt,name=ReqParam" json:"ReqParam"`
}

func (m *GetcmRequest) Reset()                    { *m = GetcmRequest{} }
func (m *GetcmRequest) String() string            { return proto.CompactTextString(m) }
func (*GetcmRequest) ProtoMessage()               {}
func (*GetcmRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetcmRequest) GetClientname() string {
	if m != nil {
		return m.Clientname
	}
	return ""
}

func (m *GetcmRequest) GetTransaction() uint32 {
	if m != nil {
		return m.Transaction
	}
	return 0
}

func (m *GetcmRequest) GetGetoper() CmMgmtOper {
	if m != nil {
		return m.Getoper
	}
	return CmMgmtOper_GET_BY_MAC
}

func (m *GetcmRequest) GetReqParam() *comm_util_pb.ModemStateMsg {
	if m != nil {
		return m.ReqParam
	}
	return nil
}

// CMMgmt service to update Cable Modem infomation
type SetcmRequest struct {
	// Client who initiate this request
	Clientname string `protobuf:"bytes,1,opt,name=clientname" json:"clientname"`
	// Transaction ID
	Transaction uint32 `protobuf:"varint,2,opt,name=transaction" json:"transaction"`
	// CM Management Operation Enumeration
	Setoper CmMgmtOper `protobuf:"varint,3,opt,name=setoper,enum=cmmgmt.CmMgmtOper" json:"setoper"`
	// Parameters applicable to [CREATE] operation:
	//   Mandatory: MAC Address, Service Group, MAC Domain, DS Channel, US Channel, Primary SID
	//
	// Parameters applicable to [DELETE] operation:
	//   Mandatory: MAC Address
	//
	// Parameters applicable to [UPDATE_RNG_STATE] operation:
	//   Mandatory: MAC Address, Service Group, MAC Domain, US Channel, MAC State
	//
	// Parameters applicable to [UPDATE_NET_STATE] operation:
	//   Mandatory: MAC Address, MAC State
	//   Optional:  IP Address
	//
	ReqParam *comm_util_pb.ModemStateMsg `protobuf:"bytes,4,opt,name=ReqParam" json:"ReqParam"`
}

func (m *SetcmRequest) Reset()                    { *m = SetcmRequest{} }
func (m *SetcmRequest) String() string            { return proto.CompactTextString(m) }
func (*SetcmRequest) ProtoMessage()               {}
func (*SetcmRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SetcmRequest) GetClientname() string {
	if m != nil {
		return m.Clientname
	}
	return ""
}

func (m *SetcmRequest) GetTransaction() uint32 {
	if m != nil {
		return m.Transaction
	}
	return 0
}

func (m *SetcmRequest) GetSetoper() CmMgmtOper {
	if m != nil {
		return m.Setoper
	}
	return CmMgmtOper_GET_BY_MAC
}

func (m *SetcmRequest) GetReqParam() *comm_util_pb.ModemStateMsg {
	if m != nil {
		return m.ReqParam
	}
	return nil
}

type CmStateTmMsg struct {
	CmState  *comm_util_pb.ModemStateMsg `protobuf:"bytes,1,opt,name=CmState" json:"CmState"`
	IsDelete bool                        `protobuf:"varint,2,opt,name=is_delete,json=isDelete" json:"isDelete"`
}

func (m *CmStateTmMsg) Reset()                    { *m = CmStateTmMsg{} }
func (m *CmStateTmMsg) String() string            { return proto.CompactTextString(m) }
func (*CmStateTmMsg) ProtoMessage()               {}
func (*CmStateTmMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CmStateTmMsg) GetCmState() *comm_util_pb.ModemStateMsg {
	if m != nil {
		return m.CmState
	}
	return nil
}

func (m *CmStateTmMsg) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

type AllCmStatesTmMsg struct {
	AllCmStatesTm []*CmStateTmMsg `protobuf:"bytes,1,rep,name=AllCmStatesTm" json:"AllCmStatesTm"`
}

func (m *AllCmStatesTmMsg) Reset()                    { *m = AllCmStatesTmMsg{} }
func (m *AllCmStatesTmMsg) String() string            { return proto.CompactTextString(m) }
func (*AllCmStatesTmMsg) ProtoMessage()               {}
func (*AllCmStatesTmMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AllCmStatesTmMsg) GetAllCmStatesTm() []*CmStateTmMsg {
	if m != nil {
		return m.AllCmStatesTm
	}
	return nil
}

func init() {
	proto.RegisterType((*GetcmRequest)(nil), "cmmgmt.GetcmRequest")
	proto.RegisterType((*SetcmRequest)(nil), "cmmgmt.SetcmRequest")
	proto.RegisterType((*CmStateTmMsg)(nil), "cmmgmt.CmStateTmMsg")
	proto.RegisterType((*AllCmStatesTmMsg)(nil), "cmmgmt.AllCmStatesTmMsg")
	proto.RegisterEnum("cmmgmt.CmMgmtOper", CmMgmtOper_name, CmMgmtOper_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CMMgmt service

type CMMgmtClient interface {
	// Get Cable Modem Information from memory or database
	GetCM(ctx context.Context, in *GetcmRequest, opts ...grpc.CallOption) (*comm_util_pb.ModemStateMsg, error)
	// CREATE, UPDATE, or DELETE Cable Modem Information
	SetCM(ctx context.Context, in *SetcmRequest, opts ...grpc.CallOption) (*comm_util_pb.ModemStateMsg, error)
}

type cMMgmtClient struct {
	cc *grpc.ClientConn
}

func NewCMMgmtClient(cc *grpc.ClientConn) CMMgmtClient {
	return &cMMgmtClient{cc}
}

func (c *cMMgmtClient) GetCM(ctx context.Context, in *GetcmRequest, opts ...grpc.CallOption) (*comm_util_pb.ModemStateMsg, error) {
	out := new(comm_util_pb.ModemStateMsg)
	err := grpc.Invoke(ctx, "/cmmgmt.CMMgmt/GetCM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cMMgmtClient) SetCM(ctx context.Context, in *SetcmRequest, opts ...grpc.CallOption) (*comm_util_pb.ModemStateMsg, error) {
	out := new(comm_util_pb.ModemStateMsg)
	err := grpc.Invoke(ctx, "/cmmgmt.CMMgmt/SetCM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CMMgmt service

type CMMgmtServer interface {
	// Get Cable Modem Information from memory or database
	GetCM(context.Context, *GetcmRequest) (*comm_util_pb.ModemStateMsg, error)
	// CREATE, UPDATE, or DELETE Cable Modem Information
	SetCM(context.Context, *SetcmRequest) (*comm_util_pb.ModemStateMsg, error)
}

func RegisterCMMgmtServer(s *grpc.Server, srv CMMgmtServer) {
	s.RegisterService(&_CMMgmt_serviceDesc, srv)
}

func _CMMgmt_GetCM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetcmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CMMgmtServer).GetCM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmmgmt.CMMgmt/GetCM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CMMgmtServer).GetCM(ctx, req.(*GetcmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CMMgmt_SetCM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetcmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CMMgmtServer).SetCM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmmgmt.CMMgmt/SetCM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CMMgmtServer).SetCM(ctx, req.(*SetcmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CMMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cmmgmt.CMMgmt",
	HandlerType: (*CMMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCM",
			Handler:    _CMMgmt_GetCM_Handler,
		},
		{
			MethodName: "SetCM",
			Handler:    _CMMgmt_SetCM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wwwin-github.cisco.com/cloud-cmts/cmts-hm-common/telemetry/src/pb/cmmgmt/cmmgmt.proto",
}

func init() {
	proto.RegisterFile("wwwin-github.cisco.com/cloud-cmts/cmts-hm-common/telemetry/src/pb/cmmgmt/cmmgmt.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbb, 0x0d, 0xa4, 0xe9, 0xa4, 0xad, 0xcc, 0xaa, 0xa0, 0xa8, 0x95, 0x90, 0x95, 0x53,
	0x84, 0x88, 0x2d, 0x82, 0x10, 0x12, 0xe2, 0x92, 0xda, 0x26, 0xaa, 0x54, 0x27, 0xd5, 0xda, 0x1c,
	0x7a, 0xb2, 0x9c, 0xcd, 0xe0, 0x1a, 0xbc, 0xde, 0xd4, 0xbb, 0x39, 0x70, 0xe7, 0x9d, 0x78, 0x00,
	0x5e, 0x0c, 0x39, 0x76, 0x13, 0xe7, 0x42, 0x91, 0xda, 0xcb, 0x7a, 0xe6, 0x9b, 0x19, 0xcf, 0xff,
	0xdb, 0x5a, 0x58, 0xcc, 0x53, 0x3d, 0x5f, 0xf1, 0x1f, 0xa8, 0x87, 0x98, 0x27, 0x43, 0xf5, 0x9d,
	0xbf, 0xb3, 0x78, 0xaa, 0xb8, 0xb4, 0xb8, 0x14, 0xf6, 0xa6, 0x68, 0xf3, 0x4c, 0xae, 0x16, 0x5c,
	0x68, 0x65, 0x97, 0xc7, 0xf0, 0x56, 0x0c, 0xb9, 0x14, 0x42, 0xe6, 0xb6, 0xc6, 0x0c, 0x05, 0xea,
	0xe2, 0xa7, 0xad, 0x0a, 0x6e, 0x2f, 0xe7, 0x36, 0x17, 0x22, 0x11, 0xba, 0x7e, 0x58, 0xcb, 0x42,
	0x6a, 0x49, 0xdb, 0x55, 0x76, 0x26, 0x9e, 0x7e, 0x9b, 0x14, 0x22, 0x5a, 0xe9, 0x34, 0x8b, 0x9a,
	0x49, 0xb5, 0xb6, 0xff, 0x9b, 0xc0, 0xd1, 0x04, 0x35, 0x17, 0x0c, 0xef, 0x56, 0xa8, 0x34, 0x7d,
	0x0d, 0xc0, 0xb3, 0x14, 0x73, 0x9d, 0xc7, 0x02, 0x7b, 0xc4, 0x24, 0x83, 0x43, 0xd6, 0x20, 0xd4,
	0x84, 0xae, 0x2e, 0xe2, 0x5c, 0xc5, 0x5c, 0xa7, 0x32, 0xef, 0xed, 0x9b, 0x64, 0x70, 0xcc, 0x9a,
	0x88, 0xbe, 0x85, 0x83, 0x04, 0xb5, 0x5c, 0x62, 0xd1, 0x6b, 0x99, 0x64, 0x70, 0x32, 0xa2, 0x56,
	0xed, 0xd4, 0x11, 0x7e, 0x22, 0xf4, 0x6c, 0x89, 0x05, 0xbb, 0x6f, 0xa1, 0x1f, 0xa1, 0xc3, 0xf0,
	0xee, 0x3a, 0x2e, 0x62, 0xd1, 0x7b, 0x66, 0x92, 0x41, 0x77, 0x74, 0x6e, 0x35, 0x15, 0x5b, 0xbe,
	0x5c, 0xa0, 0x08, 0x74, 0xac, 0xd1, 0x57, 0x09, 0xdb, 0x34, 0xaf, 0x95, 0x07, 0x4f, 0xae, 0x5c,
	0x3d, 0xac, 0x5c, 0x3d, 0x56, 0xf9, 0x1c, 0x8e, 0x9c, 0x8a, 0x87, 0xc2, 0x57, 0x09, 0xfd, 0x00,
	0x07, 0x75, 0xbe, 0x56, 0xfd, 0xc0, 0x7b, 0xee, 0x7b, 0xe9, 0x39, 0x1c, 0xa6, 0x2a, 0x5a, 0x60,
	0x86, 0x1a, 0xd7, 0x6e, 0x3a, 0xac, 0x93, 0x2a, 0x77, 0x9d, 0xf7, 0xa7, 0x60, 0x8c, 0xb3, 0xac,
	0x6e, 0x55, 0xd5, 0x9e, 0x4f, 0x70, 0xbc, 0xc3, 0x7a, 0xc4, 0x6c, 0x0d, 0xba, 0xa3, 0xd3, 0xad,
	0xc9, 0xad, 0x28, 0xb6, 0xdb, 0xfa, 0xe6, 0x0f, 0x01, 0xd8, 0x7e, 0x04, 0x7a, 0x02, 0x30, 0xf1,
	0xc2, 0xe8, 0xe2, 0x26, 0xf2, 0xc7, 0x8e, 0xb1, 0xd7, 0xc8, 0x83, 0x4b, 0xd7, 0x20, 0xf4, 0x15,
	0xd0, 0x6d, 0x3d, 0xfa, 0xc2, 0x66, 0x7e, 0xe4, 0x5e, 0x18, 0xfb, 0x0d, 0x1e, 0x5c, 0xba, 0x1b,
	0xde, 0xa2, 0x00, 0x6d, 0x87, 0x79, 0xe3, 0xd0, 0x33, 0x16, 0x65, 0xfc, 0xf5, 0xda, 0x2d, 0x63,
	0x2c, 0x63, 0xd7, 0xbb, 0xf2, 0x42, 0xcf, 0xf8, 0x46, 0x5f, 0xc2, 0x8b, 0x8a, 0x47, 0x41, 0x58,
	0x9e, 0xb3, 0xe9, 0xd5, 0x8d, 0x91, 0xd0, 0x53, 0x30, 0x6a, 0xcc, 0xa6, 0x93, 0xaa, 0x64, 0xdc,
	0x36, 0xe8, 0xd4, 0x0b, 0x6b, 0x9a, 0x8e, 0x7e, 0x11, 0x68, 0x3b, 0x7e, 0xe9, 0x82, 0x7e, 0x86,
	0xe7, 0x13, 0xd4, 0x8e, 0x4f, 0x37, 0xf6, 0x9b, 0xd7, 0xe0, 0xec, 0x5f, 0xbf, 0xa0, 0xbf, 0x57,
	0x4e, 0x07, 0xbb, 0xd3, 0xc1, 0xff, 0x4f, 0xcf, 0xdb, 0xeb, 0xbb, 0xf7, 0xfe, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x41, 0xd6, 0xbb, 0x0f, 0x5a, 0x04, 0x00, 0x00,
}