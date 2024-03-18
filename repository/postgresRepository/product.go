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

// todo: optimize with join to category and relations only call db with 1 query
func (p postgresRepository) GetProductById(
	_ context.Context,
	productId string) (*synapsisv1.Product, error) {

	productData := &synapsisv1.Product{}
	findProductById := p.orm.
		Where("product.id = ?", productId).
		//Where("product.deleted_at IS NULL").
		Find(&productData).Debug()

	if findProductById.Error != nil {
		return nil, unwrapError(findProductById)
	}

	if productData.GetDeletedAt() > 0 {
		return nil, gorm.ErrRecordNotFound
	}

	relations := []*synapsisv1.ProductCategoryRelation{}
	findRelations := p.orm.
		Where("product_id = ?", productData.GetId()).
		Find(&relations).Debug()
	if findRelations.Error != nil {
		return nil, unwrapError(findRelations)
	}

	//  get the category id
	categoryId := lo.Map(relations,
		func(relation *synapsisv1.ProductCategoryRelation, _ int) string {
			return relation.GetProductCategoryId()
		})

	categories := []*synapsisv1.ProductCategory{}

	findCategories := p.orm.
		Where("id IN (?)", categoryId).
		Find(&categories)
	if findCategories.Error != nil {
		return nil, unwrapError(findCategories)
	}

	productData.ProductCategories = categories
	return productData, nil
}
