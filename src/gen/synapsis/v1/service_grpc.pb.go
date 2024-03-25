// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: synapsis/v1/service.proto

package synapsisv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SynapsisService_Ping_FullMethodName                  = "/synapsis.v1.SynapsisService/Ping"
	SynapsisService_CreateUser_FullMethodName            = "/synapsis.v1.SynapsisService/CreateUser"
	SynapsisService_GetUserById_FullMethodName           = "/synapsis.v1.SynapsisService/GetUserById"
	SynapsisService_Login_FullMethodName                 = "/synapsis.v1.SynapsisService/Login"
	SynapsisService_Logout_FullMethodName                = "/synapsis.v1.SynapsisService/Logout"
	SynapsisService_ForgotPassword_FullMethodName        = "/synapsis.v1.SynapsisService/ForgotPassword"
	SynapsisService_ResetPassword_FullMethodName         = "/synapsis.v1.SynapsisService/ResetPassword"
	SynapsisService_CreateProductCategory_FullMethodName = "/synapsis.v1.SynapsisService/CreateProductCategory"
	SynapsisService_GetProductCategory_FullMethodName    = "/synapsis.v1.SynapsisService/GetProductCategory"
	SynapsisService_UpdateProductCategory_FullMethodName = "/synapsis.v1.SynapsisService/UpdateProductCategory"
	SynapsisService_DeleteProductCategory_FullMethodName = "/synapsis.v1.SynapsisService/DeleteProductCategory"
	SynapsisService_GetProductCategories_FullMethodName  = "/synapsis.v1.SynapsisService/GetProductCategories"
	SynapsisService_CreateProduct_FullMethodName         = "/synapsis.v1.SynapsisService/CreateProduct"
	SynapsisService_GetProductById_FullMethodName        = "/synapsis.v1.SynapsisService/GetProductById"
	SynapsisService_UpdateProduct_FullMethodName         = "/synapsis.v1.SynapsisService/UpdateProduct"
	SynapsisService_DeleteProduct_FullMethodName         = "/synapsis.v1.SynapsisService/DeleteProduct"
	SynapsisService_GetProducts_FullMethodName           = "/synapsis.v1.SynapsisService/GetProducts"
	SynapsisService_CreateTransaction_FullMethodName     = "/synapsis.v1.SynapsisService/CreateTransaction"
	SynapsisService_GetTransactionById_FullMethodName    = "/synapsis.v1.SynapsisService/GetTransactionById"
	SynapsisService_UpdateTransaction_FullMethodName     = "/synapsis.v1.SynapsisService/UpdateTransaction"
	SynapsisService_GetTransactions_FullMethodName       = "/synapsis.v1.SynapsisService/GetTransactions"
)

