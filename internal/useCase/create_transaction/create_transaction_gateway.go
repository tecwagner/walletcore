package createTransaction

import (
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	transactionGateway "github.com/tecwagner/walletcore-service/internal/gateway/transaction_gateway"
)

type CreateTransactionUseCase struct {
	TransactionGateway transactionGateway.ITransactionGateway
	AccountGateway accountGateway.IAccountGateway
}
