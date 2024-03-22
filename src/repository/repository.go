package repository

import (
	"context"
	synapsisv12 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
)

// UserRepository
// it use mongodb as data store
type UserRepository interface {
	CreateUser(_ context.Context, _ *synapsisv12.CreateUserRequest, isSeeder bool) (*synapsisv12.CreateUserResponse, error)
	GetUserByEmail(_ context.Context, email string) (*synapsisv12.User, error)
	GetUserById(_ context.Context, userId string) (*synapsisv12.User, error)
	DeleteUserById(_ context.Context, userId string) error
	//UpdateUser(_ context.Context, user *synapsisv1.User) (*synapsisv1.User, error)
}

// ProductRepository
// it use postgres as data store
type ProductRepository interface {
	CreateProductCategory(_ context.Context, _ []*synapsisv12.ProductCategory) ([]*synapsisv12.ProductCategory, error)
	GetProductCategoryById(_ context.Context, categoryId ...string) ([]*synapsisv12.ProductCategory, error)
	GetProductCategories(_ context.Context, _ *synapsisv12.Pagination) (int64, []*synapsisv12.ProductCategory, error)
	//DeleteProductCategoryById(_ context.Context, categoryId string) error
	//UpdateProductCategory(_ context.Context, category *synapsisv1.ProductCategory) (*synapsisv1.ProductCategory, error)

	CreateProduct(_ context.Context, _ *synapsisv12.Product) (*synapsisv12.CreateProductResponse, error)
	GetProductById(_ context.Context, productId string) (*synapsisv12.Product, error)
	GetProducts(_ context.Context, _ *synapsisv12.Pagination) (int64, []*synapsisv12.Product, error)
	DeleteProductById(_ context.Context, productId string) error
	UpdateProduct(_ context.Context, product *synapsisv12.Product) (*synapsisv12.Product, error)

	GetProductRelations(_ context.Context, productId string) ([]*synapsisv12.ProductCategoryRelation, error)
}

// AuthRepository
// it use Auth0 from okta to issue token and manage the session
type AuthRepository interface{}

type SynapsisRepository interface {
	UserRepository
	ProductRepository
	AuthRepository
}
