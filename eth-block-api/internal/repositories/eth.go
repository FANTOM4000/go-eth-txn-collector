package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethRepositories struct {
	ethClient *ethclient.Client
}

func NewEthRepositories(ethClient *ethclient.Client) ports.EthRepositories {
	return ethRepositories{ethClient: ethClient}
}

func (e ethRepositories) GetTransactionByBlockHash(ctx context.Context, hex string) ([]domains.Transaction, error) {
	block, err := e.ethClient.BlockByHash(ctx, common.HexToHash(hex))
	if err != nil {
		return []domains.Transaction{}, err
	}

	transactions := []domains.Transaction{}

	for _, tx := range block.Transactions() {
		from, err := e.GetFrom(tx)
		if err == nil {
			fmt.Println(from)
		}
		receipt, err := e.ethClient.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			logger.Error("error get TransactionReceipt", logger.ErrField(err))
			return []domains.Transaction{}, err
		}
		transactions = append(transactions, domains.Transaction{
			Hex:           tx.Hash().Hex(),
			Value:         tx.Value().Uint64(),
			Gas:           tx.Gas(),
			GasPrice:      tx.GasPrice().Uint64(),
			Nonce:         tx.Nonce(),
			Reciever:      tx.To().Hex(),
			Sender:        from,
			ReceiptStatus: receipt.Status,
		})
	}
	return transactions, nil
}

func (e ethRepositories) GetFrom(tx *types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from.String(), err
}
