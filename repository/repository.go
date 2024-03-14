package repository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

type userRepository interface {
	CreateUser(context.Context, *synapsisv1.CreateUserRequest) (*synapsisv1.CreateUserResponse, error)
	GetUserByEmail(_ context.Context, email string) (*synapsisv1.User, error)
}
type SynapsisRepository interface {
	userRepository
}
