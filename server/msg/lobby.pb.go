// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lobby.proto

package msg

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

// 默认消息返回协议
type DefaultReault struct {
	RetResult            int32    `protobuf:"varint,1,opt,name=RetResult,proto3" json:"RetResult,omitempty"`
	ErrorInfo            string   `protobuf:"bytes,2,opt,name=ErrorInfo,proto3" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultReault) Reset()         { *m = DefaultReault{} }
func (m *DefaultReault) String() string { return proto.CompactTextString(m) }
func (*DefaultReault) ProtoMessage()    {}
func (*DefaultReault) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{0}
}

func (m *DefaultReault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultReault.Unmarshal(m, b)
}
func (m *DefaultReault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultReault.Marshal(b, m, deterministic)
}
func (m *DefaultReault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultReault.Merge(m, src)
}
func (m *DefaultReault) XXX_Size() int {
	return xxx_messageInfo_DefaultReault.Size(m)
}
func (m *DefaultReault) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultReault.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultReault proto.InternalMessageInfo

func (m *DefaultReault) GetRetResult() int32 {
	if m != nil {
		return m.RetResult
	}
	return 0
}

func (m *DefaultReault) GetErrorInfo() string {
	if m != nil {
		return m.ErrorInfo
	}
	return ""
}

// 用户登陆协议
type UserLogin struct {
	LoginName            string   `protobuf:"bytes,1,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	LoginPW              string   `protobuf:"bytes,2,opt,name=LoginPW,proto3" json:"LoginPW,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLogin) Reset()         { *m = UserLogin{} }
func (m *UserLogin) String() string { return proto.CompactTextString(m) }
func (*UserLogin) ProtoMessage()    {}
func (*UserLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{1}
}

func (m *UserLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLogin.Unmarshal(m, b)
}
func (m *UserLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLogin.Marshal(b, m, deterministic)
}
func (m *UserLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLogin.Merge(m, src)
}
func (m *UserLogin) XXX_Size() int {
	return xxx_messageInfo_UserLogin.Size(m)
}
func (m *UserLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLogin.DiscardUnknown(m)
}

var xxx_messageInfo_UserLogin proto.InternalMessageInfo

func (m *UserLogin) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserLogin) GetLoginPW() string {
	if m != nil {
		return m.LoginPW
	}
	return ""
}

type UserLogin_Result struct {
	DefaultReault        *DefaultReault `protobuf:"bytes,1,opt,name=defaultReault,proto3" json:"defaultReault,omitempty"`
	Token                string         `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserLogin_Result) Reset()         { *m = UserLogin_Result{} }
func (m *UserLogin_Result) String() string { return proto.CompactTextString(m) }
func (*UserLogin_Result) ProtoMessage()    {}
func (*UserLogin_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{1, 0}
}

func (m *UserLogin_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLogin_Result.Unmarshal(m, b)
}
func (m *UserLogin_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLogin_Result.Marshal(b, m, deterministic)
}
func (m *UserLogin_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLogin_Result.Merge(m, src)
}
func (m *UserLogin_Result) XXX_Size() int {
	return xxx_messageInfo_UserLogin_Result.Size(m)
}
func (m *UserLogin_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLogin_Result.DiscardUnknown(m)
}

var xxx_messageInfo_UserLogin_Result proto.InternalMessageInfo

func (m *UserLogin_Result) GetDefaultReault() *DefaultReault {
	if m != nil {
		return m.DefaultReault
	}
	return nil
}

func (m *UserLogin_Result) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 注册协议
type UserRegister struct {
	LoginName            string   `protobuf:"bytes,1,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	LoginPW              string   `protobuf:"bytes,2,opt,name=LoginPW,proto3" json:"LoginPW,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegister) Reset()         { *m = UserRegister{} }
func (m *UserRegister) String() string { return proto.CompactTextString(m) }
func (*UserRegister) ProtoMessage()    {}
func (*UserRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{2}
}

func (m *UserRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegister.Unmarshal(m, b)
}
func (m *UserRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegister.Marshal(b, m, deterministic)
}
func (m *UserRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegister.Merge(m, src)
}
func (m *UserRegister) XXX_Size() int {
	return xxx_messageInfo_UserRegister.Size(m)
}
func (m *UserRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegister.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegister proto.InternalMessageInfo

func (m *UserRegister) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserRegister) GetLoginPW() string {
	if m != nil {
		return m.LoginPW
	}
	return ""
}

