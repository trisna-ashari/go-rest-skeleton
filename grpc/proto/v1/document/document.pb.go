// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: grpc/proto/v1/document/document.proto

package document

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestDocumentsOrderBy int32

const (
	RequestDocuments_TITLE      RequestDocumentsOrderBy = 0
	RequestDocuments_CREATED_AT RequestDocumentsOrderBy = 1
)

// Enum value maps for RequestDocumentsOrderBy.
var (
	RequestDocumentsOrderBy_name = map[int32]string{
		0: "TITLE",
		1: "CREATED_AT",
	}
	RequestDocumentsOrderBy_value = map[string]int32{
		"TITLE":      0,
		"CREATED_AT": 1,
	}
)

func (x RequestDocumentsOrderBy) Enum() *RequestDocumentsOrderBy {
	p := new(RequestDocumentsOrderBy)
	*p = x
	return p
}

func (x RequestDocumentsOrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestDocumentsOrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_v1_document_document_proto_enumTypes[0].Descriptor()
}

func (RequestDocumentsOrderBy) Type() protoreflect.EnumType {
	return &file_grpc_proto_v1_document_document_proto_enumTypes[0]
}

func (x RequestDocumentsOrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestDocumentsOrderBy.Descriptor instead.
func (RequestDocumentsOrderBy) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{3, 0}
}

type RequestDocumentsOrderMode int32

const (
	RequestDocuments_ASC  RequestDocumentsOrderMode = 0
	RequestDocuments_DESC RequestDocumentsOrderMode = 1
)

// Enum value maps for RequestDocumentsOrderMode.
var (
	RequestDocumentsOrderMode_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	RequestDocumentsOrderMode_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x RequestDocumentsOrderMode) Enum() *RequestDocumentsOrderMode {
	p := new(RequestDocumentsOrderMode)
	*p = x
	return p
}

func (x RequestDocumentsOrderMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestDocumentsOrderMode) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_v1_document_document_proto_enumTypes[1].Descriptor()
}

func (RequestDocumentsOrderMode) Type() protoreflect.EnumType {
	return &file_grpc_proto_v1_document_document_proto_enumTypes[1]
}

func (x RequestDocumentsOrderMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestDocumentsOrderMode.Descriptor instead.
func (RequestDocumentsOrderMode) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{3, 1}
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Workflow  string `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_v1_document_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_v1_document_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{0}
}

func (x *Document) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Document) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

func (x *Document) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Document) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Documents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document []*Document `protobuf:"bytes,5,rep,name=document,proto3" json:"document,omitempty"`
}

func (x *Documents) Reset() {
	*x = Documents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_v1_document_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Documents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documents) ProtoMessage() {}

func (x *Documents) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_v1_document_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documents.ProtoReflect.Descriptor instead.
func (*Documents) Descriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{1}
}

func (x *Documents) GetDocument() []*Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type RequestDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RequestDocument) Reset() {
	*x = RequestDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_v1_document_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDocument) ProtoMessage() {}

func (x *RequestDocument) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_v1_document_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDocument.ProtoReflect.Descriptor instead.
func (*RequestDocument) Descriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{2}
}

func (x *RequestDocument) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RequestDocuments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Workflow string `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *RequestDocuments) Reset() {
	*x = RequestDocuments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_v1_document_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDocuments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDocuments) ProtoMessage() {}

func (x *RequestDocuments) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_v1_document_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDocuments.ProtoReflect.Descriptor instead.
func (*RequestDocuments) Descriptor() ([]byte, []int) {
	return file_grpc_proto_v1_document_document_proto_rawDescGZIP(), []int{3}
}

func (x *RequestDocuments) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RequestDocuments) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestDocuments) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

func (x *RequestDocuments) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_grpc_proto_v1_document_document_proto protoreflect.FileDescriptor

var file_grpc_proto_v1_document_document_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3b, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10,
	0x01, 0x22, 0x1f, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43,
	0x10, 0x01, 0x32, 0x90, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x13, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_v1_document_document_proto_rawDescOnce sync.Once
	file_grpc_proto_v1_document_document_proto_rawDescData = file_grpc_proto_v1_document_document_proto_rawDesc
)

