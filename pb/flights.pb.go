// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: flights.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int64                  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flights_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flights_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flights_proto_rawDescGZIP(), []int{0}
}

func (x *Flight) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Flight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flight) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Flight) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Flight) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Flight) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Flight) GetAvailableSlot() int64 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From string                 `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string                 `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Date *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flights_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flights_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_flights_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SearchRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SearchRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ListFlights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*Flight `protobuf:"bytes,1,rep,name=flights,proto3" json:"flights,omitempty"`
}

func (x *ListFlights) Reset() {
	*x = ListFlights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flights_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlights) ProtoMessage() {}

func (x *ListFlights) ProtoReflect() protoreflect.Message {
	mi := &file_flights_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlights.ProtoReflect.Descriptor instead.
func (*ListFlights) Descriptor() ([]byte, []int) {
	return file_flights_proto_rawDescGZIP(), []int{2}
}

func (x *ListFlights) GetFlights() []*Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flights_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_flights_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_flights_proto_rawDescGZIP(), []int{3}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_flights_proto protoreflect.FileDescriptor

var file_flights_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a,
	0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x77,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xef, 0x01, 0x0a, 0x07, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x36,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x46, 0x6c, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x46, 0x6c, 0x79, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flights_proto_rawDescOnce sync.Once
	file_flights_proto_rawDescData = file_flights_proto_rawDesc
)

func file_flights_proto_rawDescGZIP() []byte {
	file_flights_proto_rawDescOnce.Do(func() {
		file_flights_proto_rawDescData = protoimpl.X.CompressGZIP(file_flights_proto_rawDescData)
	})
	return file_flights_proto_rawDescData
}

var file_flights_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flights_proto_goTypes = []interface{}{
	(*Flight)(nil),                // 0: assignment.Flight
	(*SearchRequest)(nil),         // 1: assignment.SearchRequest
	(*ListFlights)(nil),           // 2: assignment.ListFlights
	(*ID)(nil),                    // 3: assignment.ID
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_flights_proto_depIdxs = []int32{
	4, // 0: assignment.Flight.date:type_name -> google.protobuf.Timestamp
	4, // 1: assignment.SearchRequest.date:type_name -> google.protobuf.Timestamp
	0, // 2: assignment.ListFlights.flights:type_name -> assignment.Flight
	0, // 3: assignment.Flights.CreateFlight:input_type -> assignment.Flight
	0, // 4: assignment.Flights.UpdateFlight:input_type -> assignment.Flight
	1, // 5: assignment.Flights.SearchFly:input_type -> assignment.SearchRequest
	3, // 6: assignment.Flights.SearchFlyByID:input_type -> assignment.ID
	0, // 7: assignment.Flights.CreateFlight:output_type -> assignment.Flight
	0, // 8: assignment.Flights.UpdateFlight:output_type -> assignment.Flight
	2, // 9: assignment.Flights.SearchFly:output_type -> assignment.ListFlights
	0, // 10: assignment.Flights.SearchFlyByID:output_type -> assignment.Flight
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flights_proto_init() }
func file_flights_proto_init() {
	if File_flights_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flights_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_flights_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_flights_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlights); i {
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
		file_flights_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
			RawDescriptor: file_flights_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flights_proto_goTypes,
		DependencyIndexes: file_flights_proto_depIdxs,
		MessageInfos:      file_flights_proto_msgTypes,
	}.Build()
	File_flights_proto = out.File
	file_flights_proto_rawDesc = nil
	file_flights_proto_goTypes = nil
	file_flights_proto_depIdxs = nil
}
