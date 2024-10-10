// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: omni/specs/infra.proto

package specs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MachineRequestStatusSpec_Stage int32

const (
	MachineRequestStatusSpec_UNKNOWN      MachineRequestStatusSpec_Stage = 0
	MachineRequestStatusSpec_PROVISIONING MachineRequestStatusSpec_Stage = 1
	MachineRequestStatusSpec_PROVISIONED  MachineRequestStatusSpec_Stage = 2
	MachineRequestStatusSpec_FAILED       MachineRequestStatusSpec_Stage = 3
)

// Enum value maps for MachineRequestStatusSpec_Stage.
var (
	MachineRequestStatusSpec_Stage_name = map[int32]string{
		0: "UNKNOWN",
		1: "PROVISIONING",
		2: "PROVISIONED",
		3: "FAILED",
	}
	MachineRequestStatusSpec_Stage_value = map[string]int32{
		"UNKNOWN":      0,
		"PROVISIONING": 1,
		"PROVISIONED":  2,
		"FAILED":       3,
	}
)

func (x MachineRequestStatusSpec_Stage) Enum() *MachineRequestStatusSpec_Stage {
	p := new(MachineRequestStatusSpec_Stage)
	*p = x
	return p
}

func (x MachineRequestStatusSpec_Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineRequestStatusSpec_Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_omni_specs_infra_proto_enumTypes[0].Descriptor()
}

func (MachineRequestStatusSpec_Stage) Type() protoreflect.EnumType {
	return &file_omni_specs_infra_proto_enumTypes[0]
}

func (x MachineRequestStatusSpec_Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineRequestStatusSpec_Stage.Descriptor instead.
func (MachineRequestStatusSpec_Stage) EnumDescriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{1, 0}
}

type MachineRequestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TalosVersion string         `protobuf:"bytes,1,opt,name=talos_version,json=talosVersion,proto3" json:"talos_version,omitempty"`
	Overlay      *Overlay       `protobuf:"bytes,2,opt,name=overlay,proto3" json:"overlay,omitempty"`
	Extensions   []string       `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty"`
	KernelArgs   []string       `protobuf:"bytes,4,rep,name=kernel_args,json=kernelArgs,proto3" json:"kernel_args,omitempty"`
	MetaValues   []*MetaValue   `protobuf:"bytes,5,rep,name=meta_values,json=metaValues,proto3" json:"meta_values,omitempty"`
	ProviderData string         `protobuf:"bytes,6,opt,name=provider_data,json=providerData,proto3" json:"provider_data,omitempty"`
	GrpcTunnel   GrpcTunnelMode `protobuf:"varint,7,opt,name=grpc_tunnel,json=grpcTunnel,proto3,enum=specs.GrpcTunnelMode" json:"grpc_tunnel,omitempty"`
}

func (x *MachineRequestSpec) Reset() {
	*x = MachineRequestSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineRequestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineRequestSpec) ProtoMessage() {}

