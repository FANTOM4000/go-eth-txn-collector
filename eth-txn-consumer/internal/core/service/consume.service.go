package service

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/standard"
	"errors"
)

type consumeService struct {
	ConsumeRepo ports.TransactionAdaptorApiRepositories
}

func NewConsumeService(ConsumeRepo ports.TransactionAdaptorApiRepositories) ports.ConsumerService {
	return consumeService{ConsumeRepo: ConsumeRepo}
}

func (c consumeService) Consume(txn domains.Transaction) error {
	blockAdaptorRes, err := c.ConsumeRepo.SaveTransaction(txn)
	if err != nil {
		return err
	}
	if blockAdaptorRes.Code != standard.SuccessCode {
		return errors.New(blockAdaptorRes.Message)
	}
	return nil
}
