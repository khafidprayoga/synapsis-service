package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s synapsisService) CreateProduct(
	ctx context.Context,
	request *synapsisv1.CreateProductRequest,
) (*synapsisv1.CreateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetProductById(
	ctx context.Context,
	request *synapsisv1.GetProductByIdRequest,
) (*synapsisv1.GetProductByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) UpdateProduct(
	ctx context.Context,
	request *synapsisv1.UpdateProductRequest,
) (*synapsisv1.GetProductByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) DeleteProduct(
	ctx context.Context,
	request *synapsisv1.DeleteProductRequest,
) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetProducts(
	ctx context.Context,
	request *synapsisv1.GetProductsRequest,
) (*synapsisv1.GetProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}
