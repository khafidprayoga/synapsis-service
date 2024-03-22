package service

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/khafidprayoga/synapsis-service/src/common/helper"
	"github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s synapsisService) CreateProductCategory(
	ctx context.Context,
	request *synapsisv1.CreateProductCategoryRequest,
) (*synapsisv1.CreateProductCategoryResponse, error) {
	productCategory := []*synapsisv1.ProductCategory{}
	{
		st := status.New(codes.InvalidArgument, "invalid request for create product category data")
		if len(request.GetData()) == 0 {
			formatted, _ := st.WithDetails(&epb.ErrorInfo{
				Reason: "data must not be empty",
				Domain: "productCategory",
			})
			return nil, formatted.Err()
		}

		for index, d := range request.GetData() {
			nameProductCatErr := validation.Validate(
				d.GetName(),
				validation.Required,
			)

			if nameProductCatErr != nil {
				formatted, _ := st.WithDetails(&epb.ErrorInfo{
					Reason: nameProductCatErr.Error(),
					Domain: "productCategory.name",
					Metadata: map[string]string{
						"arrayIndex": fmt.Sprintf("%v", index),
					},
				})
				return nil, formatted.Err()
			}

			productCategory = append(productCategory, &synapsisv1.ProductCategory{
				Name:        d.GetName(),
				Description: d.GetDescription(),
			})
		}
	}

	createdCategory, errCreate := s.productRepo.CreateProductCategory(ctx, productCategory)
	if errCreate != nil {
		st := status.New(codes.Internal, "error create product category data")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason: errCreate.Error(),
			Domain: "productCategory",
		})

		return nil, formatted.Err()
	}

	return &synapsisv1.CreateProductCategoryResponse{
		ProductCategory: createdCategory,
	}, nil
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
	allowedOrderBy := []string{"name"}
	metadataErr := map[string]string{
		"allowedOrderBy": "name",
	}

	st := status.New(codes.InvalidArgument, "invalid pagination request for get product category data")
	invalidPaginationReqErr := commonHelper.ValidatePaginationRequest(request.GetPagination(), allowedOrderBy...)
	if invalidPaginationReqErr != nil {
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason:   invalidPaginationReqErr.Error(),
			Domain:   "productCategory",
			Metadata: metadataErr,
		})
		return nil, formatted.Err()
	}

	count, categories, errGetCategory := s.productRepo.GetProductCategories(ctx, request.GetPagination())
	if errGetCategory != nil {
		st = status.New(codes.Internal, "error on get product category data")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason:   errGetCategory.Error(),
			Domain:   "productCategory",
			Metadata: metadataErr,
		})

		return nil, formatted.Err()
	}

	request.Pagination.Count = count
	return &synapsisv1.GetProductCategoriesResponse{
		ProductCategories: categories,
		Pagination:        request.GetPagination(),
	}, nil
}
