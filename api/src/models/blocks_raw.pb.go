// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.11.3
// source: blocks_raw.proto

package models

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

type BlocksRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*BlockRaw `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlocksRaw) Reset() {
	*x = BlocksRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_raw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlocksRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlocksRaw) ProtoMessage() {}

func (x *BlocksRaw) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_raw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlocksRaw.ProtoReflect.Descriptor instead.
func (*BlocksRaw) Descriptor() ([]byte, []int) {
	return file_blocks_raw_proto_rawDescGZIP(), []int{0}
}

func (x *BlocksRaw) GetBlocks() []*BlockRaw {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_blocks_raw_proto protoreflect.FileDescriptor

var file_blocks_raw_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x09, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x61, 0x77, 0x12, 0x28, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x61, 0x77, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocks_raw_proto_rawDescOnce sync.Once
	file_blocks_raw_proto_rawDescData = file_blocks_raw_proto_rawDesc
)

func file_blocks_raw_proto_rawDescGZIP() []byte {
	file_blocks_raw_proto_rawDescOnce.Do(func() {
		file_blocks_raw_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocks_raw_proto_rawDescData)
	})
	return file_blocks_raw_proto_rawDescData
}

var file_blocks_raw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blocks_raw_proto_goTypes = []interface{}{
	(*BlocksRaw)(nil), // 0: models.BlocksRaw
	(*BlockRaw)(nil),  // 1: models.BlockRaw
}
var file_blocks_raw_proto_depIdxs = []int32{
	1, // 0: models.BlocksRaw.blocks:type_name -> models.BlockRaw
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blocks_raw_proto_init() }
func file_blocks_raw_proto_init() {
	if File_blocks_raw_proto != nil {
		return
	}
	file_block_raw_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blocks_raw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlocksRaw); i {
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
			RawDescriptor: file_blocks_raw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocks_raw_proto_goTypes,
		DependencyIndexes: file_blocks_raw_proto_depIdxs,
		MessageInfos:      file_blocks_raw_proto_msgTypes,
	}.Build()
	File_blocks_raw_proto = out.File
	file_blocks_raw_proto_rawDesc = nil
	file_blocks_raw_proto_goTypes = nil
	file_blocks_raw_proto_depIdxs = nil
}
