package repository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

type SynapsisRepository interface {
	CreateUser(context.Context, *synapsisv1.CreateUserRequest) (*synapsisv1.CreateUserResponse, error)
}
