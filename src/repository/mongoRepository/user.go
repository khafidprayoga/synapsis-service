package mongoRepository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/khafidprayoga/synapsis-service/src/common/lib"
	synapsisv12 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const userCollection = "user"

func (s mongoRepository) GetUserByEmail(ctx context.Context, email string) (*synapsisv12.User, error) {
	col := s.mongo.Collection(userCollection)
	match := bson.M{
		"email": email,
		"dt.deleted_at": bson.M{
			"$type": bson.TypeNull,
		},
	}

	res := col.FindOne(ctx, match)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}

		s.log.Error("error finding user", zap.Error(res.Err()))
		return nil, fmt.Errorf("error find user with email: %v", email)
	}

	u := &synapsisv12.User{}
	if e := res.Decode(u); e != nil {
		s.log.Error("error decoding user", zap.Error(e))
		return nil, fmt.Errorf("error decoding user with email: %v", email)
	}

	return u, nil
}

func (s mongoRepository) CreateUser(
	ctx context.Context,
	request *synapsisv12.CreateUserRequest,
	isSeeder bool,
) (*synapsisv12.CreateUserResponse, error) {
	hashedPass, errHash := commonLib.HashPasssword(request.Password)
	if errHash != nil {
		return nil, errors.New("error hashing password")
	}

	at := timestamppb.Now()
	col := s.mongo.Collection(userCollection)
	user := &synapsisv12.User{
		Id:       uuid.New().String(),
		FullName: request.GetFullName(),
		Email:    request.GetEmail(),
		Password: string(hashedPass),
		Dob:      request.GetDob(),
		Dt: &synapsisv12.DT{
			CreatedAt: at,
			UpdatedAt: at,
			DeletedAt: nil,
		},
	}

	// flag for seed new user data
	if isSeeder {
		user.Dt.DeletedAt = timestamppb.Now()
	}

	_, errInsert := col.InsertOne(ctx, user)
	if errInsert != nil {
		return nil, errInsert
	}

	return &synapsisv12.CreateUserResponse{
		User: user,
	}, nil
}

func (s mongoRepository) GetUserById(ctx context.Context, userId string) (*synapsisv12.User, error) {
	col := s.mongo.Collection(userCollection)
	match := bson.M{
		"id": userId,
		"dt.deleted_at": bson.M{
			"$type": bson.TypeNull,
		},
	}

	res := col.FindOne(ctx, match)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}

		s.log.Error("error finding user", zap.Error(res.Err()))
		return nil, fmt.Errorf("error find user with id: %v", userId)
	}

	u := &synapsisv12.User{}
	if e := res.Decode(u); e != nil {
		s.log.Error("error decoding user", zap.Error(e))
		return nil, fmt.Errorf("error decoding user with id: %v", userId)
	}

	return u, nil
}

func (s mongoRepository) DeleteUserById(ctx context.Context, userId string) error {
	col := s.mongo.Collection(userCollection)
	match := bson.M{
		"id": userId,
		"dt.deleted_at": bson.M{
			"$type": bson.TypeNull,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"dt.deleted_at": timestamppb.Now(),
		},
	}

	_, errDelete := col.UpdateOne(ctx, match, update)
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (s mongoRepository) UpdateUser(_ context.Context, user *synapsisv12.User) (*synapsisv12.User, error) {
	//TODO implement me
	panic("implement me")
}
