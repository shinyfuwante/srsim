// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb/model/sim.proto

package model

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

type SimConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings   *SimulatorSettings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Characters []*Character       `protobuf:"bytes,2,rep,name=characters,proto3" json:"characters,omitempty"`
	Enemies    []*Enemy           `protobuf:"bytes,3,rep,name=enemies,proto3" json:"enemies,omitempty"` // TODO: waves
	Engage     *Engage            `protobuf:"bytes,6,opt,name=engage,proto3" json:"engage,omitempty"`
	// Types that are assignable to Logic:
	//
	//	*SimConfig_Gcsl
	Logic isSimConfig_Logic `protobuf_oneof:"logic"`
}

func (x *SimConfig) Reset() {
	*x = SimConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimConfig) ProtoMessage() {}

func (x *SimConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimConfig.ProtoReflect.Descriptor instead.
func (*SimConfig) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{0}
}

func (x *SimConfig) GetSettings() *SimulatorSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *SimConfig) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *SimConfig) GetEnemies() []*Enemy {
	if x != nil {
		return x.Enemies
	}
	return nil
}

func (x *SimConfig) GetEngage() *Engage {
	if x != nil {
		return x.Engage
	}
	return nil
}

func (m *SimConfig) GetLogic() isSimConfig_Logic {
	if m != nil {
		return m.Logic
	}
	return nil
}

func (x *SimConfig) GetGcsl() string {
	if x, ok := x.GetLogic().(*SimConfig_Gcsl); ok {
		return x.Gcsl
	}
	return ""
}

type isSimConfig_Logic interface {
	isSimConfig_Logic()
}

type SimConfig_Gcsl struct {
	Gcsl string `protobuf:"bytes,7,opt,name=gcsl,proto3,oneof"`
}

func (*SimConfig_Gcsl) isSimConfig_Logic() {}

type SimulatorSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CycleLimit uint32 `protobuf:"varint,1,opt,name=cycle_limit,json=cycleLimit,proto3" json:"cycle_limit,omitempty"`
}

func (x *SimulatorSettings) Reset() {
	*x = SimulatorSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulatorSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorSettings) ProtoMessage() {}

func (x *SimulatorSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorSettings.ProtoReflect.Descriptor instead.
func (*SimulatorSettings) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{1}
}

func (x *SimulatorSettings) GetCycleLimit() uint32 {
	if x != nil {
		return x.CycleLimit
	}
	return 0
}

type Wave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enemies []*Enemy `protobuf:"bytes,1,rep,name=enemies,proto3" json:"enemies,omitempty"`
}

func (x *Wave) Reset() {
	*x = Wave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wave) ProtoMessage() {}

func (x *Wave) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wave.ProtoReflect.Descriptor instead.
func (*Wave) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{2}
}

func (x *Wave) GetEnemies() []*Enemy {
	if x != nil {
		return x.Enemies
	}
	return nil
}

type Engage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ambush bool `protobuf:"varint,1,opt,name=ambush,proto3" json:"ambush,omitempty"`
}

func (x *Engage) Reset() {
	*x = Engage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_sim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engage) ProtoMessage() {}

func (x *Engage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_sim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engage.ProtoReflect.Descriptor instead.
func (*Engage) Descriptor() ([]byte, []int) {
	return file_pb_model_sim_proto_rawDescGZIP(), []int{3}
}

func (x *Engage) GetAmbush() bool {
	if x != nil {
		return x.Ambush
	}
	return false
}

var File_pb_model_sim_proto protoreflect.FileDescriptor

var file_pb_model_sim_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x14, 0x70, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x09,
	0x53, 0x69, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x30, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x67,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x04, 0x67, 0x63, 0x73, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x67, 0x63, 0x73, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x34, 0x0a, 0x11, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x2e, 0x0a, 0x04, 0x57, 0x61, 0x76, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e, 0x65,
	0x6d, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65,
	0x73, 0x22, 0x20, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x62, 0x75, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x6d, 0x62,
	0x75, 0x73, 0x68, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x2f, 0x73, 0x72, 0x73, 0x69,
	0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_model_sim_proto_rawDescOnce sync.Once
	file_pb_model_sim_proto_rawDescData = file_pb_model_sim_proto_rawDesc
)

func file_pb_model_sim_proto_rawDescGZIP() []byte {
	file_pb_model_sim_proto_rawDescOnce.Do(func() {
		file_pb_model_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_sim_proto_rawDescData)
	})
	return file_pb_model_sim_proto_rawDescData
}

var file_pb_model_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_model_sim_proto_goTypes = []interface{}{
	(*SimConfig)(nil),         // 0: model.SimConfig
	(*SimulatorSettings)(nil), // 1: model.SimulatorSettings
	(*Wave)(nil),              // 2: model.Wave
	(*Engage)(nil),            // 3: model.Engage
	(*Character)(nil),         // 4: model.Character
	(*Enemy)(nil),             // 5: model.Enemy
}
var file_pb_model_sim_proto_depIdxs = []int32{
	1, // 0: model.SimConfig.settings:type_name -> model.SimulatorSettings
	4, // 1: model.SimConfig.characters:type_name -> model.Character
	5, // 2: model.SimConfig.enemies:type_name -> model.Enemy
	3, // 3: model.SimConfig.engage:type_name -> model.Engage
	5, // 4: model.Wave.enemies:type_name -> model.Enemy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_model_sim_proto_init() }
func file_pb_model_sim_proto_init() {
	if File_pb_model_sim_proto != nil {
		return
	}
	file_pb_model_enemy_proto_init()
	file_pb_model_character_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_model_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimConfig); i {
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
		file_pb_model_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulatorSettings); i {
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
		file_pb_model_sim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wave); i {
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
		file_pb_model_sim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Engage); i {
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
	file_pb_model_sim_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SimConfig_Gcsl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_model_sim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_sim_proto_goTypes,
		DependencyIndexes: file_pb_model_sim_proto_depIdxs,
		MessageInfos:      file_pb_model_sim_proto_msgTypes,
	}.Build()
	File_pb_model_sim_proto = out.File
	file_pb_model_sim_proto_rawDesc = nil
	file_pb_model_sim_proto_goTypes = nil
	file_pb_model_sim_proto_depIdxs = nil
}
