// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: message.proto

package pb

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

type MessageStatus int32

const (
	MessageStatus_MS_NORMAL MessageStatus = 0 // 正常的
	MessageStatus_MS_RECALL MessageStatus = 1 // 撤回
)

// Enum value maps for MessageStatus.
var (
	MessageStatus_name = map[int32]string{
		0: "MS_NORMAL",
		1: "MS_RECALL",
	}
	MessageStatus_value = map[string]int32{
		"MS_NORMAL": 0,
		"MS_RECALL": 1,
	}
)

func (x MessageStatus) Enum() *MessageStatus {
	p := new(MessageStatus)
	*p = x
	return p
}

func (x MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MessageStatus) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageStatus.Descriptor instead.
func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                           // 推送码
	Content  []byte        `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                      // 推送内容
	Seq      int64         `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`                             // 用户消息发送序列号
	SendTime int64         `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`   // 消息发送时间戳，精确到毫秒
	Status   MessageStatus `protobuf:"varint,5,opt,name=status,proto3,enum=pb.MessageStatus" json:"status,omitempty"` // 消息状态
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Message) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *Message) GetStatus() MessageStatus {
	if x != nil {
		return x.Status
	}
	return MessageStatus_MS_NORMAL
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2d, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x5f, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x5f, 0x52, 0x45,
	0x43, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_proto_goTypes = []interface{}{
	(MessageStatus)(0), // 0: pb.MessageStatus
	(*Message)(nil),    // 1: pb.Message
}
var file_message_proto_depIdxs = []int32{
	0, // 0: pb.Message.status:type_name -> pb.MessageStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
