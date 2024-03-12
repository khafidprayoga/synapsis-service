package mongoRepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s mongoRepository) CreateUser(
	ctx context.Context,
	request *synapsisv1.CreateUserRequest,
) (*synapsisv1.CreateUserResponse, error) {
	hashedPass, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return nil, errors.New("error hashing password")
	}

	at := timestamppb.Now()
	return &synapsisv1.CreateUserResponse{User: &synapsisv1.User{
		Id:        uuid.New().String(),
		FullName:  request.GetFullName(),
		Email:     request.GetEmail(),
		Password:  string(hashedPass),
		Dob:       request.GetDob(),
		CreatedAt: at,
		UpdatedAt: at,
	}}, nil
}
