// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: login_service.proto

package login_service_v1

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

var File_login_service_proto protoreflect.FileDescriptor

var file_login_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x29, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_login_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_login_service_proto_goTypes = []interface{}{
	(*CaptchaMessage)(nil),  // 0: login.service.v1.CaptchaMessage
	(*CaptchaResponse)(nil), // 1: login.service.v1.CaptchaResponse
}
var file_login_service_proto_depIdxs = []int32{
	0, // 0: login.service.v1.LoginService.GetCaptcha:input_type -> login.service.v1.CaptchaMessage
	1, // 1: login.service.v1.LoginService.GetCaptcha:output_type -> login.service.v1.CaptchaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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