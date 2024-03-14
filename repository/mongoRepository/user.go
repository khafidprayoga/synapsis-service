package mongoRepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const userCollection = "user"

func (s mongoRepository) GetUserByEmail(_ context.Context, email string) (*synapsisv1.User, error) {
	//TODO implement me
	return nil, errors.New("not implemented")
}

func (s mongoRepository) CreateUser(
	ctx context.Context,
	request *synapsisv1.CreateUserRequest,
) (*synapsisv1.CreateUserResponse, error) {
	//validate get user by email not exist
	{
		existingUser, errGetExistingUser := s.GetUserByEmail(ctx, request.GetEmail())
		if errGetExistingUser != nil {
			return nil, errGetExistingUser
		}
		if existingUser != nil {
			return nil, errors.New("user already exist try another email")
		}
	}
	hashedPass, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return nil, errors.New("error hashing password")
	}

	at := timestamppb.Now()
	col := s.mongo.Collection(userCollection)
	user := &synapsisv1.User{
		Id:       uuid.New().String(),
		FullName: request.GetFullName(),
		Email:    request.GetEmail(),
		Password: string(hashedPass),
		Dob:      request.GetDob(),
		Dt: &synapsisv1.DT{
			CreatedAt: at,
			UpdatedAt: at,
			DeletedAt: nil,
		},
	}

	_, errInsert := col.InsertOne(ctx, user)
	if errInsert != nil {
		return nil, errInsert
	}

	return &synapsisv1.CreateUserResponse{
		User: user,
	}, nil
}
