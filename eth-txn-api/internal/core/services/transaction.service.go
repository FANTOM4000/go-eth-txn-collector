package services

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"app/pkg/logger"
	"app/pkg/standard"
	"context"
)

type transactionService struct {
	transactionRepo ports.TransactionRepositories
	addressRepo     ports.AddressRepositories
}

func NewTransactionService(transactionRepo ports.TransactionRepositories, addressRepo ports.AddressRepositories) ports.TransactionService {
	return transactionService{transactionRepo: transactionRepo, addressRepo: addressRepo}
}
func (t transactionService) Save(ctx context.Context, txn domains.Transaction) (dto.BaseOKResponseWithData[domains.Transaction], error) {

	senderFound, err := t.addressRepo.CheckAddressExist(ctx, txn.Sender)
	if err != nil {
		logger.Error("check address sender error", logger.ErrField(err))
		return dto.BaseOKResponseWithData[domains.Transaction]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.GenericError,
				Message: "check address sender error",
			},
		}, err
	}
	recieverFound, err := t.addressRepo.CheckAddressExist(ctx, txn.Reciever)
	if err != nil {
		logger.Error("check address reciever error", logger.ErrField(err))
		return dto.BaseOKResponseWithData[domains.Transaction]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.GenericError,
				Message: "check address reciever error",
			},
		}, err
	}

	if senderFound || recieverFound {
		tx, err := t.transactionRepo.Save(ctx, txn)
		if err != nil {
			logger.Error("error save transaction", logger.ErrField(err))
			return dto.BaseOKResponseWithData[domains.Transaction]{
				BaseOKResponse: dto.BaseOKResponse{
					Code:    standard.CreateError,
					Message: "error save transaction",
				},
			}, err
		}
		return dto.BaseOKResponseWithData[domains.Transaction]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.SuccessCode,
				Message: "success",
			},
			Data: tx,
		}, nil
	}

	return dto.BaseOKResponseWithData[domains.Transaction]{
		BaseOKResponse: dto.BaseOKResponse{
			Code:    standard.SuccessCode,
			Message: "success",
		},
	}, nil
}
func (t transactionService) GetIncomingAndOutgoingOfAddress(ctx context.Context, addr string, page int, perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error) {
	txns, err := t.transactionRepo.GetByContainAddress(ctx, addr, page, perpage)
	if err != nil {
		logger.Error("error get tansactions", logger.ErrField(err))
		return dto.BaseOKResponseWithData[[]domains.Transaction]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.GetDataError,
				Message: "error get tansactions",
			},
		}, err
	}
	return dto.BaseOKResponseWithData[[]domains.Transaction]{
		BaseOKResponse: dto.BaseOKResponse{
			Code:    standard.SuccessCode,
			Message: "success",
		},
		Data: txns,
	}, nil

}
