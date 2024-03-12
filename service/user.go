package service

import (
	"context"
	"errors"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func (s synapsisService) CreateUser(
	ctx context.Context,
	request *synapsisv1.CreateUserRequest) (*synapsisv1.CreateUserResponse, error) {
	if request.GetEmail() == "" {
		return nil, errors.New("email is required")
	}

	return &synapsisv1.CreateUserResponse{User: &synapsisv1.User{
		Id:    "asd",
		Email: request.GetEmail(),
	}}, nil
}