// SynapsisServiceClient is the client API for SynapsisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SynapsisServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	// User
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	// Auth
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Product Category
	CreateProductCategory(ctx context.Context, in *CreateProductCategoryRequest, opts ...grpc.CallOption) (*CreateProductCategoryResponse, error)
	GetProductCategory(ctx context.Context, in *GetProductCategoryRequest, opts ...grpc.CallOption) (*GetProductCategoryResponse, error)
	UpdateProductCategory(ctx context.Context, in *UpdateProductCategoryRequest, opts ...grpc.CallOption) (*GetProductCategoryResponse, error)
	DeleteProductCategory(ctx context.Context, in *DeleteProductCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProductCategories(ctx context.Context, in *GetProductCategoriesRequest, opts ...grpc.CallOption) (*GetProductCategoriesResponse, error)
	// Product
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	// Transaction
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetTransactionById(ctx context.Context, in *GetTransactionByIdRequest, opts ...grpc.CallOption) (*GetTransactionByIdResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*GetTransactionByIdResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
}

type synapsisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSynapsisServiceClient(cc grpc.ClientConnInterface) SynapsisServiceClient {
	return &synapsisServiceClient{cc}
}

func (c *synapsisServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, SynapsisService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, SynapsisService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SynapsisService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SynapsisService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, SynapsisService_ForgotPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SynapsisService_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) CreateProductCategory(ctx context.Context, in *CreateProductCategoryRequest, opts ...grpc.CallOption) (*CreateProductCategoryResponse, error) {
	out := new(CreateProductCategoryResponse)
	err := c.cc.Invoke(ctx, SynapsisService_CreateProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetProductCategory(ctx context.Context, in *GetProductCategoryRequest, opts ...grpc.CallOption) (*GetProductCategoryResponse, error) {
	out := new(GetProductCategoryResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) UpdateProductCategory(ctx context.Context, in *UpdateProductCategoryRequest, opts ...grpc.CallOption) (*GetProductCategoryResponse, error) {
	out := new(GetProductCategoryResponse)
	err := c.cc.Invoke(ctx, SynapsisService_UpdateProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) DeleteProductCategory(ctx context.Context, in *DeleteProductCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SynapsisService_DeleteProductCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetProductCategories(ctx context.Context, in *GetProductCategoriesRequest, opts ...grpc.CallOption) (*GetProductCategoriesResponse, error) {
	out := new(GetProductCategoriesResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetProductCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, SynapsisService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error) {
	out := new(GetProductByIdResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetProductById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error) {
	out := new(GetProductByIdResponse)
	err := c.cc.Invoke(ctx, SynapsisService_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SynapsisService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, SynapsisService_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetTransactionById(ctx context.Context, in *GetTransactionByIdRequest, opts ...grpc.CallOption) (*GetTransactionByIdResponse, error) {
	out := new(GetTransactionByIdResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetTransactionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*GetTransactionByIdResponse, error) {
	out := new(GetTransactionByIdResponse)
	err := c.cc.Invoke(ctx, SynapsisService_UpdateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synapsisServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, SynapsisService_GetTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SynapsisServiceServer is the server API for SynapsisService service.
// All implementations must embed UnimplementedSynapsisServiceServer
// for forward compatibility
type SynapsisServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	// User
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	// Auth
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error)
	// Product Category
	CreateProductCategory(context.Context, *CreateProductCategoryRequest) (*CreateProductCategoryResponse, error)
	GetProductCategory(context.Context, *GetProductCategoryRequest) (*GetProductCategoryResponse, error)
	UpdateProductCategory(context.Context, *UpdateProductCategoryRequest) (*GetProductCategoryResponse, error)
	DeleteProductCategory(context.Context, *DeleteProductCategoryRequest) (*emptypb.Empty, error)
	GetProductCategories(context.Context, *GetProductCategoriesRequest) (*GetProductCategoriesResponse, error)
	// Product
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	GetProductById(context.Context, *GetProductByIdRequest) (*GetProductByIdResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*GetProductByIdResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	// Transaction
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetTransactionById(context.Context, *GetTransactionByIdRequest) (*GetTransactionByIdResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*GetTransactionByIdResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	mustEmbedUnimplementedSynapsisServiceServer()
}

// UnimplementedSynapsisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSynapsisServiceServer struct {
}

func (UnimplementedSynapsisServiceServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSynapsisServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSynapsisServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedSynapsisServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSynapsisServiceServer) Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSynapsisServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedSynapsisServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedSynapsisServiceServer) CreateProductCategory(context.Context, *CreateProductCategoryRequest) (*CreateProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductCategory not implemented")
}
func (UnimplementedSynapsisServiceServer) GetProductCategory(context.Context, *GetProductCategoryRequest) (*GetProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCategory not implemented")
}
func (UnimplementedSynapsisServiceServer) UpdateProductCategory(context.Context, *UpdateProductCategoryRequest) (*GetProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductCategory not implemented")
}
func (UnimplementedSynapsisServiceServer) DeleteProductCategory(context.Context, *DeleteProductCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductCategory not implemented")
}
func (UnimplementedSynapsisServiceServer) GetProductCategories(context.Context, *GetProductCategoriesRequest) (*GetProductCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCategories not implemented")
}
func (UnimplementedSynapsisServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedSynapsisServiceServer) GetProductById(context.Context, *GetProductByIdRequest) (*GetProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedSynapsisServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*GetProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedSynapsisServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedSynapsisServiceServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedSynapsisServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedSynapsisServiceServer) GetTransactionById(context.Context, *GetTransactionByIdRequest) (*GetTransactionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionById not implemented")
}
func (UnimplementedSynapsisServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*GetTransactionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedSynapsisServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedSynapsisServiceServer) mustEmbedUnimplementedSynapsisServiceServer() {}

// UnsafeSynapsisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SynapsisServiceServer will
// result in compilation errors.
type UnsafeSynapsisServiceServer interface {
	mustEmbedUnimplementedSynapsisServiceServer()
}

func RegisterSynapsisServiceServer(s grpc.ServiceRegistrar, srv SynapsisServiceServer) {
	s.RegisterService(&SynapsisService_ServiceDesc, srv)
}

func _SynapsisService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_ForgotPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_CreateProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).CreateProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_CreateProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).CreateProductCategory(ctx, req.(*CreateProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetProductCategory(ctx, req.(*GetProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_UpdateProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).UpdateProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_UpdateProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).UpdateProductCategory(ctx, req.(*UpdateProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_DeleteProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).DeleteProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_DeleteProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).DeleteProductCategory(ctx, req.(*DeleteProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetProductCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetProductCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetProductCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetProductCategories(ctx, req.(*GetProductCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetProductById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetProductById(ctx, req.(*GetProductByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetTransactionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetTransactionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetTransactionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetTransactionById(ctx, req.(*GetTransactionByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_UpdateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynapsisService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynapsisServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SynapsisService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynapsisServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SynapsisService_ServiceDesc is the grpc.ServiceDesc for SynapsisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SynapsisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "synapsis.v1.SynapsisService",
	HandlerType: (*SynapsisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SynapsisService_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _SynapsisService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _SynapsisService_GetUserById_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SynapsisService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SynapsisService_Logout_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _SynapsisService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _SynapsisService_ResetPassword_Handler,
		},
		{
			MethodName: "CreateProductCategory",
			Handler:    _SynapsisService_CreateProductCategory_Handler,
		},
		{
			MethodName: "GetProductCategory",
			Handler:    _SynapsisService_GetProductCategory_Handler,
		},
		{
			MethodName: "UpdateProductCategory",
			Handler:    _SynapsisService_UpdateProductCategory_Handler,
		},
		{
			MethodName: "DeleteProductCategory",
			Handler:    _SynapsisService_DeleteProductCategory_Handler,
		},
		{
			MethodName: "GetProductCategories",
			Handler:    _SynapsisService_GetProductCategories_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _SynapsisService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _SynapsisService_GetProductById_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _SynapsisService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _SynapsisService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _SynapsisService_GetProducts_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _SynapsisService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransactionById",
			Handler:    _SynapsisService_GetTransactionById_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _SynapsisService_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _SynapsisService_GetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "synapsis/v1/service.proto",
}