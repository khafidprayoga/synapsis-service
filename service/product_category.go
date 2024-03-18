package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s synapsisService) CreateProductCategory(
	ctx context.Context,
	request *synapsisv1.CreateProductCategoryRequest,
) (*synapsisv1.CreateProductCategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetProductCategory(
	ctx context.Context,
	request *synapsisv1.GetProductCategoryRequest,
) (*synapsisv1.GetProductCategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) UpdateProductCategory(
	ctx context.Context,
	request *synapsisv1.UpdateProductCategoryRequest,
) (*synapsisv1.GetProductCategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) DeleteProductCategory(
	ctx context.Context,
	request *synapsisv1.DeleteProductCategoryRequest,
) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetProductCategories(
	ctx context.Context,
	request *synapsisv1.GetProductCategoriesRequest,
) (*synapsisv1.GetProductCategoriesResponse, error) {
	//TODO implement me
	panic("implement me")
}
