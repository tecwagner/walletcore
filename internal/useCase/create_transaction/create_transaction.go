package createTransaction

import (
	"context"

	"github.com/tecwagner/walletcore-service/internal/entity"
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	transactionGateway "github.com/tecwagner/walletcore-service/internal/gateway/transaction_gateway"
	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/uow"
)

func NewCreateTransactionUseCase(
	Uow uow.IUowInterface,
	eventDispatcher events.IEventDispatcherInterface,
	transactionCreated events.IEventInterface,
	balanceUpdated events.IEventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                Uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
		BalanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}

	// Aplicando o metodo Unit Of Work
	// O Do Executa toda a operação ao mesmo tempo
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {

		accountRepository := uc.getAccountGateway(ctx)
		transactionRepository := uc.getTransactionGateway(ctx)

		// Consultando a conta do cliente pagador
		accountFrom, err := accountRepository.FindByID(input.AccountFromID)
		if err != nil {
			return err
		}

		// Consultador a conta do cliente recebedor
		accountTo, err := accountRepository.FindByID(input.AccountToID)
		if err != nil {
			return err
		}

		// Iniciando a nova transação
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		// Metodo que realiza alteração no saldo da conta
		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		// Criando a transação envia os dados para Banco de Dados
		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		output.ID = transaction.ID
		output.AccountFromID = input.AccountFromID
		output.AccountToID = input.AccountToID
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountFromID = input.AccountFromID
		balanceUpdatedOutput.AccountToID = input.AccountToID
		balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance

		return nil

	})

	if err != nil {
		return nil, err
	}

	// Registrando e Enviado evento kafka Transaction
	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	// Registrando e Enviado evento kafka Balance
	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountGateway(ctx context.Context) accountGateway.IAccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(accountGateway.IAccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionGateway(ctx context.Context) transactionGateway.ITransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(transactionGateway.ITransactionGateway)
}
