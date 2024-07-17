// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: registration.proto

package api

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

type RegistrationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels      []*Label           `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Resources   map[string]float64 `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	BindAddress string             `protobuf:"bytes,3,opt,name=bindAddress,proto3" json:"bindAddress,omitempty"`
}

func (x *RegistrationReq) Reset() {
	*x = RegistrationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationReq) ProtoMessage() {}

func (x *RegistrationReq) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationReq.ProtoReflect.Descriptor instead.
func (*RegistrationReq) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationReq) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RegistrationReq) GetResources() map[string]float64 {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *RegistrationReq) GetBindAddress() string {
	if x != nil {
		return x.BindAddress
	}
	return ""
}

type RegistrationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
}

func (x *RegistrationResp) Reset() {
	*x = RegistrationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResp) ProtoMessage() {}

func (x *RegistrationResp) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResp.ProtoReflect.Descriptor instead.
func (*RegistrationResp) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationResp) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_registration_proto protoreflect.FileDescriptor

var file_registration_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x61, 0x67,
	0x6e, 0x65, 0x74, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x31, 0x32, 0x73, 0x2f,
	0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x61, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registration_proto_rawDescOnce sync.Once
	file_registration_proto_rawDescData = file_registration_proto_rawDesc
)

func file_registration_proto_rawDescGZIP() []byte {
	file_registration_proto_rawDescOnce.Do(func() {
		file_registration_proto_rawDescData = protoimpl.X.CompressGZIP(file_registration_proto_rawDescData)
	})
	return file_registration_proto_rawDescData
}

var file_registration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_registration_proto_goTypes = []interface{}{
	(*RegistrationReq)(nil),  // 0: proto.RegistrationReq
	(*RegistrationResp)(nil), // 1: proto.RegistrationResp
	nil,                      // 2: proto.RegistrationReq.ResourcesEntry
	(*Label)(nil),            // 3: proto.Label
}
var file_registration_proto_depIdxs = []int32{
	3, // 0: proto.RegistrationReq.labels:type_name -> proto.Label
	2, // 1: proto.RegistrationReq.resources:type_name -> proto.RegistrationReq.ResourcesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_registration_proto_init() }
func file_registration_proto_init() {
	if File_registration_proto != nil {
		return
	}
	file_magnetar_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_registration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationReq); i {
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
		file_registration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResp); i {
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
			RawDescriptor: file_registration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registration_proto_goTypes,
		DependencyIndexes: file_registration_proto_depIdxs,
		MessageInfos:      file_registration_proto_msgTypes,
	}.Build()
	File_registration_proto = out.File
	file_registration_proto_rawDesc = nil
	file_registration_proto_goTypes = nil
	file_registration_proto_depIdxs = nil
}
