package infrastructure

import (
	"app/config"
	"app/pkg/logger"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitEthereumClient() *ethclient.Client {
	client, err := ethclient.Dial(config.Get().Ethereum.NodeUrl)
	if err != nil {
		logger.Fatal("init eth client error", logger.ErrField(err))
	}
	return client
}
