// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/movie_service.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year      int32                  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Runtime   int32                  `protobuf:"varint,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genres    []string               `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *Movie) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Movie) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Movie) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMovieRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy string   `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Filter []string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListMoviesRequest) Reset() {
	*x = ListMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesRequest) ProtoMessage() {}

func (x *ListMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesRequest.ProtoReflect.Descriptor instead.
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListMoviesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListMoviesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListMoviesRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListMoviesRequest) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

type CreateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year    int32    `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Runtime int32    `protobuf:"varint,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genres  []string `protobuf:"bytes,4,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *CreateMovieRequest) Reset() {
	*x = CreateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieRequest) ProtoMessage() {}

func (x *CreateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMovieRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CreateMovieRequest) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *CreateMovieRequest) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

type UpdateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year    int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Runtime int32  `protobuf:"varint,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (x *UpdateMovieRequest) Reset() {
	*x = UpdateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieRequest) ProtoMessage() {}

func (x *UpdateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMovieRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMovieRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *UpdateMovieRequest) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

type DeleteMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMovieRequest) Reset() {
	*x = DeleteMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieRequest) ProtoMessage() {}

func (x *DeleteMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMovieRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *ListMoviesResponse) Reset() {
	*x = ListMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesResponse) ProtoMessage() {}

func (x *ListMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesResponse.ProtoReflect.Descriptor instead.
func (*ListMoviesResponse) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListMoviesResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

type DeleteMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMovieResponse) Reset() {
	*x = DeleteMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_movie_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieResponse) ProtoMessage() {}

func (x *DeleteMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieResponse.ProtoReflect.Descriptor instead.
func (*DeleteMovieResponse) Descriptor() ([]byte, []int) {
	return file_api_movie_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_movie_service_proto protoreflect.FileDescriptor

var file_api_movie_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01,
	0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x70, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8f, 0x03,
	0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01,
	0x2a, 0x22, 0x07, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x2a, 0x0c, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_movie_service_proto_rawDescOnce sync.Once
	file_api_movie_service_proto_rawDescData = file_api_movie_service_proto_rawDesc
)

func file_api_movie_service_proto_rawDescGZIP() []byte {
	file_api_movie_service_proto_rawDescOnce.Do(func() {
		file_api_movie_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_movie_service_proto_rawDescData)
	})
	return file_api_movie_service_proto_rawDescData
}

var file_api_movie_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_movie_service_proto_goTypes = []interface{}{
	(*Movie)(nil),                 // 0: api.Movie
	(*GetMovieRequest)(nil),       // 1: api.GetMovieRequest
	(*ListMoviesRequest)(nil),     // 2: api.ListMoviesRequest
	(*CreateMovieRequest)(nil),    // 3: api.CreateMovieRequest
	(*UpdateMovieRequest)(nil),    // 4: api.UpdateMovieRequest
	(*DeleteMovieRequest)(nil),    // 5: api.DeleteMovieRequest
	(*ListMoviesResponse)(nil),    // 6: api.ListMoviesResponse
	(*DeleteMovieResponse)(nil),   // 7: api.DeleteMovieResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_api_movie_service_proto_depIdxs = []int32{
	8, // 0: api.Movie.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: api.Movie.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: api.ListMoviesResponse.movies:type_name -> api.Movie
	1, // 3: api.MovieService.GetMovie:input_type -> api.GetMovieRequest
	2, // 4: api.MovieService.ListMovies:input_type -> api.ListMoviesRequest
	3, // 5: api.MovieService.CreateMovie:input_type -> api.CreateMovieRequest
	4, // 6: api.MovieService.UpdateMovie:input_type -> api.UpdateMovieRequest
	5, // 7: api.MovieService.DeleteMovie:input_type -> api.DeleteMovieRequest
	0, // 8: api.MovieService.GetMovie:output_type -> api.Movie
	6, // 9: api.MovieService.ListMovies:output_type -> api.ListMoviesResponse
	0, // 10: api.MovieService.CreateMovie:output_type -> api.Movie
	0, // 11: api.MovieService.UpdateMovie:output_type -> api.Movie
	7, // 12: api.MovieService.DeleteMovie:output_type -> api.DeleteMovieResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_movie_service_proto_init() }
func file_api_movie_service_proto_init() {
	if File_api_movie_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_movie_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_api_movie_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieRequest); i {
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
		file_api_movie_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesRequest); i {
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
		file_api_movie_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMovieRequest); i {
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
		file_api_movie_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovieRequest); i {
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
		file_api_movie_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieRequest); i {
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
		file_api_movie_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesResponse); i {
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
		file_api_movie_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieResponse); i {
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
			RawDescriptor: file_api_movie_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_movie_service_proto_goTypes,
		DependencyIndexes: file_api_movie_service_proto_depIdxs,
		MessageInfos:      file_api_movie_service_proto_msgTypes,
	}.Build()
	File_api_movie_service_proto = out.File
	file_api_movie_service_proto_rawDesc = nil
	file_api_movie_service_proto_goTypes = nil
	file_api_movie_service_proto_depIdxs = nil
}
