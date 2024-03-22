package service

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	"golang.org/x/text/language"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (svc synapsisService) CreateUser(
	ctx context.Context,
	request *synapsisv1.CreateUserRequest) (_ *synapsisv1.CreateUserResponse, err error) {
	{
		fullNameErr := validation.Validate(
			request.GetFullName(),
			validation.Required,
			validation.Length(0, 50),
		)

		if fullNameErr != nil {
			err = status.Errorf(codes.InvalidArgument, "invalid full name: %v", fullNameErr)
			return
		}

		emailErr := validation.Validate(request.GetEmail(), validation.Required, is.Email)
		if emailErr != nil {
			err = status.Errorf(codes.InvalidArgument, "invalid email: %v", emailErr)
			return
		}

		passwordErr := validation.Validate(
			request.GetPassword(),
			validation.Required,
			validation.Length(0, 32),
		)

		if passwordErr != nil {
			err = status.Errorf(codes.InvalidArgument, "invalid password: %v", passwordErr)
			return
		}

		errDob := validation.Validate(request.GetDob(),
			validation.Required,
		)

		if errDob != nil {
			err = status.Errorf(codes.InvalidArgument, "invalid dob: %v", errDob)
			return
		}
		min18YO := time.Now().AddDate(-18, 0, 0)
		if request.GetDob().AsTime().After(min18YO) {
			err = status.Errorf(codes.InvalidArgument, "dob must be at least 18 years old")
			return
		}
	}

	existingUser, errGetExistingUser := svc.userRepo.GetUserByEmail(ctx, request.GetEmail())
	if errGetExistingUser != nil {
		return nil, status.Errorf(codes.Internal, errGetExistingUser.Error())
	}

	if existingUser != nil {
		st := status.New(codes.InvalidArgument, "email already taken")
		_, e := st.WithDetails(
			&epb.LocalizedMessage{
				Locale:  language.Indonesian.String(),
				Message: "email sudah terpakai",
			})

		return nil, e
	}

	return svc.userRepo.CreateUser(ctx, request, false)
}

func (s synapsisService) GetUserById(
	ctx context.Context,
	request *synapsisv1.GetUserByIdRequest,
) (*synapsisv1.GetUserByIdResponse, error) {
	{
		idReqErr := validation.Validate(request.GetId(), validation.Required)
		if idReqErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", idReqErr)
		}
	}

	userData, err := s.userRepo.GetUserById(ctx, request.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if userData == nil && err == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &synapsisv1.GetUserByIdResponse{
		User: userData,
	}, nil
}
