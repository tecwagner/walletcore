package createTransaction

import (
	"github.com/tecwagner/walletcore-service/internal/entity"
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	transactionGateway "github.com/tecwagner/walletcore-service/internal/gateway/transaction_gateway"
)



func NewCreateTransactiontUseCase(transactionGateway transactionGateway.ITransactionGateway, accountGateway accountGateway.IAccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway: accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	// Consultando a conta do cliente pagador
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	// Consultador a conta do cliente recebedor
	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	// Iniciando a nova transação 
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	// Criando a transação envia os dados para Banco de Dados
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{ID: transaction.ID}, nil
}
