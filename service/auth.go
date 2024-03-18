package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s synapsisService) Login(
	ctx context.Context,
	request *synapsisv1.LoginRequest,
) (*synapsisv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) Logout(
	ctx context.Context,
	request *synapsisv1.LogoutRequest,
) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) ForgotPassword(
	ctx context.Context,
	request *synapsisv1.ForgotPasswordRequest,
) (*synapsisv1.ForgotPasswordResponse, error) {

	//TODO implement me
	panic("implement me")
}

func (s synapsisService) ResetPassword(
	ctx context.Context,
	request *synapsisv1.ResetPasswordRequest,
) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
