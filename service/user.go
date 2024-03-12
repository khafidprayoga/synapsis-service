package service

import (
	"context"
	"errors"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func (svc synapsisService) CreateUser(
	ctx context.Context,
	request *synapsisv1.CreateUserRequest) (*synapsisv1.CreateUserResponse, error) {
	// validations
	{
		if request.GetEmail() == "" {
			return nil, errors.New("email is required")
		}

		if request.GetFullName() == "" {
			return nil, errors.New("full name is required")
		}

		if request.GetPassword() == "" {
			return nil, errors.New("password is required")
		}

		if request.GetDob() == nil {
			return nil, errors.New("dob is required")
		}
	}

	return svc.Repo.CreateUser(ctx, request)
}
