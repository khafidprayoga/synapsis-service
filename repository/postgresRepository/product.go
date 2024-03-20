package postgresRepository

import (
	"context"
	"fmt"
	commonHelper "github.com/khafidprayoga/synapsis-service/common/helper"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

func (p postgresRepository) CreateProduct(
	_ context.Context,
	productData *synapsisv1.Product,
) (*synapsisv1.CreateProductResponse, error) {
	errDoTX := p.orm.Transaction(func(tx *gorm.DB) error {
		errInsertProduct := tx.Create(&productData)
		if errInsertProduct.Error != nil {
			return errInsertProduct.Error
		}

		var productCatRel []*synapsisv1.ProductCategoryRelation
		for _, category := range productData.GetProductCategories() {
			productCatRel = append(productCatRel, &synapsisv1.ProductCategoryRelation{
				ProductId:         productData.GetId(),
				ProductCategoryId: category.GetId(),
			})
		}

		insertRelation := tx.Create(&productCatRel)
		if insertRelation.Error != nil {
			return insertRelation.Error
		}
		return nil
	})

	if errDoTX != nil {
		return nil, errDoTX
	}

	return &synapsisv1.CreateProductResponse{
		Product: productData,
	}, nil
}

func (p postgresRepository) GetProductById(
	ctx context.Context,
	productId string) (*synapsisv1.Product, error) {

	productData := &synapsisv1.Product{}
	findProductById := p.orm.
		WithContext(ctx).
		Where("product.id = ?", productId).
		//Where("product.deleted_at IS NULL").
		Find(&productData).Debug()

	if findProductById.Error != nil {
		return nil, unwrapError(findProductById)
	}

	if productData.GetDeletedAt() > 0 {
		return nil, gorm.ErrRecordNotFound
	}

	relations, findRelationsErr := p.GetProductRelations(ctx, productData.GetId())
	if findRelationsErr != nil {
		return nil, findRelationsErr
	}

	//  get the category id
	categoryId := lo.Map(relations,
		func(relation *synapsisv1.ProductCategoryRelation, _ int) string {
			return relation.GetProductCategoryId()
		})

	categories, findCategoriesErr := p.GetProductCategoryById(ctx, categoryId...)
	if findCategoriesErr != nil {
		return nil, findCategoriesErr
	}

	productData.ProductCategories = categories
	return productData, nil
}

func (p postgresRepository) GetProducts(
	ctx context.Context,
	paging *synapsisv1.Pagination,
) (int64, []*synapsisv1.Product, error) {
	products := []*synapsisv1.Product{}
	q := p.orm.
		WithContext(ctx).
		Table("product").
		Where("deleted_at is NULL")

	if paging.GetSearch() != "" {
		search := "%" + paging.GetSearch() + "%"
		q = q.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	if paging.GetStartAt() != nil && paging.GetEndAt() != nil {
		q = q.Where("created_at BETWEEN ? AND ?", paging.GetStartAt().AsTime(), paging.GetEndAt().AsTime())
	}

	if paging.GetOrderBy() != "" {
		q = q.Order(
			fmt.Sprintf(
				"%s %s",
				paging.GetOrderBy(),
				commonHelper.PaginationSortSQL[paging.GetSort()]),
		)
	} else {
		q = q.Order("name ASC")
	}

	var categoriesCount int64
	count := q.Count(&categoriesCount)
	if count.Error != nil {
		return 0, products, unwrapError(count)
	}

	findAll := q.
		Limit(int(paging.GetPage().GetLimit())).
		Offset(int(paging.GetPage().GetOffset())).
		Find(&products)
	if findAll.Error != nil {
		return 0, products, unwrapError(findAll)
	}

	return categoriesCount, products, nil
}

func (p postgresRepository) DeleteProductById(
	ctx context.Context,
	productId string) error {
	deleteProduct := p.orm.
		WithContext(ctx).
		Model(&synapsisv1.Product{}).
		Where("id = ?", productId).
		Update("deleted_at", time.Now().UnixMilli())
	if deleteProduct.Error != nil {
		return unwrapError(deleteProduct)
	}

	return nil
}
