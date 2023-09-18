package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethRepositories struct {
	ethClient *ethclient.Client
}

func NewEthRepositories(ethClient *ethclient.Client) ports.EthRepositories {
	return ethRepositories{ethClient: ethClient}
}

func (e ethRepositories) GetTransactionByBlockHash(ctx context.Context, number uint64) ([]domains.Transaction, error) {
	block, err := e.ethClient.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return []domains.Transaction{}, err
	}

	transactions := []domains.Transaction{}
	txn := block.Transactions()
	for _, tx := range txn {
		from, _ := e.GetFrom(tx)
		value := uint64(0)
		if tx.Value() != nil {
			value = tx.Value().Uint64()
		}
		gasPrice := uint64(0)
		if tx.GasPrice() != nil {
			gasPrice = tx.GasPrice().Uint64()
		}
		reciever := ""
		if tx.To() != nil {
			reciever = tx.To().Hex()
		}

		transactions = append(transactions, domains.Transaction{
			Hex:      strings.ToLower(tx.Hash().Hex()),
			Value:    value,
			Gas:      tx.Gas(),
			GasPrice: gasPrice,
			Nonce:    tx.Nonce(),
			Reciever: strings.ToLower(reciever),
			Sender:   strings.ToLower(from),
		})
	}
	return transactions, nil
}

func (e ethRepositories) GetFrom(tx *types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		logger.Error("error get sender")
	}
	return from.String(), err
}
