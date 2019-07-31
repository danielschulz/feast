// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/DatasetService.proto

package core

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DatasetServiceTypes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetServiceTypes) Reset()         { *m = DatasetServiceTypes{} }
func (m *DatasetServiceTypes) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes) ProtoMessage()    {}
func (*DatasetServiceTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edc37a8b0d37b39, []int{0}
}

func (m *DatasetServiceTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes.Unmarshal(m, b)
}
func (m *DatasetServiceTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes.Marshal(b, m, deterministic)
}
func (m *DatasetServiceTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes.Merge(m, src)
}
func (m *DatasetServiceTypes) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes.Size(m)
}
func (m *DatasetServiceTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes proto.InternalMessageInfo

type DatasetServiceTypes_CreateDatasetRequest struct {
	// set of features for which its training data should be created
	FeatureSet *FeatureSet `protobuf:"bytes,1,opt,name=featureSet,proto3" json:"featureSet,omitempty"`
	// start date of the training data (inclusive)
	StartDate *timestamp.Timestamp `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	// end date of the training data (inclusive)
	EndDate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	// (optional) number of row that should be generated
	// (default: none)
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// (optional) prefix for dataset name
	NamePrefix string `protobuf:"bytes,5,opt,name=namePrefix,proto3" json:"namePrefix,omitempty"`
	// (optional) additional WHERE clause, all filter entry will be combined with logic AND
	Filters              map[string]string `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DatasetServiceTypes_CreateDatasetRequest) Reset() {
	*m = DatasetServiceTypes_CreateDatasetRequest{}
}
func (m *DatasetServiceTypes_CreateDatasetRequest) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes_CreateDatasetRequest) ProtoMessage()    {}
func (*DatasetServiceTypes_CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edc37a8b0d37b39, []int{0, 0}
}

func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Unmarshal(m, b)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Marshal(b, m, deterministic)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Merge(m, src)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Size(m)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest proto.InternalMessageInfo

