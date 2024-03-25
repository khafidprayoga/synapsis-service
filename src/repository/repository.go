package repository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
)

// UserRepository
// it use mongodb as data store
type UserRepository interface {
	CreateUser(_ context.Context, _ *synapsisv1.CreateUserRequest, isSeeder bool) (*synapsisv1.CreateUserResponse, error)
	GetUserByEmail(_ context.Context, email string) (*synapsisv1.User, error)
	GetUserById(_ context.Context, userId string) (*synapsisv1.User, error)
	DeleteUserById(_ context.Context, userId string) error
	//UpdateUser(_ context.Context, user *synapsisv1.User) (*synapsisv1.User, error)
}

// ProductRepository
// it use postgres as data store
type ProductRepository interface {
	CreateProductCategory(_ context.Context, _ []*synapsisv1.ProductCategory) ([]*synapsisv1.ProductCategory, error)
	GetProductCategoryById(_ context.Context, categoryId ...string) ([]*synapsisv1.ProductCategory, error)
	GetProductCategories(_ context.Context, _ *synapsisv1.Pagination) (int64, []*synapsisv1.ProductCategory, error)
	//DeleteProductCategoryById(_ context.Context, categoryId string) error
	//UpdateProductCategory(_ context.Context, category *synapsisv1.ProductCategory) (*synapsisv1.ProductCategory, error)

	CreateProduct(_ context.Context, _ *synapsisv1.Product) (*synapsisv1.CreateProductResponse, error)
	GetProductById(_ context.Context, productId string) (*synapsisv1.Product, error)
	GetProducts(_ context.Context, _ *synapsisv1.Pagination) (int64, []*synapsisv1.Product, error)
	DeleteProductById(_ context.Context, productId string) error
	UpdateProduct(_ context.Context, product *synapsisv1.Product) (*synapsisv1.Product, error)

	GetProductRelations(_ context.Context, productId string) ([]*synapsisv1.ProductCategoryRelation, error)
}

// AuthRepository
// it use Auth0 from okta to issue token and manage the session
type AuthRepository interface {
	Login(_ context.Context, _ *synapsisv1.LoginRequest) (*synapsisv1.LoginResponse, error)
	Logout(_ context.Context, _ *synapsisv1.LogoutRequest) error
}

type SynapsisRepository interface {
	UserRepository
	ProductRepository
	AuthRepository
}
