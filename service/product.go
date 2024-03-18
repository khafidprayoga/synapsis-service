package service

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/samber/lo"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s synapsisService) CreateProduct(
	ctx context.Context,
	request *synapsisv1.CreateProductRequest,
) (*synapsisv1.CreateProductResponse, error) {
	{
		st := status.New(codes.InvalidArgument, "invalid request for create product data")

		nameProductErr := validation.Validate(
			request.GetName(),
			validation.Required,
			validation.Length(3, 100),
		)
		if nameProductErr != nil {
			formatted, _ := st.WithDetails(&epb.ErrorInfo{
				Reason: nameProductErr.Error(),
				Domain: "product.name",
			})
			return nil, formatted.Err()
		}

		if request.GetPrice() <= 0 {
			formatted, _ := st.WithDetails(&epb.ErrorInfo{
				Reason: "price must be greater than 0",
				Domain: "product.price",
			})
			return nil, formatted.Err()
		}

		productIdErr := validation.Validate(
			request.GetProductCategoryIds(),
			validation.Required,
			validation.Each(validation.Required),
		)

		if productIdErr != nil {
			formatted, _ := st.WithDetails(&epb.ErrorInfo{
				Reason: productIdErr.Error(),
				Domain: "product.productCategoryIds",
			})
			return nil, formatted.Err()
		}
		request.ProductCategoryIds = lo.Uniq(request.GetProductCategoryIds())

		if request.GetImageURL() == "" {
			request.ImageURL = "https://picsum.photos/seed/picsum/854/480"
		}
	}

	productCat, errGetProductCat := s.
		productRepo.
		GetProductCategoryById(ctx, request.GetProductCategoryIds()...)

	if errGetProductCat != nil {
		errDetails := &epb.ErrorInfo{
			Reason:   errGetProductCat.Error(),
			Domain:   "product.productCategoryIds",
			Metadata: nil,
		}
		st := status.New(codes.Internal, "error get product category data")
		formatted, _ := st.WithDetails(errDetails)
		return nil, formatted.Err()
	}

	if len(productCat) == 0 {
		return nil, status.Errorf(codes.NotFound, "contains invalid product category (empty)")
	}
	product := &synapsisv1.Product{
		Name:              request.GetName(),
		Description:       request.GetDescription(),
		Price:             request.GetPrice(),
		ImageURL:          request.GetImageURL(),
		ProductCategories: productCat,
	}

	productData, errCreateProduct := s.productRepo.CreateProduct(ctx, product)
	if errCreateProduct != nil {
		errDetails := &epb.ErrorInfo{
			Reason:   errCreateProduct.Error(),
			Domain:   "product",
			Metadata: nil,
		}
		st := status.New(codes.Internal, "error create product data")
		formatted, _ := st.WithDetails(errDetails)
		return nil, formatted.Err()
	}

	return productData, nil
}

func (s synapsisService) GetProductById(
	ctx context.Context,
	request *synapsisv1.GetProductByIdRequest,
) (*synapsisv1.GetProductByIdResponse, error) {
	if validateProductIdErr := validation.Validate(
		request.GetId(),
		validation.Required); validateProductIdErr != nil {
		return nil, status.Error(codes.InvalidArgument, validateProductIdErr.Error())
	}

	productData, errGetProduct := s.productRepo.GetProductById(ctx, request.GetId())
	if errGetProduct != nil {
		st := status.New(codes.Internal, "error get product data by id")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason: errGetProduct.Error(),
			Domain: "product",
		})
		return nil, formatted.Err()
	}

	return &synapsisv1.GetProductByIdResponse{
		Product: productData,
	}, nil
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
