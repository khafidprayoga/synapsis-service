// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: synapsis/v1/transaction.proto

package synapsisv1

import (
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

type TransactionState int32

const (
	TransactionState_TRANSACTION_STATE_UNSPECIFIED TransactionState = 0
	TransactionState_TRANSACTION_STATE_CART        TransactionState = 1
	TransactionState_TRANSACTION_STATE_PROCESS     TransactionState = 2
	TransactionState_TRANSACTION_STATE_PAID        TransactionState = 3
)

// Enum value maps for TransactionState.
var (
	TransactionState_name = map[int32]string{
		0: "TRANSACTION_STATE_UNSPECIFIED",
		1: "TRANSACTION_STATE_CART",
		2: "TRANSACTION_STATE_PROCESS",
		3: "TRANSACTION_STATE_PAID",
	}
	TransactionState_value = map[string]int32{
		"TRANSACTION_STATE_UNSPECIFIED": 0,
		"TRANSACTION_STATE_CART":        1,
		"TRANSACTION_STATE_PROCESS":     2,
		"TRANSACTION_STATE_PAID":        3,
	}
)

func (x TransactionState) Enum() *TransactionState {
	p := new(TransactionState)
	*p = x
	return p
}

func (x TransactionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionState) Descriptor() protoreflect.EnumDescriptor {
	return file_synapsis_v1_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionState) Type() protoreflect.EnumType {
	return &file_synapsis_v1_transaction_proto_enumTypes[0]
}

func (x TransactionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionState.Descriptor instead.
func (TransactionState) EnumDescriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{0}
}

// Transaction
// index: product.id, user.id, dt.created_at, state
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User     *User                   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Products []*Product              `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	Payment  *TransactionPaymentData `protobuf:"bytes,6,opt,name=payment,proto3" json:"payment,omitempty"`
	Total    float64                 `protobuf:"fixed64,7,opt,name=total,proto3" json:"total,omitempty"`
	State    TransactionState        `protobuf:"varint,5,opt,name=state,proto3,enum=synapsis.v1.TransactionState" json:"state,omitempty"`
	Dt       *DT                     `protobuf:"bytes,99,opt,name=dt,proto3" json:"dt,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Transaction) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Transaction) GetPayment() *TransactionPaymentData {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Transaction) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Transaction) GetState() TransactionState {
	if x != nil {
		return x.State
	}
	return TransactionState_TRANSACTION_STATE_UNSPECIFIED
}

func (x *Transaction) GetDt() *DT {
	if x != nil {
		return x.Dt
	}
	return nil
}

type TransactionPaymentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrisCode string                 `protobuf:"bytes,1,opt,name=qrisCode,proto3" json:"qrisCode,omitempty" bson:"qris_code"`
	Amount   float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	IssuedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=issuedAt,proto3" json:"issuedAt,omitempty" bson:"issued_at"`
	PaidAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=paidAt,proto3" json:"paidAt,omitempty" bson:"paid_at"`
}

func (x *TransactionPaymentData) Reset() {
	*x = TransactionPaymentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionPaymentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionPaymentData) ProtoMessage() {}

func (x *TransactionPaymentData) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionPaymentData.ProtoReflect.Descriptor instead.
func (*TransactionPaymentData) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionPaymentData) GetQrisCode() string {
	if x != nil {
		return x.QrisCode
	}
	return ""
}

func (x *TransactionPaymentData) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionPaymentData) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *TransactionPaymentData) GetPaidAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PaidAt
	}
	return nil
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	User      *User      `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Products  []*Product `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTransactionRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CreateTransactionRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateTransactionRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type CreateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *CreateTransactionResponse) Reset() {
	*x = CreateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionResponse) ProtoMessage() {}

func (x *CreateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type GetTransactionByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTransactionByIdRequest) Reset() {
	*x = GetTransactionByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionByIdRequest) ProtoMessage() {}

func (x *GetTransactionByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionByIdRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTransactionByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTransactionByIdResponse) Reset() {
	*x = GetTransactionByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionByIdResponse) ProtoMessage() {}

