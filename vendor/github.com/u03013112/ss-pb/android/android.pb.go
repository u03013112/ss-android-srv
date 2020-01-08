// Code generated by protoc-gen-go. DO NOT EDIT.
// source: android/android.proto

/*
Package android is a generated protocol buffer package.

It is generated from these files:
	android/android.proto

It has these top-level messages:
	LoginRequest
	LoginReply
	GetConfigRequest
	GetConfigReply
	KeepaliveRequest
	KeepaliveReply
	GetProdectionListRequest
	Prodection
	GetProdectionListReply
	BuyTestRequest
	BuyTestReply
	GetGoogleAdRequest
	GetGoogleAdReply
	GetUserInfoRequest
	GetUserInfoReply
	GetConfigNewReply
*/
package android

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type LoginRequest struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *LoginRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LoginReply struct {
	Token       string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ExpiresDate int64  `protobuf:"varint,2,opt,name=expiresDate" json:"expiresDate,omitempty"`
	Total       int64  `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Used        int64  `protobuf:"varint,4,opt,name=used" json:"used,omitempty"`
	Error       string `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

func (m *LoginReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *LoginReply) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *LoginReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetConfigRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetConfigRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetConfigReply struct {
	IP     string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Port   string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	Method string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Passwd string `protobuf:"bytes,4,opt,name=passwd" json:"passwd,omitempty"`
	Error  string `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *GetConfigReply) Reset()                    { *m = GetConfigReply{} }
func (m *GetConfigReply) String() string            { return proto.CompactTextString(m) }
func (*GetConfigReply) ProtoMessage()               {}
func (*GetConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetConfigReply) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *GetConfigReply) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *GetConfigReply) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *GetConfigReply) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *GetConfigReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type KeepaliveRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Rx    int64  `protobuf:"varint,2,opt,name=rx" json:"rx,omitempty"`
}

func (m *KeepaliveRequest) Reset()                    { *m = KeepaliveRequest{} }
func (m *KeepaliveRequest) String() string            { return proto.CompactTextString(m) }
func (*KeepaliveRequest) ProtoMessage()               {}
func (*KeepaliveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KeepaliveRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *KeepaliveRequest) GetRx() int64 {
	if m != nil {
		return m.Rx
	}
	return 0
}

type KeepaliveReply struct {
	NeedStop    bool   `protobuf:"varint,1,opt,name=needStop" json:"needStop,omitempty"`
	ExpiresDate int64  `protobuf:"varint,2,opt,name=expiresDate" json:"expiresDate,omitempty"`
	Total       int64  `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Used        int64  `protobuf:"varint,4,opt,name=used" json:"used,omitempty"`
	Error       string `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *KeepaliveReply) Reset()                    { *m = KeepaliveReply{} }
func (m *KeepaliveReply) String() string            { return proto.CompactTextString(m) }
func (*KeepaliveReply) ProtoMessage()               {}
func (*KeepaliveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *KeepaliveReply) GetNeedStop() bool {
	if m != nil {
		return m.NeedStop
	}
	return false
}

func (m *KeepaliveReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

func (m *KeepaliveReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *KeepaliveReply) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *KeepaliveReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetProdectionListRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetProdectionListRequest) Reset()                    { *m = GetProdectionListRequest{} }
func (m *GetProdectionListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProdectionListRequest) ProtoMessage()               {}
func (*GetProdectionListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetProdectionListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Prodection struct {
	ID          int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Time        int64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Total       int64  `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
}

func (m *Prodection) Reset()                    { *m = Prodection{} }
func (m *Prodection) String() string            { return proto.CompactTextString(m) }
func (*Prodection) ProtoMessage()               {}
func (*Prodection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Prodection) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Prodection) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Prodection) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Prodection) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Prodection) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetProdectionListReply struct {
	ProdectionList []*Prodection `protobuf:"bytes,1,rep,name=prodectionList" json:"prodectionList,omitempty"`
	Error          string        `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *GetProdectionListReply) Reset()                    { *m = GetProdectionListReply{} }
func (m *GetProdectionListReply) String() string            { return proto.CompactTextString(m) }
func (*GetProdectionListReply) ProtoMessage()               {}
func (*GetProdectionListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetProdectionListReply) GetProdectionList() []*Prodection {
	if m != nil {
		return m.ProdectionList
	}
	return nil
}

func (m *GetProdectionListReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type BuyTestRequest struct {
	Token        string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ProdectionID int64  `protobuf:"varint,2,opt,name=prodectionID" json:"prodectionID,omitempty"`
}

func (m *BuyTestRequest) Reset()                    { *m = BuyTestRequest{} }
func (m *BuyTestRequest) String() string            { return proto.CompactTextString(m) }
func (*BuyTestRequest) ProtoMessage()               {}
func (*BuyTestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BuyTestRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BuyTestRequest) GetProdectionID() int64 {
	if m != nil {
		return m.ProdectionID
	}
	return 0
}

type BuyTestReply struct {
	ExpiresDate int64  `protobuf:"varint,1,opt,name=expiresDate" json:"expiresDate,omitempty"`
	Total       int64  `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Used        int64  `protobuf:"varint,3,opt,name=used" json:"used,omitempty"`
	Error       string `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *BuyTestReply) Reset()                    { *m = BuyTestReply{} }
func (m *BuyTestReply) String() string            { return proto.CompactTextString(m) }
func (*BuyTestReply) ProtoMessage()               {}
func (*BuyTestReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BuyTestReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

func (m *BuyTestReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *BuyTestReply) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *BuyTestReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetGoogleAdRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetGoogleAdRequest) Reset()                    { *m = GetGoogleAdRequest{} }
func (m *GetGoogleAdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetGoogleAdRequest) ProtoMessage()               {}
func (*GetGoogleAdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetGoogleAdRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetGoogleAdReply struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetGoogleAdReply) Reset()                    { *m = GetGoogleAdReply{} }
func (m *GetGoogleAdReply) String() string            { return proto.CompactTextString(m) }
func (*GetGoogleAdReply) ProtoMessage()               {}
func (*GetGoogleAdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetGoogleAdReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserInfoRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetUserInfoRequest) Reset()                    { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()               {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetUserInfoRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetUserInfoReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *GetUserInfoReply) Reset()                    { *m = GetUserInfoReply{} }
func (m *GetUserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoReply) ProtoMessage()               {}
func (*GetUserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetUserInfoReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetConfigNewReply struct {
	Config string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *GetConfigNewReply) Reset()                    { *m = GetConfigNewReply{} }
func (m *GetConfigNewReply) String() string            { return proto.CompactTextString(m) }
func (*GetConfigNewReply) ProtoMessage()               {}
func (*GetConfigNewReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetConfigNewReply) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "android.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "android.LoginReply")
	proto.RegisterType((*GetConfigRequest)(nil), "android.GetConfigRequest")
	proto.RegisterType((*GetConfigReply)(nil), "android.GetConfigReply")
	proto.RegisterType((*KeepaliveRequest)(nil), "android.KeepaliveRequest")
	proto.RegisterType((*KeepaliveReply)(nil), "android.KeepaliveReply")
	proto.RegisterType((*GetProdectionListRequest)(nil), "android.GetProdectionListRequest")
	proto.RegisterType((*Prodection)(nil), "android.Prodection")
	proto.RegisterType((*GetProdectionListReply)(nil), "android.GetProdectionListReply")
	proto.RegisterType((*BuyTestRequest)(nil), "android.BuyTestRequest")
	proto.RegisterType((*BuyTestReply)(nil), "android.BuyTestReply")
	proto.RegisterType((*GetGoogleAdRequest)(nil), "android.GetGoogleAdRequest")
	proto.RegisterType((*GetGoogleAdReply)(nil), "android.GetGoogleAdReply")
	proto.RegisterType((*GetUserInfoRequest)(nil), "android.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoReply)(nil), "android.GetUserInfoReply")
	proto.RegisterType((*GetConfigNewReply)(nil), "android.GetConfigNewReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Android service

type AndroidClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error)
	GetConfigNew(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigNewReply, error)
	Keepalive(ctx context.Context, in *KeepaliveRequest, opts ...grpc.CallOption) (*KeepaliveReply, error)
	GetProdectionList(ctx context.Context, in *GetProdectionListRequest, opts ...grpc.CallOption) (*GetProdectionListReply, error)
	BuyTest(ctx context.Context, in *BuyTestRequest, opts ...grpc.CallOption) (*BuyTestReply, error)
	GetGoogleAd(ctx context.Context, in *GetGoogleAdRequest, opts ...grpc.CallOption) (*GetGoogleAdReply, error)
	GetGoogleAdFree(ctx context.Context, in *GetGoogleAdRequest, opts ...grpc.CallOption) (*GetGoogleAdReply, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error)
}

type androidClient struct {
	cc *grpc.ClientConn
}

func NewAndroidClient(cc *grpc.ClientConn) AndroidClient {
	return &androidClient{cc}
}

func (c *androidClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/android.Android/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error) {
	out := new(GetConfigReply)
	err := grpc.Invoke(ctx, "/android.Android/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetConfigNew(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigNewReply, error) {
	out := new(GetConfigNewReply)
	err := grpc.Invoke(ctx, "/android.Android/GetConfigNew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) Keepalive(ctx context.Context, in *KeepaliveRequest, opts ...grpc.CallOption) (*KeepaliveReply, error) {
	out := new(KeepaliveReply)
	err := grpc.Invoke(ctx, "/android.Android/Keepalive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetProdectionList(ctx context.Context, in *GetProdectionListRequest, opts ...grpc.CallOption) (*GetProdectionListReply, error) {
	out := new(GetProdectionListReply)
	err := grpc.Invoke(ctx, "/android.Android/GetProdectionList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) BuyTest(ctx context.Context, in *BuyTestRequest, opts ...grpc.CallOption) (*BuyTestReply, error) {
	out := new(BuyTestReply)
	err := grpc.Invoke(ctx, "/android.Android/BuyTest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetGoogleAd(ctx context.Context, in *GetGoogleAdRequest, opts ...grpc.CallOption) (*GetGoogleAdReply, error) {
	out := new(GetGoogleAdReply)
	err := grpc.Invoke(ctx, "/android.Android/GetGoogleAd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetGoogleAdFree(ctx context.Context, in *GetGoogleAdRequest, opts ...grpc.CallOption) (*GetGoogleAdReply, error) {
	out := new(GetGoogleAdReply)
	err := grpc.Invoke(ctx, "/android.Android/GetGoogleAdFree", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *androidClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := grpc.Invoke(ctx, "/android.Android/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Android service

type AndroidServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigReply, error)
	GetConfigNew(context.Context, *GetConfigRequest) (*GetConfigNewReply, error)
	Keepalive(context.Context, *KeepaliveRequest) (*KeepaliveReply, error)
	GetProdectionList(context.Context, *GetProdectionListRequest) (*GetProdectionListReply, error)
	BuyTest(context.Context, *BuyTestRequest) (*BuyTestReply, error)
	GetGoogleAd(context.Context, *GetGoogleAdRequest) (*GetGoogleAdReply, error)
	GetGoogleAdFree(context.Context, *GetGoogleAdRequest) (*GetGoogleAdReply, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
}

func RegisterAndroidServer(s *grpc.Server, srv AndroidServer) {
	s.RegisterService(&_Android_serviceDesc, srv)
}

func _Android_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetConfigNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetConfigNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetConfigNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetConfigNew(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_Keepalive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepaliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).Keepalive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/Keepalive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).Keepalive(ctx, req.(*KeepaliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetProdectionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProdectionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetProdectionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetProdectionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetProdectionList(ctx, req.(*GetProdectionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_BuyTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).BuyTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/BuyTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).BuyTest(ctx, req.(*BuyTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetGoogleAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetGoogleAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetGoogleAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetGoogleAd(ctx, req.(*GetGoogleAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetGoogleAdFree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetGoogleAdFree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetGoogleAdFree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetGoogleAdFree(ctx, req.(*GetGoogleAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Android_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndroidServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/android.Android/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndroidServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Android_serviceDesc = grpc.ServiceDesc{
	ServiceName: "android.Android",
	HandlerType: (*AndroidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Android_Login_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Android_GetConfig_Handler,
		},
		{
			MethodName: "GetConfigNew",
			Handler:    _Android_GetConfigNew_Handler,
		},
		{
			MethodName: "Keepalive",
			Handler:    _Android_Keepalive_Handler,
		},
		{
			MethodName: "GetProdectionList",
			Handler:    _Android_GetProdectionList_Handler,
		},
		{
			MethodName: "BuyTest",
			Handler:    _Android_BuyTest_Handler,
		},
		{
			MethodName: "GetGoogleAd",
			Handler:    _Android_GetGoogleAd_Handler,
		},
		{
			MethodName: "GetGoogleAdFree",
			Handler:    _Android_GetGoogleAdFree_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Android_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "android/android.proto",
}

func init() { proto.RegisterFile("android/android.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x4f, 0xe3, 0x46,
	0x10, 0xaf, 0x1d, 0x42, 0xc8, 0x10, 0xa5, 0xb0, 0x90, 0x60, 0x0c, 0xb4, 0x61, 0x55, 0x55, 0x11,
	0x15, 0x84, 0xc0, 0x4b, 0xd5, 0xf6, 0x05, 0x1a, 0x35, 0x4a, 0x8b, 0x2a, 0xea, 0xb6, 0x0f, 0x77,
	0x0f, 0x27, 0x39, 0xf1, 0x12, 0x56, 0x04, 0xaf, 0xb1, 0x37, 0xfc, 0xb9, 0x87, 0x93, 0xee, 0x74,
	0xcf, 0x77, 0x0f, 0xf7, 0xd1, 0xee, 0x2b, 0xdc, 0x07, 0x39, 0xed, 0x7a, 0xed, 0xac, 0x4d, 0x12,
	0xa4, 0x7b, 0xb8, 0x27, 0x3c, 0xb3, 0xbf, 0x99, 0xdf, 0x6f, 0x66, 0x98, 0x51, 0xa0, 0xe6, 0xfa,
	0x5e, 0xc8, 0xa8, 0xd7, 0x52, 0x7f, 0x0f, 0x82, 0x90, 0x71, 0x86, 0x4a, 0xca, 0xb4, 0xb7, 0x87,
	0x8c, 0x0d, 0x47, 0xa4, 0xe5, 0x06, 0xb4, 0xe5, 0xfa, 0x3e, 0xe3, 0x2e, 0xa7, 0xcc, 0x8f, 0x62,
	0x18, 0xfe, 0x0d, 0x2a, 0x67, 0x6c, 0x48, 0x7d, 0x87, 0xdc, 0x8c, 0x49, 0xc4, 0x11, 0x82, 0x85,
	0xf1, 0x98, 0x7a, 0x96, 0xd1, 0x30, 0x9a, 0x65, 0x47, 0x7e, 0x23, 0x0b, 0x4a, 0xb7, 0x24, 0x8c,
	0x28, 0xf3, 0x2d, 0x53, 0xba, 0x13, 0x13, 0xbf, 0x35, 0x00, 0x54, 0x78, 0x30, 0x7a, 0x40, 0xeb,
	0x50, 0xe4, 0xec, 0x8a, 0xf8, 0x2a, 0x3a, 0x36, 0x50, 0x03, 0x96, 0xc9, 0x7d, 0x40, 0x43, 0x12,
	0x75, 0x5c, 0x4e, 0x64, 0x8a, 0x82, 0xa3, 0xbb, 0xe2, 0x38, 0xee, 0x8e, 0xac, 0x82, 0x7c, 0x8b,
	0x0d, 0x29, 0x25, 0x22, 0x9e, 0xb5, 0x20, 0x9d, 0xf2, 0x5b, 0x20, 0x49, 0x18, 0xb2, 0xd0, 0x2a,
	0xc7, 0x0c, 0xd2, 0xc0, 0x4d, 0x58, 0xe9, 0x12, 0xfe, 0x3b, 0xf3, 0x2f, 0xe8, 0x30, 0x29, 0x64,
	0xaa, 0x16, 0xfc, 0x12, 0xaa, 0x1a, 0x52, 0x68, 0xae, 0x82, 0xd9, 0x3b, 0x57, 0x20, 0xb3, 0x77,
	0x2e, 0x58, 0x03, 0x16, 0x72, 0x55, 0xa9, 0xfc, 0x46, 0x75, 0x58, 0xbc, 0x26, 0xfc, 0x92, 0x79,
	0x52, 0x60, 0xd9, 0x51, 0x96, 0xf0, 0x07, 0x6e, 0x14, 0xdd, 0xc5, 0x1a, 0xcb, 0x8e, 0xb2, 0x66,
	0xa8, 0xfc, 0x19, 0x56, 0xfe, 0x22, 0x24, 0x70, 0x47, 0xf4, 0x96, 0xcc, 0x55, 0x29, 0x34, 0x85,
	0xf7, 0xaa, 0x51, 0x66, 0x78, 0x8f, 0xdf, 0x19, 0x50, 0xd5, 0x42, 0x85, 0x6c, 0x1b, 0x96, 0x7c,
	0x42, 0xbc, 0x7f, 0x39, 0x0b, 0x64, 0xec, 0x92, 0x93, 0xda, 0x5f, 0xa1, 0xe1, 0x87, 0x60, 0x75,
	0x09, 0x3f, 0x0f, 0x99, 0x47, 0x06, 0xe2, 0xbf, 0xe9, 0x8c, 0x46, 0x7c, 0x7e, 0xe3, 0x5f, 0x01,
	0x4c, 0xe0, 0xb2, 0xe9, 0x1d, 0x09, 0x28, 0x38, 0x66, 0xaf, 0x23, 0x98, 0x39, 0xbd, 0x4e, 0xa4,
	0xca, 0xef, 0x19, 0x1a, 0xd7, 0xa1, 0x18, 0x84, 0x74, 0x40, 0x94, 0xc8, 0xd8, 0x10, 0x15, 0x7b,
	0x24, 0x1a, 0x84, 0x34, 0x10, 0xe9, 0xad, 0xa2, 0x64, 0xd6, 0x5d, 0xf8, 0x0a, 0xea, 0x53, 0x14,
	0x8b, 0x4e, 0xfe, 0x0a, 0xd5, 0x20, 0xe3, 0xb6, 0x8c, 0x46, 0xa1, 0xb9, 0x7c, 0xb4, 0x76, 0x90,
	0x2c, 0xd4, 0x24, 0xca, 0xc9, 0x41, 0x67, 0xb4, 0xe7, 0x4f, 0xa8, 0x9e, 0x8e, 0x1f, 0xfe, 0x23,
	0x4f, 0x34, 0x05, 0x61, 0xa8, 0x4c, 0xf2, 0xf5, 0x3a, 0xaa, 0xfc, 0x8c, 0x0f, 0x07, 0x50, 0x49,
	0x73, 0x09, 0xb9, 0xb9, 0xe1, 0x1a, 0x73, 0x86, 0x6b, 0x4e, 0x1b, 0x6e, 0xe1, 0xc9, 0xe1, 0xee,
	0x01, 0xea, 0x12, 0xde, 0x95, 0x57, 0xe3, 0xc4, 0x9b, 0x3f, 0x56, 0x2c, 0x37, 0x6f, 0x82, 0x55,
	0x1b, 0x95, 0x1e, 0x10, 0x93, 0x7a, 0x2a, 0xdf, 0xff, 0x11, 0x09, 0x7b, 0xfe, 0x05, 0x9b, 0x9f,
	0x6f, 0x4f, 0xe6, 0x9b, 0x60, 0x45, 0xbe, 0x3a, 0x2c, 0x46, 0xdc, 0xe5, 0xe3, 0x48, 0x41, 0x95,
	0x85, 0x7f, 0x82, 0xd5, 0x74, 0x97, 0xff, 0x26, 0x77, 0x29, 0x78, 0x20, 0x3d, 0x09, 0x38, 0xb6,
	0x8e, 0xde, 0x97, 0xa0, 0x74, 0x12, 0xcf, 0x13, 0xfd, 0x03, 0x45, 0x79, 0xb4, 0x50, 0x2d, 0x1d,
	0xb1, 0x7e, 0x03, 0xed, 0xb5, 0xbc, 0x3b, 0x18, 0x3d, 0xe0, 0xed, 0x37, 0x1f, 0x3f, 0x7d, 0x30,
	0xeb, 0x78, 0xb5, 0x75, 0xdb, 0x4e, 0x4e, 0x6d, 0x6b, 0x24, 0xde, 0x7f, 0x31, 0xf6, 0xd0, 0x0b,
	0x28, 0xa7, 0x5a, 0xd0, 0x66, 0x1a, 0x9f, 0xbf, 0x4a, 0xf6, 0xc6, 0xb4, 0x27, 0x91, 0x7e, 0x47,
	0xa6, 0xdf, 0xc0, 0x48, 0x4f, 0x1f, 0x6b, 0x17, 0xf9, 0x2f, 0xa1, 0xa2, 0xd7, 0x3a, 0x8f, 0xc2,
	0x7e, 0xfc, 0x94, 0x74, 0x07, 0xef, 0x4a, 0x96, 0x2d, 0x5c, 0xd7, 0x59, 0x86, 0x84, 0xef, 0x4f,
	0x98, 0x5c, 0x28, 0xa7, 0xa7, 0x46, 0xa3, 0xc9, 0x5f, 0x2e, 0xad, 0x92, 0xec, 0x65, 0xc2, 0x0d,
	0xc9, 0x61, 0xe3, 0x9a, 0xce, 0x71, 0x95, 0x60, 0x04, 0xc5, 0x6b, 0x43, 0x4e, 0x2e, 0xbb, 0x8c,
	0x68, 0x57, 0xd7, 0x3d, 0xf5, 0xb4, 0xd8, 0xdf, 0xcf, 0x83, 0x08, 0xee, 0xa6, 0xe4, 0xc6, 0x78,
	0x27, 0x57, 0x5f, 0x16, 0x2b, 0x34, 0x3c, 0x83, 0x92, 0x5a, 0x2b, 0x34, 0xa9, 0x24, 0xbb, 0xb4,
	0x76, 0xed, 0xf1, 0x83, 0x20, 0xf9, 0x4e, 0x92, 0x58, 0x78, 0x4d, 0x27, 0xe9, 0xc7, 0x08, 0x91,
	0x9a, 0xc2, 0xb2, 0xb6, 0x13, 0x68, 0x4b, 0x17, 0x9d, 0xdb, 0x2a, 0x7b, 0x73, 0xfa, 0xa3, 0xa0,
	0xc1, 0x92, 0x66, 0x1b, 0x6f, 0xe4, 0x6a, 0x49, 0x50, 0x82, 0xea, 0x06, 0xbe, 0xd5, 0xe2, 0xfe,
	0x08, 0x09, 0xf9, 0x62, 0xba, 0x1f, 0x25, 0x5d, 0x03, 0x6f, 0xcd, 0xa0, 0x13, 0xc9, 0x05, 0x65,
	0x57, 0x56, 0x97, 0x6c, 0x68, 0x96, 0x2e, 0xb7, 0xe3, 0x59, 0xba, 0xcc, 0x52, 0xe3, 0x6f, 0x4e,
	0x7f, 0x78, 0x8e, 0x87, 0x94, 0x5f, 0x8e, 0xfb, 0x07, 0x03, 0x76, 0xdd, 0x1a, 0x1f, 0x1e, 0x1f,
	0xb6, 0x8f, 0xdb, 0xed, 0xa3, 0x56, 0x14, 0xed, 0x07, 0xfd, 0x44, 0x41, 0x7f, 0x51, 0xfe, 0x4c,
	0x39, 0xfe, 0x1c, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x53, 0x5c, 0x25, 0xe6, 0x08, 0x00, 0x00,
}
