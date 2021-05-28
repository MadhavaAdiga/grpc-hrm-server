// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: permission_message.proto

package hrm

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

// todo: update enum
type PermissionType int32

const (
	PermissionType_UNKNOWN          PermissionType = 0
	PermissionType_FULLACCESS       PermissionType = 1
	PermissionType_RESTRICTEDACCESS PermissionType = 2
	PermissionType_DENIED           PermissionType = 3
)

// Enum value maps for PermissionType.
var (
	PermissionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "FULLACCESS",
		2: "RESTRICTEDACCESS",
		3: "DENIED",
	}
	PermissionType_value = map[string]int32{
		"UNKNOWN":          0,
		"FULLACCESS":       1,
		"RESTRICTEDACCESS": 2,
		"DENIED":           3,
	}
)

func (x PermissionType) Enum() *PermissionType {
	p := new(PermissionType)
	*p = x
	return p
}

func (x PermissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_permission_message_proto_enumTypes[0].Descriptor()
}

func (PermissionType) Type() protoreflect.EnumType {
	return &file_permission_message_proto_enumTypes[0]
}

func (x PermissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionType.Descriptor instead.
func (PermissionType) EnumDescriptor() ([]byte, []int) {
	return file_permission_message_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   PermissionType `protobuf:"varint,2,opt,name=type,proto3,enum=PermissionType" json:"type,omitempty"`
	Active bool           `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"` // role?
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_permission_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_permission_message_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetType() PermissionType {
	if x != nil {
		return x.Type
	}
	return PermissionType_UNKNOWN
}

func (x *Permission) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

var File_permission_message_proto protoreflect.FileDescriptor

var file_permission_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0a, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2a, 0x4f, 0x0a, 0x0e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x55, 0x4c, 0x4c,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x54,
	0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x68,
	0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_message_proto_rawDescOnce sync.Once
	file_permission_message_proto_rawDescData = file_permission_message_proto_rawDesc
)

func file_permission_message_proto_rawDescGZIP() []byte {
	file_permission_message_proto_rawDescOnce.Do(func() {
		file_permission_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_message_proto_rawDescData)
	})
	return file_permission_message_proto_rawDescData
}

var file_permission_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_permission_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_permission_message_proto_goTypes = []interface{}{
	(PermissionType)(0), // 0: PermissionType
	(*Permission)(nil),  // 1: Permission
}
var file_permission_message_proto_depIdxs = []int32{
	0, // 0: Permission.type:type_name -> PermissionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_permission_message_proto_init() }
func file_permission_message_proto_init() {
	if File_permission_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
			RawDescriptor: file_permission_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_message_proto_goTypes,
		DependencyIndexes: file_permission_message_proto_depIdxs,
		EnumInfos:         file_permission_message_proto_enumTypes,
		MessageInfos:      file_permission_message_proto_msgTypes,
	}.Build()
	File_permission_message_proto = out.File
	file_permission_message_proto_rawDesc = nil
	file_permission_message_proto_goTypes = nil
	file_permission_message_proto_depIdxs = nil
}
