// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: chat.proto

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_Message EventType = 0
	EventType_Info    EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "Message",
		1: "Info",
	}
	EventType_value = map[string]int32{
		"Message": 0,
		"Info":    1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

type ChatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ChatEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      EventType              `protobuf:"varint,1,opt,name=type,proto3,enum=chat.EventType" json:"type,omitempty"`
	UserID    int64                  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName  string                 `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	Text      string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Color     int32                  `protobuf:"varint,5,opt,name=color,proto3" json:"color,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChatEvent) Reset() {
	*x = ChatEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatEvent) ProtoMessage() {}

func (x *ChatEvent) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatEvent.ProtoReflect.Descriptor instead.
func (*ChatEvent) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_Message
}

func (x *ChatEvent) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ChatEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ChatEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ChatEvent) GetColor() int32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *ChatEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Userid []int64 `protobuf:"varint,2,rep,packed,name=userid,proto3" json:"userid,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetUserid() []int64 {
	if x != nil {
		return x.Userid
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID int64 `protobuf:"varint,1,opt,name=chatID,proto3" json:"chatID,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ConnectRequest) GetChatID() int64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID  int64      `protobuf:"varint,1,opt,name=chatID,proto3" json:"chatID,omitempty"`
	Message *ChatEvent `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *SendRequest) GetChatID() int64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

func (x *SendRequest) GetMessage() *ChatEvent {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID int64 `protobuf:"varint,1,opt,name=chatID,proto3" json:"chatID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRequest) GetChatID() int64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{11}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*ChatInfo `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{12}
}

func (x *ListResponse) GetChats() []*ChatInfo {
	if x != nil {
		return x.Chats
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x22, 0x50, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2a, 0x22,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x10, 0x01, 0x32, 0xce, 0x02, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chat_proto_goTypes = []any{
	(EventType)(0),                // 0: chat.EventType
	(*HealthRequest)(nil),         // 1: chat.HealthRequest
	(*HealthResponse)(nil),        // 2: chat.HealthResponse
	(*ChatInfo)(nil),              // 3: chat.ChatInfo
	(*ChatEvent)(nil),             // 4: chat.ChatEvent
	(*CreateRequest)(nil),         // 5: chat.CreateRequest
	(*CreateResponse)(nil),        // 6: chat.CreateResponse
	(*ConnectRequest)(nil),        // 7: chat.ConnectRequest
	(*SendRequest)(nil),           // 8: chat.SendRequest
	(*SendResponse)(nil),          // 9: chat.SendResponse
	(*DeleteRequest)(nil),         // 10: chat.DeleteRequest
	(*DeleteResponse)(nil),        // 11: chat.DeleteResponse
	(*ListRequest)(nil),           // 12: chat.ListRequest
	(*ListResponse)(nil),          // 13: chat.ListResponse
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_chat_proto_depIdxs = []int32{
	0,  // 0: chat.ChatEvent.type:type_name -> chat.EventType
	14, // 1: chat.ChatEvent.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 2: chat.SendRequest.message:type_name -> chat.ChatEvent
	3,  // 3: chat.ListResponse.chats:type_name -> chat.ChatInfo
	1,  // 4: chat.Chat.HealthCheck:input_type -> chat.HealthRequest
	5,  // 5: chat.Chat.Create:input_type -> chat.CreateRequest
	7,  // 6: chat.Chat.Conect:input_type -> chat.ConnectRequest
	8,  // 7: chat.Chat.SendMessage:input_type -> chat.SendRequest
	10, // 8: chat.Chat.Delete:input_type -> chat.DeleteRequest
	12, // 9: chat.Chat.List:input_type -> chat.ListRequest
	2,  // 10: chat.Chat.HealthCheck:output_type -> chat.HealthResponse
	6,  // 11: chat.Chat.Create:output_type -> chat.CreateResponse
	4,  // 12: chat.Chat.Conect:output_type -> chat.ChatEvent
	9,  // 13: chat.Chat.SendMessage:output_type -> chat.SendResponse
	11, // 14: chat.Chat.Delete:output_type -> chat.DeleteResponse
	13, // 15: chat.Chat.List:output_type -> chat.ListResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HealthRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HealthResponse); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChatInfo); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChatEvent); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectRequest); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SendRequest); i {
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
		file_chat_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SendResponse); i {
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
		file_chat_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_chat_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_chat_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ListRequest); i {
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
		file_chat_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
