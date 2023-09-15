package services

import "app/internal/core/ports"

type blockQueService struct {
	blockQueRepo ports.BlockQueRepositories
}

func NewBlockQueService(blockQueRepo ports.BlockQueRepositories)ports.BlockQueService {
	return blockQueService{blockQueRepo: blockQueRepo}
}

func(b blockQueService) Produce(hex string) error {
	err := b.blockQueRepo.Produce(hex)
	if err != nil {
		return err
	}

	return nil
}