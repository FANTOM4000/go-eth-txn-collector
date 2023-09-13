package service

import (
	"app/internal/core/ports"
	"app/pkg/standard"
	"errors"
)

type consumeService struct {
	ConsumeRepo ports.BlockAdaptorApiRepositories
}

func NewConsumeService(ConsumeRepo ports.BlockAdaptorApiRepositories) ports.ConsumerService {
	return consumeService{ConsumeRepo: ConsumeRepo}
}

func (c consumeService) Consume(blockHash string) error {
	blockAdaptorRes,err := c.ConsumeRepo.ProduceTransaction(blockHash)
	if err!=nil {
		return err
	}
	if blockAdaptorRes.Code != standard.SuccessCode {
		return errors.New(blockAdaptorRes.Message)
	}
	return nil
}
