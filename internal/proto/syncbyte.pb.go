// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/proto/syncbyte.proto

package proto

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

type BackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourcePath string `protobuf:"bytes,1,opt,name=SourcePath,proto3" json:"SourcePath,omitempty"`
}

func (x *BackupRequest) Reset() {
	*x = BackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_syncbyte_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRequest) ProtoMessage() {}

func (x *BackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_syncbyte_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRequest.ProtoReflect.Descriptor instead.
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_syncbyte_proto_rawDescGZIP(), []int{0}
}

func (x *BackupRequest) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

type BackupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Path       string      `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	Size       int64       `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	MD5        string      `protobuf:"bytes,4,opt,name=MD5,proto3" json:"MD5,omitempty"`
	GID        uint32      `protobuf:"varint,5,opt,name=GID,proto3" json:"GID,omitempty"`
	UID        uint32      `protobuf:"varint,6,opt,name=UID,proto3" json:"UID,omitempty"`
	Device     uint64      `protobuf:"varint,7,opt,name=Device,proto3" json:"Device,omitempty"`
	DeviceID   uint64      `protobuf:"varint,8,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	BlockSize  int64       `protobuf:"varint,9,opt,name=BlockSize,proto3" json:"BlockSize,omitempty"`
	Blocks     int64       `protobuf:"varint,10,opt,name=Blocks,proto3" json:"Blocks,omitempty"`
	AccessTime int64       `protobuf:"varint,11,opt,name=AccessTime,proto3" json:"AccessTime,omitempty"`
	ModTime    int64       `protobuf:"varint,12,opt,name=ModTime,proto3" json:"ModTime,omitempty"`
	ChangeTime int64       `protobuf:"varint,13,opt,name=ChangeTime,proto3" json:"ChangeTime,omitempty"`
	PartInfos  []*PartInfo `protobuf:"bytes,14,rep,name=PartInfos,proto3" json:"PartInfos,omitempty"`
}

func (x *BackupResponse) Reset() {
	*x = BackupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_syncbyte_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupResponse) ProtoMessage() {}

func (x *BackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_syncbyte_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupResponse.ProtoReflect.Descriptor instead.
func (*BackupResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_syncbyte_proto_rawDescGZIP(), []int{1}
}

func (x *BackupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BackupResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BackupResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BackupResponse) GetMD5() string {
	if x != nil {
		return x.MD5
	}
	return ""
}

func (x *BackupResponse) GetGID() uint32 {
	if x != nil {
		return x.GID
	}
	return 0
}

func (x *BackupResponse) GetUID() uint32 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *BackupResponse) GetDevice() uint64 {
	if x != nil {
		return x.Device
	}
	return 0
}

func (x *BackupResponse) GetDeviceID() uint64 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

func (x *BackupResponse) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *BackupResponse) GetBlocks() int64 {
	if x != nil {
		return x.Blocks
	}
	return 0
}

func (x *BackupResponse) GetAccessTime() int64 {
	if x != nil {
		return x.AccessTime
	}
	return 0
}

func (x *BackupResponse) GetModTime() int64 {
	if x != nil {
		return x.ModTime
	}
	return 0
}

func (x *BackupResponse) GetChangeTime() int64 {
	if x != nil {
		return x.ChangeTime
	}
	return 0
}

func (x *BackupResponse) GetPartInfos() []*PartInfo {
	if x != nil {
		return x.PartInfos
	}
	return nil
}

type PartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	MD5   string `protobuf:"bytes,2,opt,name=MD5,proto3" json:"MD5,omitempty"`
	Size  int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *PartInfo) Reset() {
	*x = PartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_syncbyte_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartInfo) ProtoMessage() {}

func (x *PartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_syncbyte_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartInfo.ProtoReflect.Descriptor instead.
func (*PartInfo) Descriptor() ([]byte, []int) {
	return file_internal_proto_syncbyte_proto_rawDescGZIP(), []int{2}
}

func (x *PartInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PartInfo) GetMD5() string {
	if x != nil {
		return x.MD5
	}
	return ""
}

func (x *PartInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_internal_proto_syncbyte_proto protoreflect.FileDescriptor

var file_internal_proto_syncbyte_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0xf5, 0x02, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x44, 0x35, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x44, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x47, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22,
	0x46, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x44, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x44, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x45, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x62,
	0x79, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79,
	0x6c, 0x69, 0x6e, 0x65, 0x39, 0x33, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x79, 0x74, 0x65, 0x2d,
	0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_syncbyte_proto_rawDescOnce sync.Once
	file_internal_proto_syncbyte_proto_rawDescData = file_internal_proto_syncbyte_proto_rawDesc
)

func file_internal_proto_syncbyte_proto_rawDescGZIP() []byte {
	file_internal_proto_syncbyte_proto_rawDescOnce.Do(func() {
		file_internal_proto_syncbyte_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_syncbyte_proto_rawDescData)
	})
	return file_internal_proto_syncbyte_proto_rawDescData
}

var file_internal_proto_syncbyte_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_proto_syncbyte_proto_goTypes = []interface{}{
	(*BackupRequest)(nil),  // 0: proto.BackupRequest
	(*BackupResponse)(nil), // 1: proto.BackupResponse
	(*PartInfo)(nil),       // 2: proto.PartInfo
}
var file_internal_proto_syncbyte_proto_depIdxs = []int32{
	2, // 0: proto.BackupResponse.PartInfos:type_name -> proto.PartInfo
	0, // 1: proto.Syncbyte.Backup:input_type -> proto.BackupRequest
	1, // 2: proto.Syncbyte.Backup:output_type -> proto.BackupResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_syncbyte_proto_init() }
func file_internal_proto_syncbyte_proto_init() {
	if File_internal_proto_syncbyte_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_syncbyte_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupRequest); i {
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
		file_internal_proto_syncbyte_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupResponse); i {
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
		file_internal_proto_syncbyte_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartInfo); i {
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
			RawDescriptor: file_internal_proto_syncbyte_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_syncbyte_proto_goTypes,
		DependencyIndexes: file_internal_proto_syncbyte_proto_depIdxs,
		MessageInfos:      file_internal_proto_syncbyte_proto_msgTypes,
	}.Build()
	File_internal_proto_syncbyte_proto = out.File
	file_internal_proto_syncbyte_proto_rawDesc = nil
	file_internal_proto_syncbyte_proto_goTypes = nil
	file_internal_proto_syncbyte_proto_depIdxs = nil
}