func (x *GetTransactionByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionByIdResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionByIdResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionByIdResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type GetTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     []string           `protobuf:"bytes,1,rep,name=userId,proto3" json:"userId,omitempty"`
	ProductId  []string           `protobuf:"bytes,2,rep,name=productId,proto3" json:"productId,omitempty"`
	State      []TransactionState `protobuf:"varint,3,rep,packed,name=state,proto3,enum=synapsis.v1.TransactionState" json:"state,omitempty"`
	Pagination *Pagination        `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetTransactionsRequest) Reset() {
	*x = GetTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsRequest) ProtoMessage() {}

func (x *GetTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransactionsRequest) GetUserId() []string {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *GetTransactionsRequest) GetProductId() []string {
	if x != nil {
		return x.ProductId
	}
	return nil
}

func (x *GetTransactionsRequest) GetState() []TransactionState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *GetTransactionsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Pagination   *Pagination    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetTransactionsResponse) Reset() {
	*x = GetTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse) ProtoMessage() {}

func (x *GetTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{7}
}

func (x *GetTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *GetTransactionsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UpdateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *UpdateTransactionRequest) Reset() {
	*x = UpdateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_transaction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionRequest) ProtoMessage() {}

func (x *UpdateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_transaction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionRequest.ProtoReflect.Descriptor instead.
func (*UpdateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_transaction_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_synapsis_v1_transaction_proto protoreflect.FileDescriptor

var file_synapsis_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a,
	0x02, 0x64, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x6e, 0x61,
	0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x54, 0x52, 0x02, 0x64, 0x74, 0x22, 0xb8,
	0x01, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x72, 0x69,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x72, 0x69,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x06, 0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x57, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x01,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x56, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x8c, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x1d,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x42, 0xb5, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x61, 0x66, 0x69,
	0x64, 0x70, 0x72, 0x61, 0x79, 0x6f, 0x67, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x79, 0x6e, 0x61,
	0x70, 0x73, 0x69, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synapsis_v1_transaction_proto_rawDescOnce sync.Once
	file_synapsis_v1_transaction_proto_rawDescData = file_synapsis_v1_transaction_proto_rawDesc
)

func file_synapsis_v1_transaction_proto_rawDescGZIP() []byte {
	file_synapsis_v1_transaction_proto_rawDescOnce.Do(func() {
		file_synapsis_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_synapsis_v1_transaction_proto_rawDescData)
	})
	return file_synapsis_v1_transaction_proto_rawDescData
}

var file_synapsis_v1_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synapsis_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_synapsis_v1_transaction_proto_goTypes = []interface{}{
	(TransactionState)(0),              // 0: synapsis.v1.TransactionState
	(*Transaction)(nil),                // 1: synapsis.v1.Transaction
	(*TransactionPaymentData)(nil),     // 2: synapsis.v1.TransactionPaymentData
	(*CreateTransactionRequest)(nil),   // 3: synapsis.v1.CreateTransactionRequest
	(*CreateTransactionResponse)(nil),  // 4: synapsis.v1.CreateTransactionResponse
	(*GetTransactionByIdRequest)(nil),  // 5: synapsis.v1.GetTransactionByIdRequest
	(*GetTransactionByIdResponse)(nil), // 6: synapsis.v1.GetTransactionByIdResponse
	(*GetTransactionsRequest)(nil),     // 7: synapsis.v1.GetTransactionsRequest
	(*GetTransactionsResponse)(nil),    // 8: synapsis.v1.GetTransactionsResponse
	(*UpdateTransactionRequest)(nil),   // 9: synapsis.v1.UpdateTransactionRequest
	(*User)(nil),                       // 10: synapsis.v1.User
	(*Product)(nil),                    // 11: synapsis.v1.Product
	(*DT)(nil),                         // 12: synapsis.v1.DT
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
	(*Pagination)(nil),                 // 14: synapsis.v1.Pagination
}
var file_synapsis_v1_transaction_proto_depIdxs = []int32{
	10, // 0: synapsis.v1.Transaction.user:type_name -> synapsis.v1.User
	11, // 1: synapsis.v1.Transaction.products:type_name -> synapsis.v1.Product
	2,  // 2: synapsis.v1.Transaction.payment:type_name -> synapsis.v1.TransactionPaymentData
	0,  // 3: synapsis.v1.Transaction.state:type_name -> synapsis.v1.TransactionState
	12, // 4: synapsis.v1.Transaction.dt:type_name -> synapsis.v1.DT
	13, // 5: synapsis.v1.TransactionPaymentData.issuedAt:type_name -> google.protobuf.Timestamp
	13, // 6: synapsis.v1.TransactionPaymentData.paidAt:type_name -> google.protobuf.Timestamp
	10, // 7: synapsis.v1.CreateTransactionRequest.user:type_name -> synapsis.v1.User
	11, // 8: synapsis.v1.CreateTransactionRequest.products:type_name -> synapsis.v1.Product
	1,  // 9: synapsis.v1.CreateTransactionResponse.transaction:type_name -> synapsis.v1.Transaction
	1,  // 10: synapsis.v1.GetTransactionByIdResponse.transaction:type_name -> synapsis.v1.Transaction
	0,  // 11: synapsis.v1.GetTransactionsRequest.state:type_name -> synapsis.v1.TransactionState
	14, // 12: synapsis.v1.GetTransactionsRequest.pagination:type_name -> synapsis.v1.Pagination
	1,  // 13: synapsis.v1.GetTransactionsResponse.transactions:type_name -> synapsis.v1.Transaction
	14, // 14: synapsis.v1.GetTransactionsResponse.pagination:type_name -> synapsis.v1.Pagination
	1,  // 15: synapsis.v1.UpdateTransactionRequest.transaction:type_name -> synapsis.v1.Transaction
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_synapsis_v1_transaction_proto_init() }
func file_synapsis_v1_transaction_proto_init() {
	if File_synapsis_v1_transaction_proto != nil {
		return
	}
	file_synapsis_v1_common_proto_init()
	file_synapsis_v1_product_proto_init()
	file_synapsis_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synapsis_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionPaymentData); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionResponse); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionByIdRequest); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionByIdResponse); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsRequest); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsResponse); i {
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
		file_synapsis_v1_transaction_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTransactionRequest); i {
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
			RawDescriptor: file_synapsis_v1_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synapsis_v1_transaction_proto_goTypes,
		DependencyIndexes: file_synapsis_v1_transaction_proto_depIdxs,
		EnumInfos:         file_synapsis_v1_transaction_proto_enumTypes,
		MessageInfos:      file_synapsis_v1_transaction_proto_msgTypes,
	}.Build()
	File_synapsis_v1_transaction_proto = out.File
	file_synapsis_v1_transaction_proto_rawDesc = nil
	file_synapsis_v1_transaction_proto_goTypes = nil
	file_synapsis_v1_transaction_proto_depIdxs = nil
}
