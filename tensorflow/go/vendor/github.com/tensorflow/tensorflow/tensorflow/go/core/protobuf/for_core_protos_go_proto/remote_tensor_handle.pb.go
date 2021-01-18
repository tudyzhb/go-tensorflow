// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/remote_tensor_handle.proto

package for_core_protos_go_proto // import "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
import types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResourceDtypeAndShape struct {
	Dtype                types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ResourceDtypeAndShape) Reset()         { *m = ResourceDtypeAndShape{} }
func (m *ResourceDtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceDtypeAndShape) ProtoMessage()    {}
func (*ResourceDtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_tensor_handle_3b027cb90946ee66, []int{0}
}
func (m *ResourceDtypeAndShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDtypeAndShape.Unmarshal(m, b)
}
func (m *ResourceDtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDtypeAndShape.Marshal(b, m, deterministic)
}
func (dst *ResourceDtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDtypeAndShape.Merge(dst, src)
}
func (m *ResourceDtypeAndShape) XXX_Size() int {
	return xxx_messageInfo_ResourceDtypeAndShape.Size(m)
}
func (m *ResourceDtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDtypeAndShape proto.InternalMessageInfo

func (m *ResourceDtypeAndShape) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *ResourceDtypeAndShape) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

type RemoteTensorHandle struct {
	// The ID of the operation that produced this tensor.
	OpId int64 `protobuf:"varint,1,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// The index into the outputs of the operation that produced this tensor.
	OutputNum int32 `protobuf:"varint,2,opt,name=output_num,json=outputNum,proto3" json:"output_num,omitempty"`
	// Device where the tensor is located. Cannot be empty.
	// For multi-device functions, it's the default device passed to placer.
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// Device of the operation producing this tensor. Can be empty if the
	// operation producing this tensor is a multi-device function.
	OpDevice string `protobuf:"bytes,4,opt,name=op_device,json=opDevice,proto3" json:"op_device,omitempty"`
	// Tensor type.
	Dtype types_go_proto.DataType `protobuf:"varint,5,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Optional data types and shapes of a remote resource variable.
	ResourceDtypesAndShapes []*ResourceDtypeAndShape `protobuf:"bytes,6,rep,name=resource_dtypes_and_shapes,json=resourceDtypesAndShapes,proto3" json:"resource_dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *RemoteTensorHandle) Reset()         { *m = RemoteTensorHandle{} }
func (m *RemoteTensorHandle) String() string { return proto.CompactTextString(m) }
func (*RemoteTensorHandle) ProtoMessage()    {}
func (*RemoteTensorHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_tensor_handle_3b027cb90946ee66, []int{1}
}
func (m *RemoteTensorHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteTensorHandle.Unmarshal(m, b)
}
func (m *RemoteTensorHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteTensorHandle.Marshal(b, m, deterministic)
}
func (dst *RemoteTensorHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteTensorHandle.Merge(dst, src)
}
func (m *RemoteTensorHandle) XXX_Size() int {
	return xxx_messageInfo_RemoteTensorHandle.Size(m)
}
func (m *RemoteTensorHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteTensorHandle.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteTensorHandle proto.InternalMessageInfo

func (m *RemoteTensorHandle) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *RemoteTensorHandle) GetOutputNum() int32 {
	if m != nil {
		return m.OutputNum
	}
	return 0
}

func (m *RemoteTensorHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RemoteTensorHandle) GetOpDevice() string {
	if m != nil {
		return m.OpDevice
	}
	return ""
}

func (m *RemoteTensorHandle) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *RemoteTensorHandle) GetResourceDtypesAndShapes() []*ResourceDtypeAndShape {
	if m != nil {
		return m.ResourceDtypesAndShapes
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDtypeAndShape)(nil), "tensorflow.eager.ResourceDtypeAndShape")
	proto.RegisterType((*RemoteTensorHandle)(nil), "tensorflow.eager.RemoteTensorHandle")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/remote_tensor_handle.proto", fileDescriptor_remote_tensor_handle_3b027cb90946ee66)
}

var fileDescriptor_remote_tensor_handle_3b027cb90946ee66 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x45, 0x49, 0x1c, 0x36, 0x0a, 0x2c, 0x8b, 0xf6, 0xcb, 0x64, 0xb7, 0x60, 0x02, 0xa5, 0xa6,
	0x14, 0x1b, 0x9c, 0x5f, 0xd0, 0x90, 0x43, 0x7b, 0x29, 0x41, 0x4d, 0x2f, 0xbd, 0x08, 0xc7, 0x92,
	0x9d, 0xd0, 0xd8, 0x23, 0x24, 0xbb, 0x21, 0xbf, 0xa2, 0xfd, 0xb9, 0x3d, 0x16, 0x4b, 0x6e, 0x6a,
	0xd2, 0xd0, 0xde, 0xc6, 0xef, 0xbd, 0x79, 0xe3, 0x79, 0x23, 0x3c, 0x29, 0x45, 0xa1, 0x41, 0xa5,
	0x1b, 0xd8, 0x86, 0x09, 0x28, 0x11, 0x4a, 0x05, 0x25, 0x2c, 0xab, 0x34, 0x54, 0x22, 0x87, 0x52,
	0x30, 0xcb, 0xb3, 0x55, 0x5c, 0xf0, 0x8d, 0x08, 0x0c, 0x4b, 0x7e, 0xbc, 0x37, 0x05, 0x22, 0xce,
	0x84, 0x1a, 0x5d, 0x1c, 0xda, 0xa4, 0x2a, 0xce, 0xc5, 0x16, 0xd4, 0x43, 0xd8, 0x18, 0xe8, 0x55,
	0x2c, 0x9b, 0xfe, 0xd1, 0xe9, 0x27, 0xea, 0x9d, 0x14, 0xda, 0xca, 0xc6, 0x5b, 0xfc, 0x9b, 0x0a,
	0x0d, 0x95, 0x4a, 0xc4, 0xac, 0xc6, 0x2f, 0x0b, 0x7e, 0x5b, 0xbb, 0x90, 0x73, 0xec, 0xf0, 0x1a,
	0x70, 0x91, 0x87, 0xfc, 0xef, 0xd1, 0xaf, 0xa0, 0xf5, 0x3f, 0xb3, 0xb8, 0x8c, 0x17, 0x3b, 0x29,
	0xa8, 0x95, 0x90, 0x08, 0x3b, 0x66, 0xb4, 0xdb, 0xf1, 0x90, 0x3f, 0x8c, 0xfe, 0xb7, 0xb5, 0x0b,
	0x53, 0x1a, 0xcf, 0x79, 0x3d, 0x91, 0x5a, 0xe9, 0xf8, 0xa9, 0x83, 0x09, 0x35, 0xeb, 0x5b, 0xc5,
	0x95, 0x59, 0x9e, 0xfc, 0xc4, 0x0e, 0x48, 0xb6, 0xe6, 0x66, 0x6c, 0x97, 0xf6, 0x40, 0x5e, 0x73,
	0x72, 0x82, 0x31, 0x54, 0xa5, 0xac, 0x4a, 0x56, 0x54, 0xb9, 0x19, 0xe2, 0xd0, 0x81, 0x45, 0x6e,
	0xaa, 0x9c, 0xfc, 0xc1, 0x7d, 0x2e, 0x1e, 0xd7, 0x89, 0x70, 0xbb, 0x1e, 0xf2, 0x07, 0xb4, 0xf9,
	0x22, 0xff, 0xf0, 0x00, 0x24, 0x6b, 0xa8, 0x9e, 0xa1, 0xbe, 0x81, 0x9c, 0x59, 0x72, 0xbf, 0x9f,
	0xf3, 0xf5, 0x7e, 0x1c, 0x8f, 0x54, 0x13, 0x12, 0x33, 0x88, 0x66, 0x71, 0xc1, 0x6d, 0xdc, 0xda,
	0xed, 0x7b, 0x5d, 0x7f, 0x18, 0x9d, 0x05, 0x87, 0x07, 0x0b, 0x8e, 0x06, 0x4b, 0xff, 0xaa, 0x36,
	0xac, 0xdf, 0x70, 0x3d, 0x7d, 0x46, 0xd8, 0x05, 0x95, 0xb5, 0x7d, 0xf6, 0x37, 0x9b, 0xba, 0x1f,
	0xb3, 0x32, 0x71, 0xea, 0x39, 0xba, 0xbf, 0xcb, 0xd6, 0xe5, 0xaa, 0x5a, 0x06, 0x09, 0xe4, 0x61,
	0xeb, 0xea, 0xc7, 0xcb, 0x0c, 0x0e, 0xde, 0x60, 0x0a, 0x8a, 0xd5, 0x08, 0x33, 0x88, 0x66, 0x19,
	0xd8, 0xea, 0x05, 0xa1, 0x65, 0xdf, 0x54, 0x93, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x2b,
	0x6f, 0x0c, 0xc2, 0x02, 0x00, 0x00,
}
