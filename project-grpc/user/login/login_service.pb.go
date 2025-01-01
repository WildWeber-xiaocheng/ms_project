// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: login_service.proto

package login

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CaptchaMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *CaptchaMessage) Reset() {
	*x = CaptchaMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaMessage) ProtoMessage() {}

func (x *CaptchaMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaMessage.ProtoReflect.Descriptor instead.
func (*CaptchaMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{0}
}

func (x *CaptchaMessage) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type CaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CaptchaResponse) Reset() {
	*x = CaptchaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaResponse) ProtoMessage() {}

func (x *CaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaResponse.ProtoReflect.Descriptor instead.
func (*CaptchaResponse) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{1}
}

func (x *CaptchaResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RegisterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Captcha  string `protobuf:"bytes,5,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *RegisterMessage) Reset() {
	*x = RegisterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMessage) ProtoMessage() {}

func (x *RegisterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMessage.ProtoReflect.Descriptor instead.
func (*RegisterMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterMessage) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterMessage) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{3}
}

type LoginMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginMessage) Reset() {
	*x = LoginMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMessage) ProtoMessage() {}

func (x *LoginMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMessage.ProtoReflect.Descriptor instead.
func (*LoginMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{4}
}

func (x *LoginMessage) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginMessage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member           *MemberMessage         `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	OrganizationList []*OrganizationMessage `protobuf:"bytes,2,rep,name=organizationList,proto3" json:"organizationList,omitempty"`
	TokenList        *TokenMessage          `protobuf:"bytes,3,opt,name=tokenList,proto3" json:"tokenList,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetMember() *MemberMessage {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *LoginResponse) GetOrganizationList() []*OrganizationMessage {
	if x != nil {
		return x.OrganizationList
	}
	return nil
}

func (x *LoginResponse) GetTokenList() *TokenMessage {
	if x != nil {
		return x.TokenList
	}
	return nil
}

type MemberMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile        string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Realname      string `protobuf:"bytes,4,opt,name=realname,proto3" json:"realname,omitempty"`
	Account       string `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
	Status        int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	LastLoginTime string `protobuf:"bytes,7,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	Address       string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Province      int32  `protobuf:"varint,9,opt,name=province,proto3" json:"province,omitempty"`
	City          int32  `protobuf:"varint,10,opt,name=city,proto3" json:"city,omitempty"`
	Area          int32  `protobuf:"varint,11,opt,name=area,proto3" json:"area,omitempty"`
	Email         string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Code          string `protobuf:"bytes,13,opt,name=code,proto3" json:"code,omitempty"`
	CreateTime    string `protobuf:"bytes,14,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *MemberMessage) Reset() {
	*x = MemberMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberMessage) ProtoMessage() {}

func (x *MemberMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberMessage.ProtoReflect.Descriptor instead.
func (*MemberMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{6}
}

func (x *MemberMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberMessage) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MemberMessage) GetRealname() string {
	if x != nil {
		return x.Realname
	}
	return ""
}

func (x *MemberMessage) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *MemberMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberMessage) GetLastLoginTime() string {
	if x != nil {
		return x.LastLoginTime
	}
	return ""
}

func (x *MemberMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberMessage) GetProvince() int32 {
	if x != nil {
		return x.Province
	}
	return 0
}

func (x *MemberMessage) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

func (x *MemberMessage) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *MemberMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MemberMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type OrganizationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MemberId    int64  `protobuf:"varint,5,opt,name=memberId,proto3" json:"memberId,omitempty"`
	CreateTime  string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Personal    int32  `protobuf:"varint,7,opt,name=personal,proto3" json:"personal,omitempty"`
	Address     string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Province    int32  `protobuf:"varint,9,opt,name=province,proto3" json:"province,omitempty"`
	City        int32  `protobuf:"varint,10,opt,name=city,proto3" json:"city,omitempty"`
	Area        int32  `protobuf:"varint,11,opt,name=area,proto3" json:"area,omitempty"`
	Code        string `protobuf:"bytes,12,opt,name=code,proto3" json:"code,omitempty"`
	OwnerCode   string `protobuf:"bytes,13,opt,name=ownerCode,proto3" json:"ownerCode,omitempty"`
}

func (x *OrganizationMessage) Reset() {
	*x = OrganizationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationMessage) ProtoMessage() {}

func (x *OrganizationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationMessage.ProtoReflect.Descriptor instead.
func (*OrganizationMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{7}
}

func (x *OrganizationMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationMessage) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *OrganizationMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrganizationMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *OrganizationMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *OrganizationMessage) GetPersonal() int32 {
	if x != nil {
		return x.Personal
	}
	return 0
}

func (x *OrganizationMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationMessage) GetProvince() int32 {
	if x != nil {
		return x.Province
	}
	return 0
}

func (x *OrganizationMessage) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

func (x *OrganizationMessage) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *OrganizationMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OrganizationMessage) GetOwnerCode() string {
	if x != nil {
		return x.OwnerCode
	}
	return ""
}

type TokenMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken    string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken   string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	TokenType      string `protobuf:"bytes,3,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	AccessTokenExp int64  `protobuf:"varint,4,opt,name=accessTokenExp,proto3" json:"accessTokenExp,omitempty"`
}

func (x *TokenMessage) Reset() {
	*x = TokenMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMessage) ProtoMessage() {}

func (x *TokenMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMessage.ProtoReflect.Descriptor instead.
func (*TokenMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{8}
}

func (x *TokenMessage) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenMessage) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *TokenMessage) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *TokenMessage) GetAccessTokenExp() int64 {
	if x != nil {
		return x.AccessTokenExp
	}
	return 0
}

type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemId int64 `protobuf:"varint,1,opt,name=memId,proto3" json:"memId,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{9}
}

func (x *UserMessage) GetMemId() int64 {
	if x != nil {
		return x.MemId
	}
	return 0
}

type OrgListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationList []*OrganizationMessage `protobuf:"bytes,1,rep,name=organizationList,proto3" json:"organizationList,omitempty"`
}

func (x *OrgListResponse) Reset() {
	*x = OrgListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgListResponse) ProtoMessage() {}

func (x *OrgListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgListResponse.ProtoReflect.Descriptor instead.
func (*OrgListResponse) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{10}
}

func (x *OrgListResponse) GetOrganizationList() []*OrganizationMessage {
	if x != nil {
		return x.OrganizationList
	}
	return nil
}

var File_login_service_proto protoreflect.FileDescriptor

var file_login_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x51, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xe7, 0x02, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xdb, 0x02, 0x0a, 0x13, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x78, 0x70, 0x22, 0x23, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x0f, 0x4f, 0x72,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0xa7, 0x03, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12,
	0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x09, 0x4d, 0x79, 0x4f,
	0x72, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_service_proto_rawDescOnce sync.Once
	file_login_service_proto_rawDescData = file_login_service_proto_rawDesc
)

func file_login_service_proto_rawDescGZIP() []byte {
	file_login_service_proto_rawDescOnce.Do(func() {
		file_login_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_service_proto_rawDescData)
	})
	return file_login_service_proto_rawDescData
}

var file_login_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_login_service_proto_goTypes = []interface{}{
	(*CaptchaMessage)(nil),      // 0: login.service.v1.CaptchaMessage
	(*CaptchaResponse)(nil),     // 1: login.service.v1.CaptchaResponse
	(*RegisterMessage)(nil),     // 2: login.service.v1.RegisterMessage
	(*RegisterResponse)(nil),    // 3: login.service.v1.RegisterResponse
	(*LoginMessage)(nil),        // 4: login.service.v1.LoginMessage
	(*LoginResponse)(nil),       // 5: login.service.v1.LoginResponse
	(*MemberMessage)(nil),       // 6: login.service.v1.MemberMessage
	(*OrganizationMessage)(nil), // 7: login.service.v1.OrganizationMessage
	(*TokenMessage)(nil),        // 8: login.service.v1.TokenMessage
	(*UserMessage)(nil),         // 9: login.service.v1.UserMessage
	(*OrgListResponse)(nil),     // 10: login.service.v1.OrgListResponse
}
var file_login_service_proto_depIdxs = []int32{
	6,  // 0: login.service.v1.LoginResponse.member:type_name -> login.service.v1.MemberMessage
	7,  // 1: login.service.v1.LoginResponse.organizationList:type_name -> login.service.v1.OrganizationMessage
	8,  // 2: login.service.v1.LoginResponse.tokenList:type_name -> login.service.v1.TokenMessage
	7,  // 3: login.service.v1.OrgListResponse.organizationList:type_name -> login.service.v1.OrganizationMessage
	0,  // 4: login.service.v1.LoginService.GetCaptcha:input_type -> login.service.v1.CaptchaMessage
	2,  // 5: login.service.v1.LoginService.Register:input_type -> login.service.v1.RegisterMessage
	4,  // 6: login.service.v1.LoginService.Login:input_type -> login.service.v1.LoginMessage
	4,  // 7: login.service.v1.LoginService.TokenVerify:input_type -> login.service.v1.LoginMessage
	9,  // 8: login.service.v1.LoginService.MyOrgList:input_type -> login.service.v1.UserMessage
	1,  // 9: login.service.v1.LoginService.GetCaptcha:output_type -> login.service.v1.CaptchaResponse
	3,  // 10: login.service.v1.LoginService.Register:output_type -> login.service.v1.RegisterResponse
	5,  // 11: login.service.v1.LoginService.Login:output_type -> login.service.v1.LoginResponse
	5,  // 12: login.service.v1.LoginService.TokenVerify:output_type -> login.service.v1.LoginResponse
	10, // 13: login.service.v1.LoginService.MyOrgList:output_type -> login.service.v1.OrgListResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_login_service_proto_init() }
func file_login_service_proto_init() {
	if File_login_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_service_proto_goTypes,
		DependencyIndexes: file_login_service_proto_depIdxs,
		MessageInfos:      file_login_service_proto_msgTypes,
	}.Build()
	File_login_service_proto = out.File
	file_login_service_proto_rawDesc = nil
	file_login_service_proto_goTypes = nil
	file_login_service_proto_depIdxs = nil
}
