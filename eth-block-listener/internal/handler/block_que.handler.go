package handler

import (
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockQueHandler struct {
	ethereumClient  *ethclient.Client
	blockQueService ports.BlockQueService
}

func NewBlockQueHandler(ethereumClient *ethclient.Client, blockQueService ports.BlockQueService) ports.BlockQueHandler {
	return blockQueHandler{ethereumClient: ethereumClient, blockQueService: blockQueService}
}

func (b blockQueHandler) ListenNewBlock() {
	headers := make(chan *types.Header)
	sub, err := b.ethereumClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		logger.Error("error subscribe eth wss", logger.ErrField(err))
		return
	}
	for {
		select {
		case err := <-sub.Err():
			logger.Error("error listen eth wss", logger.ErrField(err))
			return
		case header := <-headers:
			logger.Info("New block", logger.StringField("hex", header.Hash().Hex()))
			err = b.blockQueService.Produce(header.Hash().Hex())
			if err != nil {
				logger.Error("error produce new block", logger.StringField("hex", header.Hash().Hex()))
			}
		}
	}
}
