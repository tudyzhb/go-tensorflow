// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/control_flow.proto

package core_protos_go_proto // import "github.com/tensorflow/tensorflow/tensorflow/go/core/core_protos_go_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Protocol buffer representing the values in ControlFlowContext.
type ValuesDef struct {
	// Value names that have been seen in this context.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Value names referenced by but external to this context.
	ExternalValues       map[string]string `protobuf:"bytes,2,rep,name=external_values,json=externalValues,proto3" json:"external_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ValuesDef) Reset()         { *m = ValuesDef{} }
func (m *ValuesDef) String() string { return proto.CompactTextString(m) }
func (*ValuesDef) ProtoMessage()    {}
func (*ValuesDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_flow_ed32bd4a24b0fb72, []int{0}
}
func (m *ValuesDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValuesDef.Unmarshal(m, b)
}
func (m *ValuesDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValuesDef.Marshal(b, m, deterministic)
}
func (dst *ValuesDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValuesDef.Merge(dst, src)
}
func (m *ValuesDef) XXX_Size() int {
	return xxx_messageInfo_ValuesDef.Size(m)
}
func (m *ValuesDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ValuesDef.DiscardUnknown(m)
}

var xxx_messageInfo_ValuesDef proto.InternalMessageInfo

func (m *ValuesDef) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ValuesDef) GetExternalValues() map[string]string {
	if m != nil {
		return m.ExternalValues
	}
	return nil
}

// Container for any kind of control flow context. Any other control flow
// contexts that are added below should also be added here.
type ControlFlowContextDef struct {
	// Types that are valid to be assigned to Ctxt:
	//	*ControlFlowContextDef_CondCtxt
	//	*ControlFlowContextDef_WhileCtxt
	Ctxt                 isControlFlowContextDef_Ctxt `protobuf_oneof:"ctxt"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ControlFlowContextDef) Reset()         { *m = ControlFlowContextDef{} }
func (m *ControlFlowContextDef) String() string { return proto.CompactTextString(m) }
func (*ControlFlowContextDef) ProtoMessage()    {}
func (*ControlFlowContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_flow_ed32bd4a24b0fb72, []int{1}
}
func (m *ControlFlowContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlFlowContextDef.Unmarshal(m, b)
}
func (m *ControlFlowContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlFlowContextDef.Marshal(b, m, deterministic)
}
func (dst *ControlFlowContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlFlowContextDef.Merge(dst, src)
}
func (m *ControlFlowContextDef) XXX_Size() int {
	return xxx_messageInfo_ControlFlowContextDef.Size(m)
}
func (m *ControlFlowContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlFlowContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_ControlFlowContextDef proto.InternalMessageInfo

type isControlFlowContextDef_Ctxt interface {
	isControlFlowContextDef_Ctxt()
}

type ControlFlowContextDef_CondCtxt struct {
	CondCtxt *CondContextDef `protobuf:"bytes,1,opt,name=cond_ctxt,json=condCtxt,proto3,oneof"`
}

type ControlFlowContextDef_WhileCtxt struct {
	WhileCtxt *WhileContextDef `protobuf:"bytes,2,opt,name=while_ctxt,json=whileCtxt,proto3,oneof"`
}

func (*ControlFlowContextDef_CondCtxt) isControlFlowContextDef_Ctxt() {}

func (*ControlFlowContextDef_WhileCtxt) isControlFlowContextDef_Ctxt() {}

func (m *ControlFlowContextDef) GetCtxt() isControlFlowContextDef_Ctxt {
	if m != nil {
		return m.Ctxt
	}
	return nil
}

func (m *ControlFlowContextDef) GetCondCtxt() *CondContextDef {
	if x, ok := m.GetCtxt().(*ControlFlowContextDef_CondCtxt); ok {
		return x.CondCtxt
	}
	return nil
}

func (m *ControlFlowContextDef) GetWhileCtxt() *WhileContextDef {
	if x, ok := m.GetCtxt().(*ControlFlowContextDef_WhileCtxt); ok {
		return x.WhileCtxt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ControlFlowContextDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ControlFlowContextDef_OneofMarshaler, _ControlFlowContextDef_OneofUnmarshaler, _ControlFlowContextDef_OneofSizer, []interface{}{
		(*ControlFlowContextDef_CondCtxt)(nil),
		(*ControlFlowContextDef_WhileCtxt)(nil),
	}
}

func _ControlFlowContextDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ControlFlowContextDef)
	// ctxt
	switch x := m.Ctxt.(type) {
	case *ControlFlowContextDef_CondCtxt:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CondCtxt); err != nil {
			return err
		}
	case *ControlFlowContextDef_WhileCtxt:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WhileCtxt); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ControlFlowContextDef.Ctxt has unexpected type %T", x)
	}
	return nil
}