type UserRegister_Result struct {
	DefaultReault        *DefaultReault `protobuf:"bytes,1,opt,name=defaultReault,proto3" json:"defaultReault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserRegister_Result) Reset()         { *m = UserRegister_Result{} }
func (m *UserRegister_Result) String() string { return proto.CompactTextString(m) }
func (*UserRegister_Result) ProtoMessage()    {}
func (*UserRegister_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{2, 0}
}

func (m *UserRegister_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegister_Result.Unmarshal(m, b)
}
func (m *UserRegister_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegister_Result.Marshal(b, m, deterministic)
}
func (m *UserRegister_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegister_Result.Merge(m, src)
}
func (m *UserRegister_Result) XXX_Size() int {
	return xxx_messageInfo_UserRegister_Result.Size(m)
}
func (m *UserRegister_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegister_Result.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegister_Result proto.InternalMessageInfo

func (m *UserRegister_Result) GetDefaultReault() *DefaultReault {
	if m != nil {
		return m.DefaultReault
	}
	return nil
}

// 玩家有角色的情况
type UserST struct {
	UID                  string   `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	ServerID             string   `protobuf:"bytes,2,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	RoleUID              string   `protobuf:"bytes,3,opt,name=RoleUID,proto3" json:"RoleUID,omitempty"`
	RoleName             string   `protobuf:"bytes,4,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
	RoleLev              string   `protobuf:"bytes,5,opt,name=RoleLev,proto3" json:"RoleLev,omitempty"`
	Coin                 string   `protobuf:"bytes,6,opt,name=Coin,proto3" json:"Coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserST) Reset()         { *m = UserST{} }
func (m *UserST) String() string { return proto.CompactTextString(m) }
func (*UserST) ProtoMessage()    {}
func (*UserST) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{3}
}

func (m *UserST) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserST.Unmarshal(m, b)
}
func (m *UserST) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserST.Marshal(b, m, deterministic)
}
func (m *UserST) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserST.Merge(m, src)
}
func (m *UserST) XXX_Size() int {
	return xxx_messageInfo_UserST.Size(m)
}
func (m *UserST) XXX_DiscardUnknown() {
	xxx_messageInfo_UserST.DiscardUnknown(m)
}

var xxx_messageInfo_UserST proto.InternalMessageInfo

func (m *UserST) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *UserST) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *UserST) GetRoleUID() string {
	if m != nil {
		return m.RoleUID
	}
	return ""
}

func (m *UserST) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *UserST) GetRoleLev() string {
	if m != nil {
		return m.RoleLev
	}
	return ""
}

func (m *UserST) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

type UserST_Result struct {
	DefaultReault        *DefaultReault `protobuf:"bytes,1,opt,name=defaultReault,proto3" json:"defaultReault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserST_Result) Reset()         { *m = UserST_Result{} }
func (m *UserST_Result) String() string { return proto.CompactTextString(m) }
func (*UserST_Result) ProtoMessage()    {}
func (*UserST_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0edfb21975fe5, []int{3, 0}
}

func (m *UserST_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserST_Result.Unmarshal(m, b)
}
func (m *UserST_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserST_Result.Marshal(b, m, deterministic)
}
func (m *UserST_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserST_Result.Merge(m, src)
}
func (m *UserST_Result) XXX_Size() int {
	return xxx_messageInfo_UserST_Result.Size(m)
}
func (m *UserST_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_UserST_Result.DiscardUnknown(m)
}

var xxx_messageInfo_UserST_Result proto.InternalMessageInfo

func (m *UserST_Result) GetDefaultReault() *DefaultReault {
	if m != nil {
		return m.DefaultReault
	}
	return nil
}

func init() {
	proto.RegisterType((*DefaultReault)(nil), "msg.DefaultReault")
	proto.RegisterType((*UserLogin)(nil), "msg.UserLogin")
	proto.RegisterType((*UserLogin_Result)(nil), "msg.UserLogin.Result")
	proto.RegisterType((*UserRegister)(nil), "msg.UserRegister")
	proto.RegisterType((*UserRegister_Result)(nil), "msg.UserRegister.Result")
	proto.RegisterType((*UserST)(nil), "msg.UserST")
	proto.RegisterType((*UserST_Result)(nil), "msg.UserST.Result")
}

func init() { proto.RegisterFile("lobby.proto", fileDescriptor_a5f0edfb21975fe5) }

var fileDescriptor_a5f0edfb21975fe5 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xd3, 0x44, 0x33, 0xb5, 0x20, 0x8b, 0x87, 0x50, 0x3c, 0x94, 0x9c, 0x7a, 0xca,
	0x41, 0x2f, 0x9e, 0x35, 0x1e, 0x82, 0x41, 0x64, 0x6b, 0xd1, 0x6b, 0x83, 0xd3, 0x10, 0x4c, 0xb2,
	0xb2, 0xbb, 0x0d, 0xf8, 0x0a, 0xbe, 0x83, 0xaf, 0xe6, 0xb3, 0xc8, 0x6c, 0x36, 0xc6, 0x5e, 0xd5,
	0x4b, 0x98, 0xff, 0xcf, 0xfe, 0xff, 0x7e, 0x03, 0x0b, 0xd3, 0x5a, 0x16, 0xc5, 0x5b, 0xf2, 0xaa,
	0xa4, 0x91, 0xdc, 0x6b, 0x74, 0x19, 0xdf, 0xc2, 0x2c, 0xc5, 0xed, 0x66, 0x57, 0x1b, 0x81, 0xf4,
	0xe5, 0x67, 0x10, 0x0a, 0x34, 0x02, 0xf5, 0xae, 0x36, 0x11, 0x5b, 0xb0, 0xa5, 0x2f, 0x46, 0x83,
	0xfe, 0xde, 0x28, 0x25, 0x55, 0xd6, 0x6e, 0x65, 0x74, 0xb0, 0x60, 0xcb, 0x50, 0x8c, 0x46, 0xfc,
	0xc1, 0x20, 0x5c, 0x6b, 0x54, 0xb9, 0x2c, 0xab, 0x96, 0xce, 0xda, 0xe1, 0x6e, 0xd3, 0xa0, 0x6d,
	0x0a, 0xc5, 0x68, 0xf0, 0x08, 0x0e, 0xad, 0xb8, 0x7f, 0x74, 0x3d, 0x83, 0x9c, 0x3f, 0x41, 0xe0,
	0x6e, 0xbb, 0x84, 0xd9, 0xf3, 0x4f, 0x38, 0xdb, 0x32, 0x3d, 0xe7, 0x49, 0xa3, 0xcb, 0x64, 0x0f,
	0x5b, 0xec, 0x1f, 0xe4, 0xa7, 0xe0, 0x1b, 0xf9, 0x82, 0xad, 0xeb, 0xee, 0x45, 0xfc, 0xce, 0xe0,
	0x98, 0xf8, 0x04, 0x96, 0x95, 0x36, 0xa8, 0x7e, 0x8d, 0x78, 0xf5, 0x77, 0xc4, 0xf8, 0x93, 0x41,
	0x40, 0x30, 0xab, 0x07, 0x7e, 0x02, 0xde, 0x3a, 0x4b, 0x1d, 0x00, 0x8d, 0x7c, 0x0e, 0x47, 0x2b,
	0x54, 0x1d, 0xaa, 0x2c, 0x75, 0x77, 0x7f, 0x6b, 0xc2, 0x12, 0xb2, 0x46, 0x4a, 0x78, 0x3d, 0x96,
	0x93, 0x94, 0xa2, 0xd1, 0x6e, 0x33, 0xe9, 0x53, 0x83, 0x1e, 0x52, 0x39, 0x76, 0x91, 0x3f, 0xa6,
	0x72, 0xec, 0x38, 0x87, 0xc9, 0xb5, 0xac, 0xda, 0x28, 0xb0, 0xb6, 0x9d, 0xff, 0x63, 0xc1, 0x22,
	0xb0, 0xcf, 0xec, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xff, 0xe0, 0x01, 0x75, 0x02, 0x00,
	0x00,
}
