package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func (s synapsisService) CreateTransaction(
	ctx context.Context,
	request *synapsisv1.CreateTransactionRequest,
) (*synapsisv1.CreateTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetTransactionById(
	ctx context.Context,
	request *synapsisv1.GetTransactionByIdRequest,
) (*synapsisv1.GetTransactionByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s synapsisService) UpdateTransaction(
	ctx context.Context,
	request *synapsisv1.UpdateTransactionRequest,
) (*synapsisv1.GetTransactionByIdResponse, error) {

	//TODO implement me
	panic("implement me")
}

func (s synapsisService) GetTransactions(
	ctx context.Context,
	request *synapsisv1.GetTransactionsRequest,
) (*synapsisv1.GetTransactionsResponse, error) {
	//TODO implement me
	panic("implement me")
}