func _ControlFlowContextDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ControlFlowContextDef)
	switch tag {
	case 1: // ctxt.cond_ctxt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CondContextDef)
		err := b.DecodeMessage(msg)
		m.Ctxt = &ControlFlowContextDef_CondCtxt{msg}
		return true, err
	case 2: // ctxt.while_ctxt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WhileContextDef)
		err := b.DecodeMessage(msg)
		m.Ctxt = &ControlFlowContextDef_WhileCtxt{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ControlFlowContextDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ControlFlowContextDef)
	// ctxt
	switch x := m.Ctxt.(type) {
	case *ControlFlowContextDef_CondCtxt:
		s := proto.Size(x.CondCtxt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ControlFlowContextDef_WhileCtxt:
		s := proto.Size(x.WhileCtxt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Protocol buffer representing a CondContext object.
type CondContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// Name of the pred tensor.
	PredName string `protobuf:"bytes,2,opt,name=pred_name,json=predName,proto3" json:"pred_name,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,3,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Branch prediction. 0 or 1.
	Branch int32 `protobuf:"varint,4,opt,name=branch,proto3" json:"branch,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,5,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Contexts contained inside this context (e.g. nested conds).
	NestedContexts       []*ControlFlowContextDef `protobuf:"bytes,6,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CondContextDef) Reset()         { *m = CondContextDef{} }
func (m *CondContextDef) String() string { return proto.CompactTextString(m) }
func (*CondContextDef) ProtoMessage()    {}
func (*CondContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_flow_ed32bd4a24b0fb72, []int{2}
}
func (m *CondContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CondContextDef.Unmarshal(m, b)
}
func (m *CondContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CondContextDef.Marshal(b, m, deterministic)
}
func (dst *CondContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CondContextDef.Merge(dst, src)
}
func (m *CondContextDef) XXX_Size() int {
	return xxx_messageInfo_CondContextDef.Size(m)
}
func (m *CondContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CondContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_CondContextDef proto.InternalMessageInfo

func (m *CondContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *CondContextDef) GetPredName() string {
	if m != nil {
		return m.PredName
	}
	return ""
}

func (m *CondContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *CondContextDef) GetBranch() int32 {
	if m != nil {
		return m.Branch
	}
	return 0
}

func (m *CondContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

func (m *CondContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if m != nil {
		return m.NestedContexts
	}
	return nil
}

// Protocol buffer representing a WhileContext object.
type WhileContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// The number of iterations allowed to run in parallel.
	ParallelIterations int32 `protobuf:"varint,2,opt,name=parallel_iterations,json=parallelIterations,proto3" json:"parallel_iterations,omitempty"`
	// Whether backprop is enabled for this while loop.
	BackProp bool `protobuf:"varint,3,opt,name=back_prop,json=backProp,proto3" json:"back_prop,omitempty"`
	// Whether GPU-CPU memory swap is enabled for this loop.
	SwapMemory bool `protobuf:"varint,4,opt,name=swap_memory,json=swapMemory,proto3" json:"swap_memory,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,5,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Name of the pivot_for_pred tensor.
	PivotForPredName string `protobuf:"bytes,6,opt,name=pivot_for_pred_name,json=pivotForPredName,proto3" json:"pivot_for_pred_name,omitempty"`
	// Name of the pivot_for_body tensor.
	PivotForBodyName string `protobuf:"bytes,7,opt,name=pivot_for_body_name,json=pivotForBodyName,proto3" json:"pivot_for_body_name,omitempty"`
	// List of names for exit tensors.
	LoopExitNames []string `protobuf:"bytes,8,rep,name=loop_exit_names,json=loopExitNames,proto3" json:"loop_exit_names,omitempty"`
	// List of names for enter tensors.
	LoopEnterNames []string `protobuf:"bytes,10,rep,name=loop_enter_names,json=loopEnterNames,proto3" json:"loop_enter_names,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,9,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Optional name of the maximum_iterations tensor.
	MaximumIterationsName string `protobuf:"bytes,11,opt,name=maximum_iterations_name,json=maximumIterationsName,proto3" json:"maximum_iterations_name,omitempty"`
	// Contexts contained inside this context (e.g. nested whiles).
	NestedContexts       []*ControlFlowContextDef `protobuf:"bytes,12,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WhileContextDef) Reset()         { *m = WhileContextDef{} }
func (m *WhileContextDef) String() string { return proto.CompactTextString(m) }
func (*WhileContextDef) ProtoMessage()    {}
func (*WhileContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_flow_ed32bd4a24b0fb72, []int{3}
}
func (m *WhileContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhileContextDef.Unmarshal(m, b)
}
func (m *WhileContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhileContextDef.Marshal(b, m, deterministic)
}
func (dst *WhileContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhileContextDef.Merge(dst, src)
}
func (m *WhileContextDef) XXX_Size() int {
	return xxx_messageInfo_WhileContextDef.Size(m)
}
func (m *WhileContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_WhileContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_WhileContextDef proto.InternalMessageInfo

func (m *WhileContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *WhileContextDef) GetParallelIterations() int32 {
	if m != nil {
		return m.ParallelIterations
	}
	return 0
}

func (m *WhileContextDef) GetBackProp() bool {
	if m != nil {
		return m.BackProp
	}
	return false
}

func (m *WhileContextDef) GetSwapMemory() bool {
	if m != nil {
		return m.SwapMemory
	}
	return false
}

func (m *WhileContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForPredName() string {
	if m != nil {
		return m.PivotForPredName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForBodyName() string {
	if m != nil {
		return m.PivotForBodyName
	}
	return ""
}

func (m *WhileContextDef) GetLoopExitNames() []string {
	if m != nil {
		return m.LoopExitNames
	}
	return nil
}

func (m *WhileContextDef) GetLoopEnterNames() []string {
	if m != nil {
		return m.LoopEnterNames
	}
	return nil
}

func (m *WhileContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

func (m *WhileContextDef) GetMaximumIterationsName() string {
	if m != nil {
		return m.MaximumIterationsName
	}
	return ""
}

func (m *WhileContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if m != nil {
		return m.NestedContexts
	}
	return nil
}

func init() {
	proto.RegisterType((*ValuesDef)(nil), "tensorflow.ValuesDef")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.ValuesDef.ExternalValuesEntry")
	proto.RegisterType((*ControlFlowContextDef)(nil), "tensorflow.ControlFlowContextDef")
	proto.RegisterType((*CondContextDef)(nil), "tensorflow.CondContextDef")
	proto.RegisterType((*WhileContextDef)(nil), "tensorflow.WhileContextDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/control_flow.proto", fileDescriptor_control_flow_ed32bd4a24b0fb72)
}

var fileDescriptor_control_flow_ed32bd4a24b0fb72 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xff, 0x9c, 0x34, 0xf9, 0xe2, 0x49, 0x49, 0x8a, 0x4b, 0x8b, 0xd5, 0x0a, 0x91, 0xe6, 0x80,
	0x82, 0x10, 0x89, 0x54, 0x10, 0x02, 0xc4, 0x85, 0x94, 0x56, 0x05, 0x09, 0x14, 0xf9, 0x00, 0x12,
	0x97, 0xd5, 0xc6, 0xde, 0xa4, 0x56, 0x6d, 0x8f, 0xb5, 0xde, 0x34, 0x8e, 0xc4, 0x0b, 0xf0, 0x06,
	0x3c, 0x04, 0x4f, 0xc0, 0x93, 0x71, 0x44, 0x3b, 0xeb, 0x36, 0x71, 0xc8, 0xa1, 0xe2, 0x12, 0xed,
	0xfe, 0xfe, 0xac, 0x67, 0x7e, 0xb3, 0x59, 0x78, 0xa2, 0x44, 0x92, 0xa1, 0x9c, 0x44, 0x38, 0x1f,
	0xf8, 0x28, 0xc5, 0x20, 0x95, 0xa8, 0x70, 0x3c, 0x9b, 0x0c, 0x7c, 0x4c, 0x94, 0xc4, 0x88, 0x69,
	0xaa, 0x4f, 0xa8, 0x03, 0x4b, 0x71, 0xf7, 0x97, 0x05, 0xf6, 0x67, 0x1e, 0xcd, 0x44, 0xf6, 0x4e,
	0x4c, 0x9c, 0x7d, 0xa8, 0x5f, 0xd1, 0xc6, 0xb5, 0x3a, 0xd5, 0x9e, 0xed, 0x15, 0x3b, 0xc7, 0x83,
	0xb6, 0xc8, 0x95, 0x90, 0x09, 0x8f, 0x58, 0x21, 0xa8, 0x74, 0xaa, 0xbd, 0xe6, 0xf1, 0xe3, 0xfe,
	0xf2, 0xac, 0xfe, 0xcd, 0x39, 0xfd, 0xd3, 0x42, 0x6c, 0x90, 0xd3, 0x44, 0xc9, 0x85, 0xd7, 0x12,
	0x25, 0xf0, 0xe0, 0x2d, 0xec, 0x6e, 0x90, 0x39, 0x3b, 0x50, 0xbd, 0x14, 0x0b, 0xd7, 0xea, 0x58,
	0x3d, 0xdb, 0xd3, 0x4b, 0xe7, 0x1e, 0xd4, 0xe8, 0x9b, 0x6e, 0x85, 0x30, 0xb3, 0x79, 0x5d, 0x79,
	0x69, 0x75, 0x7f, 0x58, 0xb0, 0x77, 0x62, 0xfa, 0x3b, 0x8b, 0x70, 0xae, 0x97, 0x22, 0x57, 0xba,
	0x91, 0x57, 0x60, 0xfb, 0x98, 0x04, 0xcc, 0x57, 0xb9, 0xa2, 0xb3, 0x9a, 0xc7, 0x07, 0xab, 0xa5,
	0x9e, 0x60, 0x12, 0x2c, 0xe5, 0xe7, 0xff, 0x79, 0x0d, 0x2d, 0x3f, 0x51, 0xb9, 0x72, 0xde, 0x00,
	0xcc, 0x2f, 0xc2, 0x48, 0x18, 0x6f, 0x85, 0xbc, 0x87, 0xab, 0xde, 0x2f, 0x9a, 0x2d, 0x99, 0x6d,
	0x32, 0x68, 0xf7, 0xb0, 0x0e, 0x5b, 0xda, 0xd7, 0xfd, 0x5e, 0x81, 0x56, 0xf9, 0x23, 0xce, 0x11,
	0x6c, 0xfb, 0x66, 0xc7, 0x12, 0x1e, 0x8b, 0xa2, 0xc5, 0x66, 0x81, 0x7d, 0xe2, 0xb1, 0x70, 0x0e,
	0xc1, 0x4e, 0xa5, 0x08, 0x0c, 0x6f, 0xda, 0x6d, 0x68, 0x80, 0xc8, 0x07, 0x00, 0x69, 0x78, 0x85,
	0x85, 0xbb, 0x4a, 0xac, 0x4d, 0x08, 0xd1, 0xfb, 0x50, 0x1f, 0x4b, 0x9e, 0xf8, 0x17, 0xee, 0x56,
	0xc7, 0xea, 0xd5, 0xbc, 0x62, 0xe7, 0x3c, 0x07, 0x30, 0x23, 0x63, 0x81, 0x98, 0xb8, 0x35, 0xea,
	0x67, 0x6f, 0xe3, 0xd8, 0x3c, 0xfb, 0xea, 0xe6, 0x26, 0x7c, 0x80, 0x76, 0x22, 0x32, 0x25, 0x02,
	0x56, 0xd4, 0x97, 0xb9, 0x75, 0x9a, 0xf8, 0xd1, 0x5a, 0x8c, 0x7f, 0x87, 0xef, 0xb5, 0x8c, 0xb3,
	0x40, 0xb2, 0xee, 0xcf, 0x2d, 0x68, 0xaf, 0x85, 0x76, 0x9b, 0x30, 0x06, 0xb0, 0x9b, 0x72, 0xc9,
	0xa3, 0x48, 0x44, 0x2c, 0x54, 0x42, 0x72, 0x15, 0x62, 0x92, 0x51, 0x2c, 0x35, 0xcf, 0xb9, 0xa6,
	0xde, 0xdf, 0x30, 0x3a, 0xbd, 0x31, 0xf7, 0x2f, 0x59, 0x2a, 0x31, 0xa5, 0x7c, 0x1a, 0x5e, 0x43,
	0x03, 0x23, 0x89, 0xa9, 0xf3, 0x10, 0x9a, 0xd9, 0x9c, 0xa7, 0x2c, 0x16, 0x31, 0xca, 0x05, 0x65,
	0xd4, 0xf0, 0x40, 0x43, 0x1f, 0x09, 0x59, 0x8b, 0xb7, 0xb6, 0x1e, 0xef, 0x53, 0xd8, 0x35, 0xf4,
	0x04, 0x25, 0x5b, 0x0e, 0xa9, 0x4e, 0xba, 0x1d, 0xa2, 0xce, 0x50, 0x8e, 0xae, 0x87, 0x55, 0x92,
	0x8f, 0x31, 0x58, 0x18, 0xf9, 0xff, 0x65, 0xf9, 0x10, 0x83, 0x05, 0xc9, 0x1f, 0x41, 0x3b, 0x42,
	0x4c, 0x99, 0xc8, 0x43, 0x53, 0x40, 0xe6, 0x36, 0xe8, 0x1f, 0x78, 0x47, 0xc3, 0xa7, 0x79, 0x48,
	0x45, 0x64, 0x4e, 0x0f, 0x76, 0x8c, 0x2e, 0x51, 0x42, 0x16, 0x42, 0x20, 0x61, 0x8b, 0x84, 0x1a,
	0x36, 0xca, 0xf2, 0xd8, 0xed, 0x5b, 0x8e, 0xfd, 0x05, 0xdc, 0x8f, 0x79, 0x1e, 0xc6, 0xb3, 0x78,
	0x25, 0x72, 0x53, 0x7a, 0x93, 0x4a, 0xdf, 0x2b, 0xe8, 0x65, 0xec, 0x54, 0xff, 0x86, 0xeb, 0xb2,
	0xfd, 0x8f, 0xd7, 0x65, 0xf8, 0x0d, 0x5c, 0x94, 0xd3, 0x55, 0xdf, 0x44, 0xf2, 0x58, 0xcc, 0x51,
	0x5e, 0x0e, 0xef, 0xae, 0x1c, 0x31, 0xd2, 0x8f, 0x59, 0x36, 0xb2, 0xbe, 0x9e, 0x4f, 0x43, 0x75,
	0x31, 0x1b, 0xf7, 0x7d, 0x8c, 0x07, 0x2b, 0xef, 0xe0, 0xe6, 0xe5, 0x14, 0xcd, 0x03, 0xa9, 0x7f,
	0x18, 0xbd, 0x87, 0x19, 0x9b, 0xa2, 0x59, 0xfd, 0xb6, 0xac, 0x71, 0x9d, 0x56, 0xcf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x59, 0x85, 0x33, 0xd1, 0x52, 0x05, 0x00, 0x00,
}
