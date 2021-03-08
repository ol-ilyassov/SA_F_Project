// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: article_db/articlepb/article.proto

package articlepb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Created string `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Expires string `protobuf:"bytes,5,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Article) GetExpires() string {
	if x != nil {
		return x.Expires
	}
	return ""
}

type GetArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetArticlesRequest) Reset() {
	*x = GetArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesRequest) ProtoMessage() {}

func (x *GetArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesRequest.ProtoReflect.Descriptor instead.
func (*GetArticlesRequest) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{1}
}

type GetArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *GetArticlesResponse) Reset() {
	*x = GetArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesResponse) ProtoMessage() {}

func (x *GetArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesResponse.ProtoReflect.Descriptor instead.
func (*GetArticlesResponse) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticlesResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	Result  string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetArticleResponse) Reset() {
	*x = GetArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResponse) ProtoMessage() {}

func (x *GetArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResponse.ProtoReflect.Descriptor instead.
func (*GetArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{4}
}

func (x *GetArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *GetArticleResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type InsertArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *InsertArticleRequest) Reset() {
	*x = InsertArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertArticleRequest) ProtoMessage() {}

func (x *InsertArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertArticleRequest.ProtoReflect.Descriptor instead.
func (*InsertArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{5}
}

func (x *InsertArticleRequest) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type InsertArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InsertArticleResponse) Reset() {
	*x = InsertArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertArticleResponse) ProtoMessage() {}

func (x *InsertArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertArticleResponse.ProtoReflect.Descriptor instead.
func (*InsertArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{6}
}

func (x *InsertArticleResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InsertArticleResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteArticleRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteArticleResponse) Reset() {
	*x = DeleteArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleResponse) ProtoMessage() {}

func (x *DeleteArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleResponse.ProtoReflect.Descriptor instead.
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArticleResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type SearchArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *SearchArticlesRequest) Reset() {
	*x = SearchArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArticlesRequest) ProtoMessage() {}

func (x *SearchArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArticlesRequest.ProtoReflect.Descriptor instead.
func (*SearchArticlesRequest) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{9}
}

func (x *SearchArticlesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type SearchArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *SearchArticlesResponse) Reset() {
	*x = SearchArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_db_articlepb_article_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArticlesResponse) ProtoMessage() {}

func (x *SearchArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_db_articlepb_article_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArticlesResponse.ProtoReflect.Descriptor instead.
func (*SearchArticlesResponse) Descriptor() ([]byte, []int) {
	return file_article_db_articlepb_article_proto_rawDescGZIP(), []int{10}
}

func (x *SearchArticlesResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

var File_article_db_articlepb_article_proto protoreflect.FileDescriptor

var file_article_db_articlepb_article_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x64, 0x62, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x22,
	0x7d, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x67, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a,
	0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x22, 0x3f, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x26, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2d, 0x0a, 0x15, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x32, 0xb7, 0x03, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_article_db_articlepb_article_proto_rawDescOnce sync.Once
	file_article_db_articlepb_article_proto_rawDescData = file_article_db_articlepb_article_proto_rawDesc
)

func file_article_db_articlepb_article_proto_rawDescGZIP() []byte {
	file_article_db_articlepb_article_proto_rawDescOnce.Do(func() {
		file_article_db_articlepb_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_db_articlepb_article_proto_rawDescData)
	})
	return file_article_db_articlepb_article_proto_rawDescData
}

var file_article_db_articlepb_article_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_article_db_articlepb_article_proto_goTypes = []interface{}{
	(*Article)(nil),                // 0: articlepb.Article
	(*GetArticlesRequest)(nil),     // 1: articlepb.getArticlesRequest
	(*GetArticlesResponse)(nil),    // 2: articlepb.getArticlesResponse
	(*GetArticleRequest)(nil),      // 3: articlepb.getArticleRequest
	(*GetArticleResponse)(nil),     // 4: articlepb.getArticleResponse
	(*InsertArticleRequest)(nil),   // 5: articlepb.insertArticleRequest
	(*InsertArticleResponse)(nil),  // 6: articlepb.insertArticleResponse
	(*DeleteArticleRequest)(nil),   // 7: articlepb.deleteArticleRequest
	(*DeleteArticleResponse)(nil),  // 8: articlepb.deleteArticleResponse
	(*SearchArticlesRequest)(nil),  // 9: articlepb.searchArticlesRequest
	(*SearchArticlesResponse)(nil), // 10: articlepb.searchArticlesResponse
}
var file_article_db_articlepb_article_proto_depIdxs = []int32{
	0,  // 0: articlepb.getArticlesResponse.article:type_name -> articlepb.Article
	0,  // 1: articlepb.getArticleResponse.article:type_name -> articlepb.Article
	0,  // 2: articlepb.insertArticleRequest.article:type_name -> articlepb.Article
	0,  // 3: articlepb.searchArticlesResponse.article:type_name -> articlepb.Article
	1,  // 4: articlepb.ArticlesService.GetArticles:input_type -> articlepb.getArticlesRequest
	3,  // 5: articlepb.ArticlesService.GetArticle:input_type -> articlepb.getArticleRequest
	5,  // 6: articlepb.ArticlesService.InsertArticle:input_type -> articlepb.insertArticleRequest
	7,  // 7: articlepb.ArticlesService.DeleteArticle:input_type -> articlepb.deleteArticleRequest
	9,  // 8: articlepb.ArticlesService.SearchArticles:input_type -> articlepb.searchArticlesRequest
	2,  // 9: articlepb.ArticlesService.GetArticles:output_type -> articlepb.getArticlesResponse
	4,  // 10: articlepb.ArticlesService.GetArticle:output_type -> articlepb.getArticleResponse
	6,  // 11: articlepb.ArticlesService.InsertArticle:output_type -> articlepb.insertArticleResponse
	8,  // 12: articlepb.ArticlesService.DeleteArticle:output_type -> articlepb.deleteArticleResponse
	10, // 13: articlepb.ArticlesService.SearchArticles:output_type -> articlepb.searchArticlesResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_article_db_articlepb_article_proto_init() }
func file_article_db_articlepb_article_proto_init() {
	if File_article_db_articlepb_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_db_articlepb_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_article_db_articlepb_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesRequest); i {
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
		file_article_db_articlepb_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesResponse); i {
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
		file_article_db_articlepb_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_article_db_articlepb_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleResponse); i {
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
		file_article_db_articlepb_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertArticleRequest); i {
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
		file_article_db_articlepb_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertArticleResponse); i {
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
		file_article_db_articlepb_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_article_db_articlepb_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleResponse); i {
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
		file_article_db_articlepb_article_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArticlesRequest); i {
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
		file_article_db_articlepb_article_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArticlesResponse); i {
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
			RawDescriptor: file_article_db_articlepb_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_db_articlepb_article_proto_goTypes,
		DependencyIndexes: file_article_db_articlepb_article_proto_depIdxs,
		MessageInfos:      file_article_db_articlepb_article_proto_msgTypes,
	}.Build()
	File_article_db_articlepb_article_proto = out.File
	file_article_db_articlepb_article_proto_rawDesc = nil
	file_article_db_articlepb_article_proto_goTypes = nil
	file_article_db_articlepb_article_proto_depIdxs = nil
}
