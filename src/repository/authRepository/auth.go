package authRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
)

func (r redisRepository) Login(
	_ context.Context,
	_ *synapsisv1.LoginRequest,
) (*synapsisv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r redisRepository) Logout(
	_ context.Context,
	_ *synapsisv1.LogoutRequest,
) error {
	//TODO implement me
	panic("implement me")
}
