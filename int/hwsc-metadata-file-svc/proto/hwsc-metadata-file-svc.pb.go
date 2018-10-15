// Code generated by protoc-gen-go. DO NOT EDIT.
// source: int/hwsc-metadata-file-svc/proto/hwsc-metadata-file-svc.proto

package hwscMetadataFileSvcPb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetadataFile struct {
	LastName             string   `protobuf:"bytes,1,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	CallTypeName         string   `protobuf:"bytes,3,opt,name=call_type_name,json=callTypeName,proto3" json:"call_type_name,omitempty"`
	GroundType           string   `protobuf:"bytes,4,opt,name=ground_type,json=groundType,proto3" json:"ground_type,omitempty"`
	Region               string   `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Ocean                string   `protobuf:"bytes,6,opt,name=ocean,proto3" json:"ocean,omitempty"`
	SensorType           string   `protobuf:"bytes,7,opt,name=sensor_type,json=sensorType,proto3" json:"sensor_type,omitempty"`
	SensorName           string   `protobuf:"bytes,8,opt,name=sensor_name,json=sensorName,proto3" json:"sensor_name,omitempty"`
	SampleRate           int32    `protobuf:"varint,9,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	Latitude             float32  `protobuf:"fixed32,10,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,11,opt,name=longitude,proto3" json:"longitude,omitempty"`
	ImageUrl             []string `protobuf:"bytes,12,rep,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	AudioUrl             []string `protobuf:"bytes,13,rep,name=audio_url,json=audioUrl,proto3" json:"audio_url,omitempty"`
	VideoUrl             []string `protobuf:"bytes,14,rep,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	FileUrl              []string `protobuf:"bytes,15,rep,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	RecordTimestamp      int64    `protobuf:"varint,16,opt,name=record_timestamp,json=recordTimestamp,proto3" json:"record_timestamp,omitempty"`
	CreateTimestamp      int64    `protobuf:"varint,17,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"`
	UpdateTimestamp      int64    `protobuf:"varint,18,opt,name=update_timestamp,json=updateTimestamp,proto3" json:"update_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataFile) Reset()         { *m = MetadataFile{} }
func (m *MetadataFile) String() string { return proto.CompactTextString(m) }
func (*MetadataFile) ProtoMessage()    {}
func (*MetadataFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f0b51bc49740371, []int{0}
}

func (m *MetadataFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataFile.Unmarshal(m, b)
}
func (m *MetadataFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataFile.Marshal(b, m, deterministic)
}
func (m *MetadataFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataFile.Merge(m, src)
}
func (m *MetadataFile) XXX_Size() int {
	return xxx_messageInfo_MetadataFile.Size(m)
}
func (m *MetadataFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataFile.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataFile proto.InternalMessageInfo

func (m *MetadataFile) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *MetadataFile) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *MetadataFile) GetCallTypeName() string {
	if m != nil {
		return m.CallTypeName
	}
	return ""
}

func (m *MetadataFile) GetGroundType() string {
	if m != nil {
		return m.GroundType
	}
	return ""
}

func (m *MetadataFile) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *MetadataFile) GetOcean() string {
	if m != nil {
		return m.Ocean
	}
	return ""
}

func (m *MetadataFile) GetSensorType() string {
	if m != nil {
		return m.SensorType
	}
	return ""
}

func (m *MetadataFile) GetSensorName() string {
	if m != nil {
		return m.SensorName
	}
	return ""
}

func (m *MetadataFile) GetSampleRate() int32 {
	if m != nil {
		return m.SampleRate
	}
	return 0
}

func (m *MetadataFile) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *MetadataFile) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *MetadataFile) GetImageUrl() []string {
	if m != nil {
		return m.ImageUrl
	}
	return nil
}

func (m *MetadataFile) GetAudioUrl() []string {
	if m != nil {
		return m.AudioUrl
	}
	return nil
}

func (m *MetadataFile) GetVideoUrl() []string {
	if m != nil {
		return m.VideoUrl
	}
	return nil
}

func (m *MetadataFile) GetFileUrl() []string {
	if m != nil {
		return m.FileUrl
	}
	return nil
}

func (m *MetadataFile) GetRecordTimestamp() int64 {
	if m != nil {
		return m.RecordTimestamp
	}
	return 0
}

func (m *MetadataFile) GetCreateTimestamp() int64 {
	if m != nil {
		return m.CreateTimestamp
	}
	return 0
}

func (m *MetadataFile) GetUpdateTimestamp() int64 {
	if m != nil {
		return m.UpdateTimestamp
	}
	return 0
}

type MetadataFileRequest struct {
	Data                 *MetadataFile `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetadataFileRequest) Reset()         { *m = MetadataFileRequest{} }
func (m *MetadataFileRequest) String() string { return proto.CompactTextString(m) }
func (*MetadataFileRequest) ProtoMessage()    {}
func (*MetadataFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f0b51bc49740371, []int{1}
}

