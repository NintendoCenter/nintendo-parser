// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: internal/protos/game.proto

package protos

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

type Shop int32

const (
	Shop_NINTENDO Shop = 0
	Shop_SAVELA   Shop = 1
	Shop_INTEREST Shop = 2
)

// Enum value maps for Shop.
var (
	Shop_name = map[int32]string{
		0: "NINTENDO",
		1: "SAVELA",
		2: "INTEREST",
	}
	Shop_value = map[string]int32{
		"NINTENDO": 0,
		"SAVELA":   1,
		"INTEREST": 2,
	}
)

func (x Shop) Enum() *Shop {
	p := new(Shop)
	*p = x
	return p
}

func (x Shop) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Shop) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_game_proto_enumTypes[0].Descriptor()
}

func (Shop) Type() protoreflect.EnumType {
	return &file_internal_protos_game_proto_enumTypes[0]
}

func (x Shop) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Shop.Descriptor instead.
func (Shop) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_game_proto_rawDescGZIP(), []int{0}
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ImageUrl    string   `protobuf:"bytes,4,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Offers      []*Offer `protobuf:"bytes,5,rep,name=Offers,proto3" json:"Offers,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_internal_protos_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Game) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Game) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Game) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Original   float32 `protobuf:"fixed32,1,opt,name=original,proto3" json:"original,omitempty"`
	Discounted float32 `protobuf:"fixed32,2,opt,name=discounted,proto3" json:"discounted,omitempty"`
	Real       float32 `protobuf:"fixed32,3,opt,name=real,proto3" json:"real,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_internal_protos_game_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetOriginal() float32 {
	if x != nil {
		return x.Original
	}
	return 0
}

func (x *Price) GetDiscounted() float32 {
	if x != nil {
		return x.Discounted
	}
	return 0
}

func (x *Price) GetReal() float32 {
	if x != nil {
		return x.Real
	}
	return 0
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shop      Shop   `protobuf:"varint,1,opt,name=shop,proto3,enum=Shop" json:"shop,omitempty"`
	IsDigital bool   `protobuf:"varint,2,opt,name=isDigital,proto3" json:"isDigital,omitempty"`
	IsUsed    bool   `protobuf:"varint,3,opt,name=isUsed,proto3" json:"isUsed,omitempty"`
	Link      string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Price     *Price `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_internal_protos_game_proto_rawDescGZIP(), []int{2}
}

func (x *Offer) GetShop() Shop {
	if x != nil {
		return x.Shop
	}
	return Shop_NINTENDO
}

func (x *Offer) GetIsDigital() bool {
	if x != nil {
		return x.IsDigital
	}
	return false
}

func (x *Offer) GetIsUsed() bool {
	if x != nil {
		return x.IsUsed
	}
	return false
}

func (x *Offer) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Offer) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

var File_internal_protos_game_proto protoreflect.FileDescriptor

var file_internal_protos_game_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0x57, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x65,
	0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04,
	0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2a,
	0x2e, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x49, 0x4e, 0x54, 0x45,
	0x4e, 0x44, 0x4f, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x41, 0x56, 0x45, 0x4c, 0x41, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x02, 0x42,
	0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protos_game_proto_rawDescOnce sync.Once
	file_internal_protos_game_proto_rawDescData = file_internal_protos_game_proto_rawDesc
)

func file_internal_protos_game_proto_rawDescGZIP() []byte {
	file_internal_protos_game_proto_rawDescOnce.Do(func() {
		file_internal_protos_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protos_game_proto_rawDescData)
	})
	return file_internal_protos_game_proto_rawDescData
}

var file_internal_protos_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_protos_game_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_protos_game_proto_goTypes = []interface{}{
	(Shop)(0),     // 0: Shop
	(*Game)(nil),  // 1: Game
	(*Price)(nil), // 2: Price
	(*Offer)(nil), // 3: Offer
}
var file_internal_protos_game_proto_depIdxs = []int32{
	3, // 0: Game.Offers:type_name -> Offer
	0, // 1: Offer.shop:type_name -> Shop
	2, // 2: Offer.price:type_name -> Price
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_protos_game_proto_init() }
func file_internal_protos_game_proto_init() {
	if File_internal_protos_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protos_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_internal_protos_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_internal_protos_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
			RawDescriptor: file_internal_protos_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protos_game_proto_goTypes,
		DependencyIndexes: file_internal_protos_game_proto_depIdxs,
		EnumInfos:         file_internal_protos_game_proto_enumTypes,
		MessageInfos:      file_internal_protos_game_proto_msgTypes,
	}.Build()
	File_internal_protos_game_proto = out.File
	file_internal_protos_game_proto_rawDesc = nil
	file_internal_protos_game_proto_goTypes = nil
	file_internal_protos_game_proto_depIdxs = nil
}