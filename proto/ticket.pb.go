// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0--rc1
// source: ticket.proto

package proto

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

type TicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Users *User   `protobuf:"bytes,3,opt,name=users,proto3" json:"users,omitempty"`
	Price float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TicketRequest) Reset() {
	*x = TicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequest) ProtoMessage() {}

func (x *TicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequest.ProtoReflect.Descriptor instead.
func (*TicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *TicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TicketRequest) GetUsers() *User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *TicketRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt *TicketReceipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Status  *TicketStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TicketResponse) Reset() {
	*x = TicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResponse) ProtoMessage() {}

func (x *TicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResponse.ProtoReflect.Descriptor instead.
func (*TicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *TicketResponse) GetReceipt() *TicketReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *TicketResponse) GetStatus() *TicketStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type TicketReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seatno  *TicketNumber  `protobuf:"bytes,1,opt,name=seatno,proto3" json:"seatno,omitempty"`
	Coachno *TicketSection `protobuf:"bytes,2,opt,name=coachno,proto3" json:"coachno,omitempty"`
}

func (x *TicketReceipt) Reset() {
	*x = TicketReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketReceipt) ProtoMessage() {}

func (x *TicketReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketReceipt.ProtoReflect.Descriptor instead.
func (*TicketReceipt) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *TicketReceipt) GetSeatno() *TicketNumber {
	if x != nil {
		return x.Seatno
	}
	return nil
}

func (x *TicketReceipt) GetCoachno() *TicketSection {
	if x != nil {
		return x.Coachno
	}
	return nil
}

type TicketNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket int32 `protobuf:"varint,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *TicketNumber) Reset() {
	*x = TicketNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketNumber) ProtoMessage() {}

func (x *TicketNumber) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketNumber.ProtoReflect.Descriptor instead.
func (*TicketNumber) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *TicketNumber) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

type TicketSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *TicketSection) Reset() {
	*x = TicketSection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketSection) ProtoMessage() {}

func (x *TicketSection) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketSection.ProtoReflect.Descriptor instead.
func (*TicketSection) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *TicketSection) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type TicketStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TicketStatus) Reset() {
	*x = TicketStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketStatus) ProtoMessage() {}

func (x *TicketStatus) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketStatus.ProtoReflect.Descriptor instead.
func (*TicketStatus) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *TicketStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TicketDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userdetails *TicketRequest  `protobuf:"bytes,1,opt,name=userdetails,proto3" json:"userdetails,omitempty"`
	Status      *TicketResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TicketDetails) Reset() {
	*x = TicketDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketDetails) ProtoMessage() {}

func (x *TicketDetails) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketDetails.ProtoReflect.Descriptor instead.
func (*TicketDetails) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *TicketDetails) GetUserdetails() *TicketRequest {
	if x != nil {
		return x.Userdetails
	}
	return nil
}

func (x *TicketDetails) GetStatus() *TicketResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x73, 0x0a, 0x0d,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x7b, 0x0a, 0x0e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7a, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x74, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x61, 0x74, 0x6e, 0x6f, 0x12, 0x35, 0x0a, 0x07, 0x63,
	0x6f, 0x61, 0x63, 0x68, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x6e, 0x6f, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x84, 0x01,
	0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x3d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xfc, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1b, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0a, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x75, 0x6b, 0x65, 0x73, 0x68, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ticket_proto_goTypes = []interface{}{
	(*TicketRequest)(nil),  // 0: train_ticket.TicketRequest
	(*User)(nil),           // 1: train_ticket.User
	(*TicketResponse)(nil), // 2: train_ticket.TicketResponse
	(*TicketReceipt)(nil),  // 3: train_ticket.TicketReceipt
	(*TicketNumber)(nil),   // 4: train_ticket.TicketNumber
	(*TicketSection)(nil),  // 5: train_ticket.TicketSection
	(*TicketStatus)(nil),   // 6: train_ticket.TicketStatus
	(*TicketDetails)(nil),  // 7: train_ticket.TicketDetails
}
var file_ticket_proto_depIdxs = []int32{
	1,  // 0: train_ticket.TicketRequest.users:type_name -> train_ticket.User
	3,  // 1: train_ticket.TicketResponse.receipt:type_name -> train_ticket.TicketReceipt
	6,  // 2: train_ticket.TicketResponse.status:type_name -> train_ticket.TicketStatus
	4,  // 3: train_ticket.TicketReceipt.seatno:type_name -> train_ticket.TicketNumber
	5,  // 4: train_ticket.TicketReceipt.coachno:type_name -> train_ticket.TicketSection
	0,  // 5: train_ticket.TicketDetails.userdetails:type_name -> train_ticket.TicketRequest
	2,  // 6: train_ticket.TicketDetails.status:type_name -> train_ticket.TicketResponse
	0,  // 7: train_ticket.TrainBooking.BookTicket:input_type -> train_ticket.TicketRequest
	3,  // 8: train_ticket.TrainBooking.GetReceipt:input_type -> train_ticket.TicketReceipt
	3,  // 9: train_ticket.TrainBooking.DeleteTicket:input_type -> train_ticket.TicketReceipt
	5,  // 10: train_ticket.TrainBooking.GetAllUsers:input_type -> train_ticket.TicketSection
	3,  // 11: train_ticket.TrainBooking.ModifySeat:input_type -> train_ticket.TicketReceipt
	2,  // 12: train_ticket.TrainBooking.BookTicket:output_type -> train_ticket.TicketResponse
	7,  // 13: train_ticket.TrainBooking.GetReceipt:output_type -> train_ticket.TicketDetails
	6,  // 14: train_ticket.TrainBooking.DeleteTicket:output_type -> train_ticket.TicketStatus
	7,  // 15: train_ticket.TrainBooking.GetAllUsers:output_type -> train_ticket.TicketDetails
	2,  // 16: train_ticket.TrainBooking.ModifySeat:output_type -> train_ticket.TicketResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketRequest); i {
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
		file_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketResponse); i {
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
		file_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketReceipt); i {
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
		file_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketNumber); i {
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
		file_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketSection); i {
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
		file_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketStatus); i {
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
		file_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketDetails); i {
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
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
