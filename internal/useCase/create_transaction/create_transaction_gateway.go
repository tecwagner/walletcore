package createTransaction

import (
	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/uow"
)

type CreateTransactionUseCase struct {
	Uow                uow.IUowInterface
	EventDispatcher    events.IEventDispatcherInterface
	TransactionCreated events.IEventInterface
	BalanceUpdated     events.IEventInterface
}
