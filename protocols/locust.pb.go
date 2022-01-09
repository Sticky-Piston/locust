// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: protocols/locust.proto

package protocols

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

// designed to be shared between all app protocols
type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// shared between all requests
	ClientVersion string `protobuf:"bytes,1,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"` // client version
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`        // unix time
	Id            string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`                       // allows requesters to use request data when processing a response
	Gossip        bool   `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`              // true to have receiver peer gossip the message to neighbors
	NodeId        string `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`               // id of node that created the message (not the peer that may have sent it). =base58(multihash(nodePubKey))
	NodePubKey    []byte `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`       // Authoring node Secp256k1 public key (32bytes) - protobufs serielized
	Sign          []byte `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`                   // signature of message data + method specific data by message authoring node.
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_locust_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_locust_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_protocols_locust_proto_rawDescGZIP(), []int{0}
}

func (x *MessageData) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *MessageData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MessageData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageData) GetGossip() bool {
	if x != nil {
		return x.Gossip
	}
	return false
}

func (x *MessageData) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *MessageData) GetNodePubKey() []byte {
	if x != nil {
		return x.NodePubKey
	}
	return nil
}

func (x *MessageData) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	Message     string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_locust_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_locust_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_protocols_locust_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileRequest) GetMessageData() *MessageData {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *ProfileRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	Title       string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Summary     string       `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *ProfileResponse) Reset() {
	*x = ProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_locust_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResponse) ProtoMessage() {}

func (x *ProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_locust_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResponse.ProtoReflect.Descriptor instead.
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return file_protocols_locust_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileResponse) GetMessageData() *MessageData {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *ProfileResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProfileResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

var File_protocols_locust_proto protoreflect.FileDescriptor

var file_protocols_locust_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x75,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x6f, 0x64,
	0x65, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x64, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x7b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_locust_proto_rawDescOnce sync.Once
	file_protocols_locust_proto_rawDescData = file_protocols_locust_proto_rawDesc
)

func file_protocols_locust_proto_rawDescGZIP() []byte {
	file_protocols_locust_proto_rawDescOnce.Do(func() {
		file_protocols_locust_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_locust_proto_rawDescData)
	})
	return file_protocols_locust_proto_rawDescData
}

var file_protocols_locust_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocols_locust_proto_goTypes = []interface{}{
	(*MessageData)(nil),     // 0: protocols.MessageData
	(*ProfileRequest)(nil),  // 1: protocols.ProfileRequest
	(*ProfileResponse)(nil), // 2: protocols.ProfileResponse
}
var file_protocols_locust_proto_depIdxs = []int32{
	0, // 0: protocols.ProfileRequest.messageData:type_name -> protocols.MessageData
	0, // 1: protocols.ProfileResponse.messageData:type_name -> protocols.MessageData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocols_locust_proto_init() }
func file_protocols_locust_proto_init() {
	if File_protocols_locust_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_locust_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
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
		file_protocols_locust_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_protocols_locust_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileResponse); i {
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
			RawDescriptor: file_protocols_locust_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocols_locust_proto_goTypes,
		DependencyIndexes: file_protocols_locust_proto_depIdxs,
		MessageInfos:      file_protocols_locust_proto_msgTypes,
	}.Build()
	File_protocols_locust_proto = out.File
	file_protocols_locust_proto_rawDesc = nil
	file_protocols_locust_proto_goTypes = nil
	file_protocols_locust_proto_depIdxs = nil
}