func file_grpc_proto_v1_document_document_proto_rawDescGZIP() []byte {
	file_grpc_proto_v1_document_document_proto_rawDescOnce.Do(func() {
		file_grpc_proto_v1_document_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_v1_document_document_proto_rawDescData)
	})
	return file_grpc_proto_v1_document_document_proto_rawDescData
}

var file_grpc_proto_v1_document_document_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grpc_proto_v1_document_document_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_proto_v1_document_document_proto_goTypes = []interface{}{
	(RequestDocumentsOrderBy)(0),   // 0: document.RequestDocuments.order_by
	(RequestDocumentsOrderMode)(0), // 1: document.RequestDocuments.order_mode
	(*Document)(nil),               // 2: document.Document
	(*Documents)(nil),              // 3: document.Documents
	(*RequestDocument)(nil),        // 4: document.RequestDocument
	(*RequestDocuments)(nil),       // 5: document.RequestDocuments
}
var file_grpc_proto_v1_document_document_proto_depIdxs = []int32{
	2, // 0: document.Documents.document:type_name -> document.Document
	4, // 1: document.DocumentService.GetDocument:input_type -> document.RequestDocument
	5, // 2: document.DocumentService.GetDocuments:input_type -> document.RequestDocuments
	2, // 3: document.DocumentService.GetDocument:output_type -> document.Document
	3, // 4: document.DocumentService.GetDocuments:output_type -> document.Documents
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_v1_document_document_proto_init() }
func file_grpc_proto_v1_document_document_proto_init() {
	if File_grpc_proto_v1_document_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_v1_document_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_grpc_proto_v1_document_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Documents); i {
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
		file_grpc_proto_v1_document_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDocument); i {
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
		file_grpc_proto_v1_document_document_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDocuments); i {
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
			RawDescriptor: file_grpc_proto_v1_document_document_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_v1_document_document_proto_goTypes,
		DependencyIndexes: file_grpc_proto_v1_document_document_proto_depIdxs,
		EnumInfos:         file_grpc_proto_v1_document_document_proto_enumTypes,
		MessageInfos:      file_grpc_proto_v1_document_document_proto_msgTypes,
	}.Build()
	File_grpc_proto_v1_document_document_proto = out.File
	file_grpc_proto_v1_document_document_proto_rawDesc = nil
	file_grpc_proto_v1_document_document_proto_goTypes = nil
	file_grpc_proto_v1_document_document_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentServiceClient interface {
	GetDocument(ctx context.Context, in *RequestDocument, opts ...grpc.CallOption) (*Document, error)
	GetDocuments(ctx context.Context, in *RequestDocuments, opts ...grpc.CallOption) (*Documents, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) GetDocument(ctx context.Context, in *RequestDocument, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/document.DocumentService/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetDocuments(ctx context.Context, in *RequestDocuments, opts ...grpc.CallOption) (*Documents, error) {
	out := new(Documents)
	err := c.cc.Invoke(ctx, "/document.DocumentService/GetDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
type DocumentServiceServer interface {
	GetDocument(context.Context, *RequestDocument) (*Document, error)
	GetDocuments(context.Context, *RequestDocuments) (*Documents, error)
}

// UnimplementedDocumentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (*UnimplementedDocumentServiceServer) GetDocument(context.Context, *RequestDocument) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (*UnimplementedDocumentServiceServer) GetDocuments(context.Context, *RequestDocuments) (*Documents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocuments not implemented")
}

func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer) {
	s.RegisterService(&_DocumentService_serviceDesc, srv)
}

func _DocumentService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDocument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.DocumentService/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetDocument(ctx, req.(*RequestDocument))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDocuments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.DocumentService/GetDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetDocuments(ctx, req.(*RequestDocuments))
	}
	return interceptor(ctx, in, info, handler)
}

var _DocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "document.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _DocumentService_GetDocument_Handler,
		},
		{
			MethodName: "GetDocuments",
			Handler:    _DocumentService_GetDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/v1/document/document.proto",
}
