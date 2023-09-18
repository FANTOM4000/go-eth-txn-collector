package service

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"context"
)

type transactionService struct {
	transactionAdaptor ports.TransactionAdaptorApi
}

func NewTransactionService(transactionAdaptor ports.TransactionAdaptorApi) ports.TransactionService {
	return transactionService{transactionAdaptor: transactionAdaptor}
}

func (t transactionService) GetIncomingAndOutgoingOfAddress(ctx context.Context, addr string, page int, perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error) {
	return t.transactionAdaptor.GetIncomingAndOutgoingOfAddress(ctx,addr,page,perpage)
}