func (x *MachineRequestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineRequestSpec.ProtoReflect.Descriptor instead.
func (*MachineRequestSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{0}
}

func (x *MachineRequestSpec) GetTalosVersion() string {
	if x != nil {
		return x.TalosVersion
	}
	return ""
}

func (x *MachineRequestSpec) GetOverlay() *Overlay {
	if x != nil {
		return x.Overlay
	}
	return nil
}

func (x *MachineRequestSpec) GetExtensions() []string {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *MachineRequestSpec) GetKernelArgs() []string {
	if x != nil {
		return x.KernelArgs
	}
	return nil
}

func (x *MachineRequestSpec) GetMetaValues() []*MetaValue {
	if x != nil {
		return x.MetaValues
	}
	return nil
}

func (x *MachineRequestSpec) GetProviderData() string {
	if x != nil {
		return x.ProviderData
	}
	return ""
}

func (x *MachineRequestSpec) GetGrpcTunnel() GrpcTunnelMode {
	if x != nil {
		return x.GrpcTunnel
	}
	return GrpcTunnelMode_UNSET
}

type MachineRequestStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stage  MachineRequestStatusSpec_Stage `protobuf:"varint,2,opt,name=stage,proto3,enum=specs.MachineRequestStatusSpec_Stage" json:"stage,omitempty"`
	Error  string                         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Status string                         `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MachineRequestStatusSpec) Reset() {
	*x = MachineRequestStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineRequestStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineRequestStatusSpec) ProtoMessage() {}

func (x *MachineRequestStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineRequestStatusSpec.ProtoReflect.Descriptor instead.
func (*MachineRequestStatusSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{1}
}

func (x *MachineRequestStatusSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MachineRequestStatusSpec) GetStage() MachineRequestStatusSpec_Stage {
	if x != nil {
		return x.Stage
	}
	return MachineRequestStatusSpec_UNKNOWN
}

func (x *MachineRequestStatusSpec) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MachineRequestStatusSpec) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InfraProviderStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema      string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Icon        string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *InfraProviderStatusSpec) Reset() {
	*x = InfraProviderStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_infra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfraProviderStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraProviderStatusSpec) ProtoMessage() {}

func (x *InfraProviderStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraProviderStatusSpec.ProtoReflect.Descriptor instead.
func (*InfraProviderStatusSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{2}
}

func (x *InfraProviderStatusSpec) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

var File_omni_specs_infra_proto protoreflect.FileDescriptor

var file_omni_specs_infra_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a,
	0x15, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x31, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xda, 0x01,
	0x0a, 0x18, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x7b, 0x0a, 0x17, 0x49, 0x6e,
	0x66, 0x72, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_omni_specs_infra_proto_rawDescOnce sync.Once
	file_omni_specs_infra_proto_rawDescData = file_omni_specs_infra_proto_rawDesc
)

func file_omni_specs_infra_proto_rawDescGZIP() []byte {
	file_omni_specs_infra_proto_rawDescOnce.Do(func() {
		file_omni_specs_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_omni_specs_infra_proto_rawDescData)
	})
	return file_omni_specs_infra_proto_rawDescData
}

var file_omni_specs_infra_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_omni_specs_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_omni_specs_infra_proto_goTypes = []any{
	(MachineRequestStatusSpec_Stage)(0), // 0: specs.MachineRequestStatusSpec.Stage
	(*MachineRequestSpec)(nil),          // 1: specs.MachineRequestSpec
	(*MachineRequestStatusSpec)(nil),    // 2: specs.MachineRequestStatusSpec
	(*InfraProviderStatusSpec)(nil),     // 3: specs.InfraProviderStatusSpec
	(*Overlay)(nil),                     // 4: specs.Overlay
	(*MetaValue)(nil),                   // 5: specs.MetaValue
	(GrpcTunnelMode)(0),                 // 6: specs.GrpcTunnelMode
}
var file_omni_specs_infra_proto_depIdxs = []int32{
	4, // 0: specs.MachineRequestSpec.overlay:type_name -> specs.Overlay
	5, // 1: specs.MachineRequestSpec.meta_values:type_name -> specs.MetaValue
	6, // 2: specs.MachineRequestSpec.grpc_tunnel:type_name -> specs.GrpcTunnelMode
	0, // 3: specs.MachineRequestStatusSpec.stage:type_name -> specs.MachineRequestStatusSpec.Stage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_omni_specs_infra_proto_init() }
func file_omni_specs_infra_proto_init() {
	if File_omni_specs_infra_proto != nil {
		return
	}
	file_omni_specs_omni_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_omni_specs_infra_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MachineRequestSpec); i {
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
		file_omni_specs_infra_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MachineRequestStatusSpec); i {
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
		file_omni_specs_infra_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InfraProviderStatusSpec); i {
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
			RawDescriptor: file_omni_specs_infra_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omni_specs_infra_proto_goTypes,
		DependencyIndexes: file_omni_specs_infra_proto_depIdxs,
		EnumInfos:         file_omni_specs_infra_proto_enumTypes,
		MessageInfos:      file_omni_specs_infra_proto_msgTypes,
	}.Build()
	File_omni_specs_infra_proto = out.File
	file_omni_specs_infra_proto_rawDesc = nil
	file_omni_specs_infra_proto_goTypes = nil
	file_omni_specs_infra_proto_depIdxs = nil
}