func (m *MetadataFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataFileRequest.Unmarshal(m, b)
}
func (m *MetadataFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataFileRequest.Marshal(b, m, deterministic)
}
func (m *MetadataFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataFileRequest.Merge(m, src)
}
func (m *MetadataFileRequest) XXX_Size() int {
	return xxx_messageInfo_MetadataFileRequest.Size(m)
}
func (m *MetadataFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataFileRequest proto.InternalMessageInfo

func (m *MetadataFileRequest) GetData() *MetadataFile {
	if m != nil {
		return m.Data
	}
	return nil
}

type MetadataFileResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataFileResponse) Reset()         { *m = MetadataFileResponse{} }
func (m *MetadataFileResponse) String() string { return proto.CompactTextString(m) }
func (*MetadataFileResponse) ProtoMessage()    {}
func (*MetadataFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f0b51bc49740371, []int{2}
}

func (m *MetadataFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataFileResponse.Unmarshal(m, b)
}
func (m *MetadataFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataFileResponse.Marshal(b, m, deterministic)
}
func (m *MetadataFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataFileResponse.Merge(m, src)
}
func (m *MetadataFileResponse) XXX_Size() int {
	return xxx_messageInfo_MetadataFileResponse.Size(m)
}
func (m *MetadataFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataFileResponse proto.InternalMessageInfo

func (m *MetadataFileResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MetadataFile)(nil), "hwscMetadataFileSvc.MetadataFile")
	proto.RegisterType((*MetadataFileRequest)(nil), "hwscMetadataFileSvc.MetadataFileRequest")
	proto.RegisterType((*MetadataFileResponse)(nil), "hwscMetadataFileSvc.MetadataFileResponse")
}

func init() {
	proto.RegisterFile("int/hwsc-metadata-file-svc/proto/hwsc-metadata-file-svc.proto", fileDescriptor_9f0b51bc49740371)
}

