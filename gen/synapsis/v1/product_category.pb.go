// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: synapsis/v1/product_category.proto

package synapsisv1

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

type ProductCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(255);UNIQUE"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" gorm:"type:text"`
	CreatedAt   int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty" gorm:"autoCreateTime:milli;INDEX" bson:"-"`
	UpdatedAt   int64  `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli" bson:"-"`
}

func (x *ProductCategory) Reset() {
	*x = ProductCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategory) ProtoMessage() {}

func (x *ProductCategory) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategory.ProtoReflect.Descriptor instead.
func (*ProductCategory) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductCategory) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProductCategory) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// ======== Service ========
type CreateProductCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateProductCategoryRequest) Reset() {
	*x = CreateProductCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductCategoryRequest) ProtoMessage() {}

func (x *CreateProductCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateProductCategoryRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateProductCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCategory *ProductCategory `protobuf:"bytes,1,opt,name=productCategory,proto3" json:"productCategory,omitempty"`
}

func (x *CreateProductCategoryResponse) Reset() {
	*x = CreateProductCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductCategoryResponse) ProtoMessage() {}

func (x *CreateProductCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateProductCategoryResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductCategoryResponse) GetProductCategory() *ProductCategory {
	if x != nil {
		return x.ProductCategory
	}
	return nil
}

type GetProductCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductCategoryRequest) Reset() {
	*x = GetProductCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategoryRequest) ProtoMessage() {}

func (x *GetProductCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetProductCategoryRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCategory *ProductCategory `protobuf:"bytes,1,opt,name=productCategory,proto3" json:"productCategory,omitempty"`
}

func (x *GetProductCategoryResponse) Reset() {
	*x = GetProductCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategoryResponse) ProtoMessage() {}

func (x *GetProductCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetProductCategoryResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductCategoryResponse) GetProductCategory() *ProductCategory {
	if x != nil {
		return x.ProductCategory
	}
	return nil
}

type UpdateProductCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateProductCategoryRequest) Reset() {
	*x = UpdateProductCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductCategoryRequest) ProtoMessage() {}

func (x *UpdateProductCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductCategoryRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProductCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteProductCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductCategoryRequest) Reset() {
	*x = DeleteProductCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductCategoryRequest) ProtoMessage() {}

func (x *DeleteProductCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductCategoryRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProductCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetProductCategoriesRequest) Reset() {
	*x = GetProductCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategoriesRequest) ProtoMessage() {}

func (x *GetProductCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetProductCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductCategoriesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetProductCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCategories []*ProductCategory `protobuf:"bytes,1,rep,name=productCategories,proto3" json:"productCategories,omitempty"`
	Pagination        *Pagination        `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetProductCategoriesResponse) Reset() {
	*x = GetProductCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_product_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategoriesResponse) ProtoMessage() {}

func (x *GetProductCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_product_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetProductCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_product_category_proto_rawDescGZIP(), []int{8}
}

func (x *GetProductCategoriesResponse) GetProductCategories() []*ProductCategory {
	if x != nil {
		return x.ProductCategories
	}
	return nil
}

func (x *GetProductCategoriesResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_synapsis_v1_product_category_proto protoreflect.FileDescriptor

var file_synapsis_v1_product_category_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x18, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x54, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x64, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x37,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xb9, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x68, 0x61, 0x66, 0x69, 0x64, 0x70, 0x72, 0x61, 0x79, 0x6f, 0x67, 0x61, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79,
	0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02,
	0x0b, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x53,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x53, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synapsis_v1_product_category_proto_rawDescOnce sync.Once
	file_synapsis_v1_product_category_proto_rawDescData = file_synapsis_v1_product_category_proto_rawDesc
)

func file_synapsis_v1_product_category_proto_rawDescGZIP() []byte {
	file_synapsis_v1_product_category_proto_rawDescOnce.Do(func() {
		file_synapsis_v1_product_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_synapsis_v1_product_category_proto_rawDescData)
	})
	return file_synapsis_v1_product_category_proto_rawDescData
}

var file_synapsis_v1_product_category_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_synapsis_v1_product_category_proto_goTypes = []interface{}{
	(*ProductCategory)(nil),               // 0: synapsis.v1.ProductCategory
	(*CreateProductCategoryRequest)(nil),  // 1: synapsis.v1.CreateProductCategoryRequest
	(*CreateProductCategoryResponse)(nil), // 2: synapsis.v1.CreateProductCategoryResponse
	(*GetProductCategoryRequest)(nil),     // 3: synapsis.v1.GetProductCategoryRequest
	(*GetProductCategoryResponse)(nil),    // 4: synapsis.v1.GetProductCategoryResponse
	(*UpdateProductCategoryRequest)(nil),  // 5: synapsis.v1.UpdateProductCategoryRequest
	(*DeleteProductCategoryRequest)(nil),  // 6: synapsis.v1.DeleteProductCategoryRequest
	(*GetProductCategoriesRequest)(nil),   // 7: synapsis.v1.GetProductCategoriesRequest
	(*GetProductCategoriesResponse)(nil),  // 8: synapsis.v1.GetProductCategoriesResponse
	(*Pagination)(nil),                    // 9: synapsis.v1.Pagination
}
var file_synapsis_v1_product_category_proto_depIdxs = []int32{
	0, // 0: synapsis.v1.CreateProductCategoryResponse.productCategory:type_name -> synapsis.v1.ProductCategory
	0, // 1: synapsis.v1.GetProductCategoryResponse.productCategory:type_name -> synapsis.v1.ProductCategory
	9, // 2: synapsis.v1.GetProductCategoriesRequest.pagination:type_name -> synapsis.v1.Pagination
	0, // 3: synapsis.v1.GetProductCategoriesResponse.productCategories:type_name -> synapsis.v1.ProductCategory
	9, // 4: synapsis.v1.GetProductCategoriesResponse.pagination:type_name -> synapsis.v1.Pagination
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_synapsis_v1_product_category_proto_init() }
func file_synapsis_v1_product_category_proto_init() {
	if File_synapsis_v1_product_category_proto != nil {
		return
	}
	file_synapsis_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synapsis_v1_product_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategory); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductCategoryRequest); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductCategoryResponse); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategoryRequest); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategoryResponse); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductCategoryRequest); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductCategoryRequest); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategoriesRequest); i {
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
		file_synapsis_v1_product_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategoriesResponse); i {
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
			RawDescriptor: file_synapsis_v1_product_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synapsis_v1_product_category_proto_goTypes,
		DependencyIndexes: file_synapsis_v1_product_category_proto_depIdxs,
		MessageInfos:      file_synapsis_v1_product_category_proto_msgTypes,
	}.Build()
	File_synapsis_v1_product_category_proto = out.File
	file_synapsis_v1_product_category_proto_rawDesc = nil
	file_synapsis_v1_product_category_proto_goTypes = nil
	file_synapsis_v1_product_category_proto_depIdxs = nil
}
