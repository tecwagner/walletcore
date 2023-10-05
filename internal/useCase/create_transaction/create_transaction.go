package createTransaction

import (
	"github.com/tecwagner/walletcore-service/internal/entity"
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	transactionGateway "github.com/tecwagner/walletcore-service/internal/gateway/transaction_gateway"
	"github.com/tecwagner/walletcore-service/pkg/events"
)

func NewCreateTransactionUseCase(
	transactionGateway transactionGateway.ITransactionGateway,
	accountGateway accountGateway.IAccountGateway,
	eventDispatcher events.IEventDispatcherInterface,
	transactionCreated events.IEventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	// Consultando a conta do cliente pagador
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	// Consultador a conta do cliente recebedor
	accountTo, err := uc.AccountGateway.FindByID(input.AccountToID)
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

	// Metodo que realiza alteração no saldo da conta
	err = uc.AccountGateway.UpdateBalance(accountFrom)
	if err != nil {
		return nil, err
	}
	err = uc.AccountGateway.UpdateBalance(accountTo)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{ID: transaction.ID}

	// Enviado evento
	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil
}
