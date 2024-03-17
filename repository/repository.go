package repository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

type userRepository interface {
	CreateUser(_ context.Context, _ *synapsisv1.CreateUserRequest, isSeeder bool) (*synapsisv1.CreateUserResponse, error)
	GetUserByEmail(_ context.Context, email string) (*synapsisv1.User, error)
	GetUserById(_ context.Context, userId string) (*synapsisv1.User, error)
	DeleteUserById(_ context.Context, userId string) error
	UpdateUser(_ context.Context, user *synapsisv1.User) (*synapsisv1.User, error)
}

type SynapsisRepository interface {
	userRepository
}
