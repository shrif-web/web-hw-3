// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: cache.proto

package go_cache_grpc

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

//service RequestCache{
////  rpc CacheNoteRpc (CacheRequest) returns(CacheResponse){}
//}
type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Type  string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Note) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CacheNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType int32  `protobuf:"varint,1,opt,name=requestType,proto3" json:"requestType,omitempty"`
	NoteId      string `protobuf:"bytes,2,opt,name=noteId,proto3" json:"noteId,omitempty"`
	NoteTitle   string `protobuf:"bytes,3,opt,name=noteTitle,proto3" json:"noteTitle,omitempty"`
	AuthorId    string `protobuf:"bytes,4,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Note        string `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CacheNoteRequest) Reset() {
	*x = CacheNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheNoteRequest) ProtoMessage() {}

func (x *CacheNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheNoteRequest.ProtoReflect.Descriptor instead.
func (*CacheNoteRequest) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{1}
}

func (x *CacheNoteRequest) GetRequestType() int32 {
	if x != nil {
		return x.RequestType
	}
	return 0
}

func (x *CacheNoteRequest) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *CacheNoteRequest) GetNoteTitle() string {
	if x != nil {
		return x.NoteTitle
	}
	return ""
}

func (x *CacheNoteRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CacheNoteRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CacheNoteRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CacheNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note      string  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Title     string  `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	NoteId    string  `protobuf:"bytes,2,opt,name=noteId,proto3" json:"noteId,omitempty"`
	Type      string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Notes     []*Note `protobuf:"bytes,7,rep,name=notes,proto3" json:"notes,omitempty"`
	Exist     bool    `protobuf:"varint,3,opt,name=exist,proto3" json:"exist,omitempty"`
	Access    bool    `protobuf:"varint,4,opt,name=access,proto3" json:"access,omitempty"`
	MissCache bool    `protobuf:"varint,5,opt,name=missCache,proto3" json:"missCache,omitempty"`
}

func (x *CacheNoteResponse) Reset() {
	*x = CacheNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheNoteResponse) ProtoMessage() {}

func (x *CacheNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheNoteResponse.ProtoReflect.Descriptor instead.
func (*CacheNoteResponse) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{2}
}

func (x *CacheNoteResponse) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CacheNoteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CacheNoteResponse) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *CacheNoteResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CacheNoteResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *CacheNoteResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *CacheNoteResponse) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

func (x *CacheNoteResponse) GetMissCache() bool {
	if x != nil {
		return x.MissCache
	}
	return false
}

type CacheLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType int32  `protobuf:"varint,1,opt,name=RequestType,proto3" json:"RequestType,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pass        string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *CacheLoginRequest) Reset() {
	*x = CacheLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheLoginRequest) ProtoMessage() {}

func (x *CacheLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheLoginRequest.ProtoReflect.Descriptor instead.
func (*CacheLoginRequest) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{3}
}

func (x *CacheLoginRequest) GetRequestType() int32 {
	if x != nil {
		return x.RequestType
	}
	return 0
}

func (x *CacheLoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CacheLoginRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CacheLoginRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type CacheLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string  `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserName  string  `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Name      string  `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
	Notes     []*Note `protobuf:"bytes,6,rep,name=notes,proto3" json:"notes,omitempty"`
	WrongPass bool    `protobuf:"varint,3,opt,name=WrongPass,proto3" json:"WrongPass,omitempty"`
	Exist     bool    `protobuf:"varint,4,opt,name=Exist,proto3" json:"Exist,omitempty"`
	MissCache bool    `protobuf:"varint,5,opt,name=MissCache,proto3" json:"MissCache,omitempty"`
}

func (x *CacheLoginResponse) Reset() {
	*x = CacheLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheLoginResponse) ProtoMessage() {}

func (x *CacheLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheLoginResponse.ProtoReflect.Descriptor instead.
func (*CacheLoginResponse) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{4}
}

func (x *CacheLoginResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CacheLoginResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CacheLoginResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CacheLoginResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *CacheLoginResponse) GetWrongPass() bool {
	if x != nil {
		return x.WrongPass
	}
	return false
}

func (x *CacheLoginResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *CacheLoginResponse) GetMissCache() bool {
	if x != nil {
		return x.MissCache
	}
	return false
}

var File_cache_proto protoreflect.FileDescriptor

var file_cache_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x54, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xae, 0x01,
	0x0a, 0x10, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x6f, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x6f, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xdb,
	0x01, 0x0a, 0x11, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x71, 0x0a, 0x11,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22,
	0xd4, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x50, 0x61, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x50, 0x61,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x73, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4d, 0x69, 0x73,
	0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x32, 0xae, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x50, 0x43, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x50, 0x43, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x68, 0x77, 0x33, 0x2f, 0x42,
	0x61, 0x63, 0x6b, 0x45, 0x6e, 0x64, 0x3b, 0x67, 0x6f, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cache_proto_rawDescOnce sync.Once
	file_cache_proto_rawDescData = file_cache_proto_rawDesc
)

func file_cache_proto_rawDescGZIP() []byte {
	file_cache_proto_rawDescOnce.Do(func() {
		file_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_proto_rawDescData)
	})
	return file_cache_proto_rawDescData
}

var file_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cache_proto_goTypes = []interface{}{
	(*Note)(nil),               // 0: usermgmt.Note
	(*CacheNoteRequest)(nil),   // 1: usermgmt.CacheNoteRequest
	(*CacheNoteResponse)(nil),  // 2: usermgmt.CacheNoteResponse
	(*CacheLoginRequest)(nil),  // 3: usermgmt.CacheLoginRequest
	(*CacheLoginResponse)(nil), // 4: usermgmt.CacheLoginResponse
}
var file_cache_proto_depIdxs = []int32{
	0, // 0: usermgmt.CacheNoteResponse.notes:type_name -> usermgmt.Note
	0, // 1: usermgmt.CacheLoginResponse.notes:type_name -> usermgmt.Note
	1, // 2: usermgmt.CacheManagement.CacheNoteRPC:input_type -> usermgmt.CacheNoteRequest
	3, // 3: usermgmt.CacheManagement.CacheLoginRPC:input_type -> usermgmt.CacheLoginRequest
	2, // 4: usermgmt.CacheManagement.CacheNoteRPC:output_type -> usermgmt.CacheNoteResponse
	4, // 5: usermgmt.CacheManagement.CacheLoginRPC:output_type -> usermgmt.CacheLoginResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cache_proto_init() }
func file_cache_proto_init() {
	if File_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheNoteRequest); i {
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
		file_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheNoteResponse); i {
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
		file_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheLoginRequest); i {
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
		file_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheLoginResponse); i {
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
			RawDescriptor: file_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_proto_goTypes,
		DependencyIndexes: file_cache_proto_depIdxs,
		MessageInfos:      file_cache_proto_msgTypes,
	}.Build()
	File_cache_proto = out.File
	file_cache_proto_rawDesc = nil
	file_cache_proto_goTypes = nil
	file_cache_proto_depIdxs = nil
}
