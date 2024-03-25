package service

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/khafidprayoga/synapsis-service/src/common/helper"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
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

	p, errGetP := s.productRepo.GetProductById(ctx, request.GetId())
	if errGetP != nil {
		return nil, status.Errorf(codes.Internal, "error get product data by id: %s", errGetP)
	}

	var product = &synapsisv1.Product{
		Id:                request.GetId(),
		Name:              request.GetName(),
		Description:       request.GetDescription(),
		Price:             request.GetPrice(),
		ImageURL:          request.GetImageURL(),
		ProductCategories: productCat,
		CreatedAt:         p.GetCreatedAt(),
	}

	productData, errCreateProduct := s.productRepo.UpdateProduct(ctx, product)
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

	return &synapsisv1.GetProductByIdResponse{
		Product: productData}, nil

}

func (s synapsisService) DeleteProduct(
	ctx context.Context,
	request *synapsisv1.DeleteProductRequest,
) (*emptypb.Empty, error) {
	if validateProductIdErr := validation.Validate(
		request.GetId(),
		validation.Required); validateProductIdErr != nil {
		return nil, status.Error(codes.InvalidArgument, validateProductIdErr.Error())
	}

	_, errGetProduct := s.productRepo.GetProductById(ctx, request.GetId())
	if errGetProduct != nil {
		st := status.New(codes.Internal, "error get product data by id")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason: errGetProduct.Error(),
			Domain: "product",
		})
		return nil, formatted.Err()
	}

	errDelete := s.productRepo.DeleteProductById(ctx, request.GetId())
	return &emptypb.Empty{}, errDelete
}

func (s synapsisService) GetProducts(
	ctx context.Context,
	request *synapsisv1.GetProductsRequest,
) (*synapsisv1.GetProductsResponse, error) {
	allowedOrderBy := []string{"name"}
	metadataErr := map[string]string{
		"allowedOrderBy": "name",
	}

	st := status.New(codes.InvalidArgument, "invalid pagination request for get products data")
	invalidPaginationReqErr := commonHelper.ValidatePaginationRequest(request.GetPagination(), allowedOrderBy...)
	if invalidPaginationReqErr != nil {
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason:   invalidPaginationReqErr.Error(),
			Domain:   "products",
			Metadata: metadataErr,
		})
		return nil, formatted.Err()
	}

	count, products, errGetCategory := s.productRepo.GetProducts(ctx, request.GetPagination())
	if errGetCategory != nil {
		st = status.New(codes.Internal, "error on get products data")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason:   errGetCategory.Error(),
			Domain:   "products",
			Metadata: metadataErr,
		})

		return nil, formatted.Err()
	}

	_, categories, errGetCategoies := s.productRepo.GetProductCategories(ctx, &synapsisv1.Pagination{
		Page: &synapsisv1.Pagination_Page{
			Limit:  100,
			Offset: 0,
		},
	})

	if errGetCategory != nil {
		st = status.New(codes.Internal, "error on get product categories data")
		formatted, _ := st.WithDetails(&epb.ErrorInfo{
			Reason:   errGetCategoies.Error(),
			Domain:   "productCategory",
			Metadata: metadataErr,
		})

		return nil, formatted.Err()
	}

	for _, p := range products {
		relations, errGetRelation := s.productRepo.GetProductRelations(ctx, p.GetId())
		if errGetRelation != nil {
			st = status.New(codes.Internal, "error on get product relation data")
			formatted, _ := st.WithDetails(&epb.ErrorInfo{
				Reason:   errGetRelation.Error(),
				Domain:   "productRelation",
				Metadata: metadataErr,
			})
			return nil, formatted.Err()
		}

		categoryId := lo.Map(relations,
			func(relation *synapsisv1.ProductCategoryRelation, _ int) string {
				return relation.GetProductCategoryId()
			})

		for _, rel := range categoryId {
			for _, cat := range categories {
				if cat.GetId() == rel {
					if p.ProductCategories != nil {
						p.ProductCategories = append(p.ProductCategories, cat)
						continue
					}
					p.ProductCategories = []*synapsisv1.ProductCategory{cat}
				}
			}

		}
	}

	request.Pagination.Count = count
	return &synapsisv1.GetProductsResponse{
		Products:   products,
		Pagination: request.GetPagination(),
	}, nil
}
