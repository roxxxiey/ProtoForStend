// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: updfwdevice.proto

package firmware_v1

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

type UPDFWTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UPDFWTypeRequest) Reset() {
	*x = UPDFWTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UPDFWTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UPDFWTypeRequest) ProtoMessage() {}

func (x *UPDFWTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UPDFWTypeRequest.ProtoReflect.Descriptor instead.
func (*UPDFWTypeRequest) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{0}
}

type UPDFWTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UPDFWTypeResponse) Reset() {
	*x = UPDFWTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UPDFWTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UPDFWTypeResponse) ProtoMessage() {}

func (x *UPDFWTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UPDFWTypeResponse.ProtoReflect.Descriptor instead.
func (*UPDFWTypeResponse) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{1}
}

func (x *UPDFWTypeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{2}
}

func (x *Settings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Settings) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateFirmwareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings []*Settings `protobuf:"bytes,1,rep,name=Settings,proto3" json:"Settings,omitempty"`
}

func (x *UpdateFirmwareRequest) Reset() {
	*x = UpdateFirmwareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFirmwareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFirmwareRequest) ProtoMessage() {}

func (x *UpdateFirmwareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFirmwareRequest.ProtoReflect.Descriptor instead.
func (*UpdateFirmwareRequest) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFirmwareRequest) GetSettings() []*Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateFirmwareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateFirmwareResponse) Reset() {
	*x = UpdateFirmwareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFirmwareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFirmwareResponse) ProtoMessage() {}

func (x *UpdateFirmwareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFirmwareResponse.ProtoReflect.Descriptor instead.
func (*UpdateFirmwareResponse) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFirmwareResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PresetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PresetRequest) Reset() {
	*x = PresetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetRequest) ProtoMessage() {}

func (x *PresetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetRequest.ProtoReflect.Descriptor instead.
func (*PresetRequest) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{5}
}

type PresetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Present string `protobuf:"bytes,1,opt,name=present,proto3" json:"present,omitempty"`
}

func (x *PresetResponse) Reset() {
	*x = PresetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updfwdevice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetResponse) ProtoMessage() {}

func (x *PresetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updfwdevice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetResponse.ProtoReflect.Descriptor instead.
func (*PresetResponse) Descriptor() ([]byte, []int) {
	return file_updfwdevice_proto_rawDescGZIP(), []int{6}
}

func (x *PresetResponse) GetPresent() string {
	if x != nil {
		return x.Present
	}
	return ""
}

var File_updfwdevice_proto protoreflect.FileDescriptor

var file_updfwdevice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x70, 0x64, 0x66, 0x77, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x55, 0x50, 0x44, 0x46, 0x57, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x55, 0x50, 0x44, 0x46, 0x57, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x58, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x30, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0f, 0x0a,
	0x0d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a,
	0x0a, 0x0e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x32, 0xce, 0x02, 0x0a, 0x0e, 0x46,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a,
	0x09, 0x55, 0x50, 0x44, 0x46, 0x57, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x50, 0x44, 0x46, 0x57, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x50, 0x44, 0x46, 0x57, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x12, 0x30, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x06,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x66,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_updfwdevice_proto_rawDescOnce sync.Once
	file_updfwdevice_proto_rawDescData = file_updfwdevice_proto_rawDesc
)

func file_updfwdevice_proto_rawDescGZIP() []byte {
	file_updfwdevice_proto_rawDescOnce.Do(func() {
		file_updfwdevice_proto_rawDescData = protoimpl.X.CompressGZIP(file_updfwdevice_proto_rawDescData)
	})
	return file_updfwdevice_proto_rawDescData
}

var file_updfwdevice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_updfwdevice_proto_goTypes = []interface{}{
	(*UPDFWTypeRequest)(nil),       // 0: DriverUpdateFirmwareProto.UPDFWTypeRequest
	(*UPDFWTypeResponse)(nil),      // 1: DriverUpdateFirmwareProto.UPDFWTypeResponse
	(*Settings)(nil),               // 2: DriverUpdateFirmwareProto.Settings
	(*UpdateFirmwareRequest)(nil),  // 3: DriverUpdateFirmwareProto.UpdateFirmwareRequest
	(*UpdateFirmwareResponse)(nil), // 4: DriverUpdateFirmwareProto.UpdateFirmwareResponse
	(*PresetRequest)(nil),          // 5: DriverUpdateFirmwareProto.PresetRequest
	(*PresetResponse)(nil),         // 6: DriverUpdateFirmwareProto.PresetResponse
}
var file_updfwdevice_proto_depIdxs = []int32{
	2, // 0: DriverUpdateFirmwareProto.UpdateFirmwareRequest.Settings:type_name -> DriverUpdateFirmwareProto.Settings
	0, // 1: DriverUpdateFirmwareProto.FirmwareDevice.UPDFWType:input_type -> DriverUpdateFirmwareProto.UPDFWTypeRequest
	3, // 2: DriverUpdateFirmwareProto.FirmwareDevice.UpdateFirmware:input_type -> DriverUpdateFirmwareProto.UpdateFirmwareRequest
	5, // 3: DriverUpdateFirmwareProto.FirmwareDevice.Preset:input_type -> DriverUpdateFirmwareProto.PresetRequest
	1, // 4: DriverUpdateFirmwareProto.FirmwareDevice.UPDFWType:output_type -> DriverUpdateFirmwareProto.UPDFWTypeResponse
	4, // 5: DriverUpdateFirmwareProto.FirmwareDevice.UpdateFirmware:output_type -> DriverUpdateFirmwareProto.UpdateFirmwareResponse
	6, // 6: DriverUpdateFirmwareProto.FirmwareDevice.Preset:output_type -> DriverUpdateFirmwareProto.PresetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_updfwdevice_proto_init() }
func file_updfwdevice_proto_init() {
	if File_updfwdevice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_updfwdevice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UPDFWTypeRequest); i {
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
		file_updfwdevice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UPDFWTypeResponse); i {
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
		file_updfwdevice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_updfwdevice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFirmwareRequest); i {
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
		file_updfwdevice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFirmwareResponse); i {
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
		file_updfwdevice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresetRequest); i {
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
		file_updfwdevice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresetResponse); i {
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
			RawDescriptor: file_updfwdevice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_updfwdevice_proto_goTypes,
		DependencyIndexes: file_updfwdevice_proto_depIdxs,
		MessageInfos:      file_updfwdevice_proto_msgTypes,
	}.Build()
	File_updfwdevice_proto = out.File
	file_updfwdevice_proto_rawDesc = nil
	file_updfwdevice_proto_goTypes = nil
	file_updfwdevice_proto_depIdxs = nil
}