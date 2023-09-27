package transactionGateway

import "github.com/tecwagner/walletcore-service/internal/entity"

type ITransactionGateway interface {
	Create(transaction *entity.Transaction) error
}