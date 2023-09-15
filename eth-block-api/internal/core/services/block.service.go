package services

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"app/pkg/logger"
	"app/pkg/standard"
	"context"
)

type blockService struct {
	ethRepo            ports.EthRepositories
	transactionQueRepo ports.TransactionQueRepositories
}

func NewBlockService(ethRepo ports.EthRepositories, transactionQueRepo ports.TransactionQueRepositories) ports.BlockService {
	return blockService{ethRepo: ethRepo, transactionQueRepo: transactionQueRepo}
}

func (b blockService) ProduceTrasactionFromBlockHash(ctx context.Context, number uint64) (dto.BaseOKResponseWithData[[]domains.Transaction], error) {
	txns, err := b.ethRepo.GetTransactionByBlockHash(ctx, number)
	if err != nil {
		logger.Error("get txns from eth client error", logger.ErrField(err))
		return dto.BaseOKResponseWithData[[]domains.Transaction]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.CreateError,
				Message: "get txns from eth client error",
			},
		}, err
	}

	for _, txn := range txns {
		err = b.transactionQueRepo.Produce(txn)
		if err != nil {
			logger.Error("produce to que error", logger.ErrField(err))
			return dto.BaseOKResponseWithData[[]domains.Transaction]{
				BaseOKResponse: dto.BaseOKResponse{
					Code:    standard.CreateError,
					Message: "produce to que error",
				},
			}, err
		}
	}

	return dto.BaseOKResponseWithData[[]domains.Transaction]{
		BaseOKResponse: dto.BaseOKResponse{
			Code:    standard.SuccessCode,
			Message: "success",
		},
		Data: txns,
	}, err
}