var fileDescriptor_9f0b51bc49740371 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0xc9, 0xfa, 0x33, 0xaf, 0x65, 0x1b, 0xde, 0x80, 0x30, 0x40, 0x84, 0x8a, 0x43, 0x76,
	0x68, 0x87, 0x86, 0x38, 0x72, 0x01, 0x89, 0x13, 0x20, 0x94, 0x8d, 0x0b, 0x97, 0xc8, 0x4b, 0xde,
	0x82, 0x85, 0x13, 0x07, 0xdb, 0x29, 0xda, 0x95, 0x7f, 0x90, 0x7f, 0x09, 0xf9, 0xb9, 0xd9, 0x1a,
	0x69, 0x08, 0x8e, 0xfe, 0xbe, 0xaf, 0xd6, 0x4b, 0xfa, 0x02, 0x6f, 0x44, 0x6d, 0x4f, 0xbe, 0xfd,
	0x34, 0xf9, 0xb2, 0x42, 0xcb, 0x0b, 0x6e, 0xf9, 0xf2, 0x52, 0x48, 0x5c, 0x9a, 0x75, 0x7e, 0xd2,
	0x68, 0x65, 0xd5, 0x5f, 0xe4, 0x8a, 0x24, 0x3b, 0x70, 0xf6, 0xe3, 0x46, 0xbe, 0x17, 0x12, 0xcf,
	0xd6, 0xf9, 0xe2, 0xf7, 0x10, 0xe6, 0xdb, 0x8c, 0x3d, 0x86, 0x50, 0x72, 0x63, 0xb3, 0x9a, 0x57,
	0x18, 0x05, 0x71, 0x90, 0x84, 0xe9, 0xd4, 0x81, 0x4f, 0xbc, 0x42, 0xf6, 0x14, 0xe0, 0x52, 0xe8,
	0xce, 0xee, 0x90, 0x0d, 0x89, 0x90, 0x7e, 0x01, 0xbb, 0x39, 0x97, 0x32, 0xb3, 0x57, 0x0d, 0xfa,
	0x64, 0x40, 0xc9, 0xdc, 0xd1, 0xf3, 0xab, 0x06, 0xa9, 0x7a, 0x06, 0xb3, 0x52, 0xab, 0xb6, 0x2e,
	0xa8, 0x8b, 0x86, 0x94, 0x80, 0x47, 0x2e, 0x62, 0x0f, 0x60, 0xac, 0xb1, 0x14, 0xaa, 0x8e, 0x46,
	0xe4, 0x36, 0x27, 0x76, 0x08, 0x23, 0x95, 0x23, 0xaf, 0xa3, 0x31, 0x61, 0x7f, 0x70, 0xd7, 0x19,
	0xac, 0x8d, 0xd2, 0xfe, 0xba, 0x89, 0xbf, 0xce, 0x23, 0xba, 0xee, 0x26, 0xa0, 0x91, 0xa6, 0xdb,
	0x41, 0x37, 0x90, 0xe1, 0x55, 0x23, 0x31, 0xd3, 0xdc, 0x62, 0x14, 0xc6, 0x41, 0x32, 0x4a, 0xc1,
	0xa3, 0x94, 0x5b, 0x64, 0x47, 0x30, 0x95, 0xdc, 0x0a, 0xdb, 0x16, 0x18, 0x41, 0x1c, 0x24, 0x3b,
	0xe9, 0xf5, 0x99, 0x3d, 0x81, 0x50, 0xaa, 0xba, 0xf4, 0x72, 0x46, 0xf2, 0x06, 0xb8, 0xb7, 0x29,
	0x2a, 0x5e, 0x62, 0xd6, 0x6a, 0x19, 0xcd, 0xe3, 0x81, 0x7b, 0x9b, 0x04, 0xbe, 0x68, 0xe9, 0x24,
	0x6f, 0x0b, 0xa1, 0x48, 0xde, 0xf5, 0x92, 0xc0, 0x46, 0xae, 0x45, 0x81, 0x5e, 0xee, 0x7a, 0x49,
	0xc0, 0xc9, 0x47, 0x30, 0x75, 0x7f, 0x2e, 0xb9, 0x3d, 0x72, 0x13, 0x77, 0x76, 0xea, 0x18, 0xf6,
	0x35, 0xe6, 0x4a, 0x17, 0x99, 0x15, 0x15, 0x1a, 0xcb, 0xab, 0x26, 0xda, 0x8f, 0x83, 0x64, 0x90,
	0xee, 0x79, 0x7e, 0xde, 0x61, 0x97, 0xe6, 0x1a, 0xb9, 0xc5, 0xad, 0xf4, 0x9e, 0x4f, 0x3d, 0xef,
	0xa5, 0x6d, 0x53, 0xf4, 0x53, 0xe6, 0x53, 0xcf, 0xaf, 0xd3, 0xc5, 0x07, 0x38, 0xd8, 0x5e, 0xa8,
	0x14, 0x7f, 0xb4, 0x68, 0x2c, 0x7b, 0x0d, 0x43, 0x87, 0x68, 0xa5, 0x66, 0xa7, 0xcf, 0x57, 0xb7,
	0x2c, 0xe3, 0xaa, 0xf7, 0x3b, 0xca, 0x17, 0x2f, 0xe1, 0xb0, 0x7f, 0x9b, 0x69, 0x54, 0x6d, 0x90,
	0x45, 0x30, 0xa9, 0xd0, 0x18, 0x5e, 0x76, 0x4b, 0xda, 0x1d, 0x4f, 0x7f, 0x05, 0xfd, 0x01, 0xce,
	0x50, 0xaf, 0x45, 0x8e, 0xec, 0x3b, 0xb0, 0x77, 0xf4, 0x54, 0xbd, 0x75, 0x4f, 0xfe, 0x3d, 0x88,
	0x7f, 0x80, 0xa3, 0xe3, 0xff, 0x28, 0xfd, 0x70, 0x8b, 0x3b, 0x6f, 0x1f, 0x7e, 0xbd, 0x7f, 0x4b,
	0xfd, 0xf9, 0xe2, 0x62, 0x4c, 0xdf, 0xe2, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x40,
	0x4d, 0x2e, 0xcc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataFileServiceClient is the client API for MetadataFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataFileServiceClient interface {
	CreateMetadataFile(ctx context.Context, in *MetadataFileRequest, opts ...grpc.CallOption) (*MetadataFileResponse, error)
}

type metadataFileServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetadataFileServiceClient(cc *grpc.ClientConn) MetadataFileServiceClient {
	return &metadataFileServiceClient{cc}
}

func (c *metadataFileServiceClient) CreateMetadataFile(ctx context.Context, in *MetadataFileRequest, opts ...grpc.CallOption) (*MetadataFileResponse, error) {
	out := new(MetadataFileResponse)
	err := c.cc.Invoke(ctx, "/hwscMetadataFileSvc.MetadataFileService/CreateMetadataFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataFileServiceServer is the server API for MetadataFileService service.
type MetadataFileServiceServer interface {
	CreateMetadataFile(context.Context, *MetadataFileRequest) (*MetadataFileResponse, error)
}

func RegisterMetadataFileServiceServer(s *grpc.Server, srv MetadataFileServiceServer) {
	s.RegisterService(&_MetadataFileService_serviceDesc, srv)
}

func _MetadataFileService_CreateMetadataFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataFileServiceServer).CreateMetadataFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hwscMetadataFileSvc.MetadataFileService/CreateMetadataFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataFileServiceServer).CreateMetadataFile(ctx, req.(*MetadataFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hwscMetadataFileSvc.MetadataFileService",
	HandlerType: (*MetadataFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMetadataFile",
			Handler:    _MetadataFileService_CreateMetadataFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "int/hwsc-metadata-file-svc/proto/hwsc-metadata-file-svc.proto",
}
