package postgresRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/samber/lo"
	"gorm.io/gorm"
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
