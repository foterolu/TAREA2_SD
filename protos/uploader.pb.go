// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: uploader.proto

package cliente

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

// Enum value maps for UploadStatusCode.
var (
	UploadStatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	UploadStatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x UploadStatusCode) Enum() *UploadStatusCode {
	p := new(UploadStatusCode)
	*p = x
	return p
}

func (x UploadStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_uploader_proto_enumTypes[0].Descriptor()
}

func (UploadStatusCode) Type() protoreflect.EnumType {
	return &file_uploader_proto_enumTypes[0]
}

func (x UploadStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatusCode.Descriptor instead.
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{0}
}

type Adress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adress []string `protobuf:"bytes,1,rep,name=Adress,proto3" json:"Adress,omitempty"`
}

func (x *Adress) Reset() {
	*x = Adress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adress) ProtoMessage() {}

func (x *Adress) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adress.ProtoReflect.Descriptor instead.
func (*Adress) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{0}
}

func (x *Adress) GetAdress() []string {
	if x != nil {
		return x.Adress
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibro    string `protobuf:"bytes,1,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
	CantidadPartes string `protobuf:"bytes,2,opt,name=CantidadPartes,proto3" json:"CantidadPartes,omitempty"`
	Ubicaciones    string `protobuf:"bytes,3,opt,name=Ubicaciones,proto3" json:"Ubicaciones,omitempty"`
	Parte          string `protobuf:"bytes,4,opt,name=parte,proto3" json:"parte,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

func (x *Log) GetCantidadPartes() string {
	if x != nil {
		return x.CantidadPartes
	}
	return ""
}

func (x *Log) GetUbicaciones() string {
	if x != nil {
		return x.Ubicaciones
	}
	return ""
}

func (x *Log) GetParte() string {
	if x != nil {
		return x.Parte
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Libro   string `protobuf:"bytes,3,opt,name=Libro,proto3" json:"Libro,omitempty"`
	Partes  int32  `protobuf:"varint,4,opt,name=Partes,proto3" json:"Partes,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chunk) GetLibro() string {
	if x != nil {
		return x.Libro
	}
	return ""
}

func (x *Chunk) GetPartes() int32 {
	if x != nil {
		return x.Partes
	}
	return 0
}

type Prop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node string `protobuf:"bytes,1,opt,name=Node,proto3" json:"Node,omitempty"`
}

func (x *Prop) Reset() {
	*x = Prop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prop) ProtoMessage() {}

func (x *Prop) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prop.ProtoReflect.Descriptor instead.
func (*Prop) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{3}
}

func (x *Prop) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type Accept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (x *Accept) Reset() {
	*x = Accept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accept) ProtoMessage() {}

func (x *Accept) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accept.ProtoReflect.Descriptor instead.
func (*Accept) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{4}
}

func (x *Accept) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=cliente.UploadStatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_uploader_proto_rawDescGZIP(), []int{5}
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadStatus) GetCode() UploadStatusCode {
	if x != nil {
		return x.Code
	}
	return UploadStatusCode_Unknown
}

var File_uploader_proto protoreflect.FileDescriptor

var file_uploader_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x41, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62,
	0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43,
	0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x62,
	0x72, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x72,
	0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x1c, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x46, 0x6c, 0x61, 0x67, 0x22, 0x57, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x33, 0x0a,
	0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x02, 0x32, 0x8e, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2d, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x12,
	0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x0f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uploader_proto_rawDescOnce sync.Once
	file_uploader_proto_rawDescData = file_uploader_proto_rawDesc
)

func file_uploader_proto_rawDescGZIP() []byte {
	file_uploader_proto_rawDescOnce.Do(func() {
		file_uploader_proto_rawDescData = protoimpl.X.CompressGZIP(file_uploader_proto_rawDescData)
	})
	return file_uploader_proto_rawDescData
}

var file_uploader_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_uploader_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_uploader_proto_goTypes = []interface{}{
	(UploadStatusCode)(0), // 0: cliente.UploadStatusCode
	(*Adress)(nil),        // 1: cliente.Adress
	(*Log)(nil),           // 2: cliente.Log
	(*Chunk)(nil),         // 3: cliente.Chunk
	(*Prop)(nil),          // 4: cliente.Prop
	(*Accept)(nil),        // 5: cliente.Accept
	(*UploadStatus)(nil),  // 6: cliente.UploadStatus
}
var file_uploader_proto_depIdxs = []int32{
	0, // 0: cliente.UploadStatus.Code:type_name -> cliente.UploadStatusCode
	3, // 1: cliente.ChunksUpload.UploadChunk:input_type -> cliente.Chunk
	4, // 2: cliente.ChunksUpload.Propuesta:input_type -> cliente.Prop
	3, // 3: cliente.ChunksUpload.SendChunk:input_type -> cliente.Chunk
	2, // 4: cliente.ChunksUpload.SendLog:input_type -> cliente.Log
	4, // 5: cliente.ChunksUpload.RequestAdress:input_type -> cliente.Prop
	6, // 6: cliente.ChunksUpload.UploadChunk:output_type -> cliente.UploadStatus
	5, // 7: cliente.ChunksUpload.Propuesta:output_type -> cliente.Accept
	6, // 8: cliente.ChunksUpload.SendChunk:output_type -> cliente.UploadStatus
	5, // 9: cliente.ChunksUpload.SendLog:output_type -> cliente.Accept
	1, // 10: cliente.ChunksUpload.RequestAdress:output_type -> cliente.Adress
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uploader_proto_init() }
func file_uploader_proto_init() {
	if File_uploader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uploader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adress); i {
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
		file_uploader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_uploader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_uploader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prop); i {
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
		file_uploader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accept); i {
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
		file_uploader_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_uploader_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uploader_proto_goTypes,
		DependencyIndexes: file_uploader_proto_depIdxs,
		EnumInfos:         file_uploader_proto_enumTypes,
		MessageInfos:      file_uploader_proto_msgTypes,
	}.Build()
	File_uploader_proto = out.File
	file_uploader_proto_rawDesc = nil
	file_uploader_proto_goTypes = nil
	file_uploader_proto_depIdxs = nil
}