func (m *DatasetServiceTypes_CreateDatasetRequest) GetFeatureSet() *FeatureSet {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetNamePrefix() string {
	if m != nil {
		return m.NamePrefix
	}
	return ""
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetFilters() map[string]string {
	if m != nil {
		return m.Filters
	}
	return nil
}

type DatasetServiceTypes_CreateDatasetResponse struct {
	// information of the created training dataset
	DatasetInfo          *DatasetInfo `protobuf:"bytes,1,opt,name=datasetInfo,proto3" json:"datasetInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DatasetServiceTypes_CreateDatasetResponse) Reset() {
	*m = DatasetServiceTypes_CreateDatasetResponse{}
}
func (m *DatasetServiceTypes_CreateDatasetResponse) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes_CreateDatasetResponse) ProtoMessage()    {}
func (*DatasetServiceTypes_CreateDatasetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edc37a8b0d37b39, []int{0, 1}
}

func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Unmarshal(m, b)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Marshal(b, m, deterministic)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Merge(m, src)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Size(m)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse proto.InternalMessageInfo

func (m *DatasetServiceTypes_CreateDatasetResponse) GetDatasetInfo() *DatasetInfo {
	if m != nil {
		return m.DatasetInfo
	}
	return nil
}

// Represent a collection of feature having same entity name
type FeatureSet struct {
	// entity related to this feature set
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// list of feature id in this feature set
	FeatureIds           []string `protobuf:"bytes,2,rep,name=featureIds,proto3" json:"featureIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSet) Reset()         { *m = FeatureSet{} }
func (m *FeatureSet) String() string { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()    {}
func (*FeatureSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edc37a8b0d37b39, []int{1}
}

func (m *FeatureSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSet.Unmarshal(m, b)
}
func (m *FeatureSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSet.Marshal(b, m, deterministic)
}
func (m *FeatureSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSet.Merge(m, src)
}
func (m *FeatureSet) XXX_Size() int {
	return xxx_messageInfo_FeatureSet.Size(m)
}
func (m *FeatureSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSet.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSet proto.InternalMessageInfo

func (m *FeatureSet) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *FeatureSet) GetFeatureIds() []string {
	if m != nil {
		return m.FeatureIds
	}
	return nil
}

// Representation of training dataset information
type DatasetInfo struct {
	// name of dataset
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL to table location of the training dataset
	TableUrl             string   `protobuf:"bytes,2,opt,name=tableUrl,proto3" json:"tableUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetInfo) Reset()         { *m = DatasetInfo{} }
func (m *DatasetInfo) String() string { return proto.CompactTextString(m) }
func (*DatasetInfo) ProtoMessage()    {}
func (*DatasetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edc37a8b0d37b39, []int{2}
}

func (m *DatasetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetInfo.Unmarshal(m, b)
}
func (m *DatasetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetInfo.Marshal(b, m, deterministic)
}
func (m *DatasetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetInfo.Merge(m, src)
}
func (m *DatasetInfo) XXX_Size() int {
	return xxx_messageInfo_DatasetInfo.Size(m)
}
func (m *DatasetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetInfo proto.InternalMessageInfo

func (m *DatasetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasetInfo) GetTableUrl() string {
	if m != nil {
		return m.TableUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*DatasetServiceTypes)(nil), "feast.core.DatasetServiceTypes")
	proto.RegisterType((*DatasetServiceTypes_CreateDatasetRequest)(nil), "feast.core.DatasetServiceTypes.CreateDatasetRequest")
	proto.RegisterMapType((map[string]string)(nil), "feast.core.DatasetServiceTypes.CreateDatasetRequest.FiltersEntry")
	proto.RegisterType((*DatasetServiceTypes_CreateDatasetResponse)(nil), "feast.core.DatasetServiceTypes.CreateDatasetResponse")
	proto.RegisterType((*FeatureSet)(nil), "feast.core.FeatureSet")
	proto.RegisterType((*DatasetInfo)(nil), "feast.core.DatasetInfo")
}

func init() { proto.RegisterFile("feast/core/DatasetService.proto", fileDescriptor_3edc37a8b0d37b39) }

var fileDescriptor_3edc37a8b0d37b39 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x3f, 0xc8, 0x04, 0x10, 0x5a, 0x0a, 0x58, 0x3e, 0x50, 0x2b, 0xa7, 0x9c, 0xd6,
	0x52, 0x68, 0x51, 0xa9, 0xc4, 0x81, 0x52, 0x2a, 0x55, 0x42, 0xa8, 0xda, 0x16, 0x09, 0xc1, 0x69,
	0x93, 0x8c, 0x8d, 0xa9, 0xed, 0x0d, 0xbb, 0xe3, 0x8a, 0x5c, 0xf8, 0x07, 0x48, 0xfc, 0x5d, 0x6e,
	0x68, 0xbd, 0x4e, 0xed, 0x44, 0x91, 0x10, 0xdc, 0x76, 0xe7, 0xbd, 0x99, 0x7d, 0xf3, 0x66, 0x07,
	0xf6, 0x13, 0x94, 0x86, 0xe2, 0xa9, 0xd2, 0x18, 0x9f, 0x4a, 0x92, 0x06, 0xe9, 0x12, 0xf5, 0x4d,
	0x36, 0x45, 0x3e, 0xd7, 0x8a, 0x14, 0x83, 0x9a, 0xc0, 0x2d, 0x21, 0xdc, 0x4f, 0x95, 0x4a, 0x73,
	0x8c, 0x6b, 0x64, 0x52, 0x25, 0x31, 0x65, 0x05, 0x1a, 0x92, 0xc5, 0xdc, 0x91, 0x87, 0xbf, 0x7d,
	0x78, 0xb4, 0x5a, 0xe5, 0x6a, 0x31, 0x47, 0x13, 0xfe, 0xf4, 0x61, 0xef, 0x8d, 0x46, 0x49, 0xd8,
	0xa0, 0x02, 0xbf, 0x55, 0x68, 0x88, 0xbd, 0x00, 0x5b, 0x9f, 0x2a, 0x8d, 0x97, 0x48, 0x81, 0x17,
	0x79, 0xa3, 0xc1, 0xf8, 0x09, 0x6f, 0x9f, 0xe4, 0x67, 0xb7, 0xa8, 0xe8, 0x30, 0xd9, 0x11, 0xf4,
	0x0d, 0x49, 0x4d, 0xa7, 0x92, 0x30, 0xe8, 0xd5, 0x69, 0x21, 0x77, 0xea, 0xf8, 0x52, 0x1d, 0xbf,
	0x5a, 0xaa, 0x13, 0x2d, 0x99, 0x1d, 0xc0, 0x2e, 0x96, 0xb3, 0x3a, 0xcf, 0xff, 0x6b, 0xde, 0x92,
	0xca, 0xf6, 0x60, 0x3b, 0xcf, 0x8a, 0x8c, 0x82, 0xad, 0xc8, 0x1b, 0xf9, 0xc2, 0x5d, 0xd8, 0x33,
	0x80, 0x52, 0x16, 0x78, 0xa1, 0x31, 0xc9, 0xbe, 0x07, 0xdb, 0x91, 0x37, 0xea, 0x8b, 0x4e, 0x84,
	0x7d, 0x86, 0xdd, 0x24, 0xcb, 0x09, 0xb5, 0x09, 0x76, 0x22, 0x7f, 0x34, 0x18, 0xbf, 0xee, 0xb6,
	0xb6, 0xc1, 0x28, 0xbe, 0xc9, 0x24, 0x7e, 0xe6, 0x6a, 0xbc, 0x2d, 0x49, 0x2f, 0xc4, 0xb2, 0x62,
	0x78, 0x0c, 0xf7, 0xba, 0x00, 0x7b, 0x08, 0xfe, 0x35, 0x2e, 0x6a, 0x0f, 0xfb, 0xc2, 0x1e, 0xad,
	0xe8, 0x1b, 0x99, 0x57, 0xce, 0xa0, 0xbe, 0x70, 0x97, 0xe3, 0xde, 0x91, 0x17, 0x0a, 0x78, 0xbc,
	0xf6, 0x92, 0x99, 0xab, 0xd2, 0x20, 0x7b, 0x09, 0x83, 0x99, 0x0b, 0x9d, 0x97, 0x89, 0x6a, 0x06,
	0xf2, 0x74, 0x83, 0x6a, 0x0b, 0x8b, 0x2e, 0x77, 0xf8, 0x0e, 0xa0, 0x1d, 0x96, 0xb5, 0x06, 0x4b,
	0xca, 0x68, 0xf1, 0x5e, 0x16, 0xd8, 0x88, 0xea, 0x44, 0x2c, 0xde, 0x8c, 0xf3, 0x7c, 0x66, 0x82,
	0x5e, 0xe4, 0x5b, 0xbc, 0x8d, 0x0c, 0x5f, 0xc1, 0xa0, 0xf3, 0x12, 0x63, 0xb0, 0x55, 0xb6, 0x85,
	0xea, 0x33, 0x0b, 0xe1, 0x2e, 0xc9, 0x49, 0x8e, 0x1f, 0x74, 0xde, 0x74, 0x78, 0x7b, 0x1f, 0xff,
	0xf2, 0xe0, 0xc1, 0xaa, 0xbf, 0xec, 0x07, 0xdc, 0x5f, 0xe9, 0x99, 0x1d, 0xfc, 0xcf, 0x30, 0xc2,
	0xc3, 0x7f, 0xcc, 0x72, 0xc6, 0x0e, 0xef, 0x9c, 0x7c, 0x84, 0xce, 0x2a, 0x9d, 0xac, 0xad, 0xc9,
	0x85, 0xfd, 0x7b, 0x9f, 0x0e, 0xd3, 0x8c, 0xbe, 0x54, 0x13, 0x3e, 0x55, 0x45, 0x9c, 0xaa, 0xaf,
	0x78, 0x1d, 0xbb, 0xed, 0xac, 0x7f, 0xa6, 0x89, 0x53, 0x2c, 0x51, 0x4b, 0xc2, 0x59, 0x9c, 0xaa,
	0xb8, 0xdd, 0xdb, 0xc9, 0x4e, 0x8d, 0x3f, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xa1, 0xf9,
	0x22, 0xcc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatasetServiceClient is the client API for DatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasetServiceClient interface {
	// Create training dataset for a feature set
	CreateDataset(ctx context.Context, in *DatasetServiceTypes_CreateDatasetRequest, opts ...grpc.CallOption) (*DatasetServiceTypes_CreateDatasetResponse, error)
}

type datasetServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatasetServiceClient(cc *grpc.ClientConn) DatasetServiceClient {
	return &datasetServiceClient{cc}
}

func (c *datasetServiceClient) CreateDataset(ctx context.Context, in *DatasetServiceTypes_CreateDatasetRequest, opts ...grpc.CallOption) (*DatasetServiceTypes_CreateDatasetResponse, error) {
	out := new(DatasetServiceTypes_CreateDatasetResponse)
	err := c.cc.Invoke(ctx, "/feast.core.DatasetService/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetServiceServer is the server API for DatasetService service.
type DatasetServiceServer interface {
	// Create training dataset for a feature set
	CreateDataset(context.Context, *DatasetServiceTypes_CreateDatasetRequest) (*DatasetServiceTypes_CreateDatasetResponse, error)
}

// UnimplementedDatasetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatasetServiceServer struct {
}

func (*UnimplementedDatasetServiceServer) CreateDataset(ctx context.Context, req *DatasetServiceTypes_CreateDatasetRequest) (*DatasetServiceTypes_CreateDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataset not implemented")
}

func RegisterDatasetServiceServer(s *grpc.Server, srv DatasetServiceServer) {
	s.RegisterService(&_DatasetService_serviceDesc, srv)
}

func _DatasetService_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasetServiceTypes_CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.DatasetService/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).CreateDataset(ctx, req.(*DatasetServiceTypes_CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.core.DatasetService",
	HandlerType: (*DatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDataset",
			Handler:    _DatasetService_CreateDataset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/core/DatasetService.proto",
}