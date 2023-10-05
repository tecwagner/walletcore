package createTransaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/event"
	"github.com/tecwagner/walletcore-service/internal/useCase/mocks"
	"github.com/tecwagner/walletcore-service/pkg/events"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	// Create a new account
	payer, _ := entity.NewClient("John", "john@example.com")
	accountFrom := entity.NewAccount(payer)
	accountFrom.Credit(1000)

	payee, _ := entity.NewClient("Joh Dou", "johdou@example.com")
	accountTo := entity.NewAccount(payee)
	accountTo.Credit(1000)

	// Uni Of Work
	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	//Nova Transação com os mock
	inputDTO := CreateTransactionInputDTO{
		AccountFromID: payer.ID,
		AccountToID:   payee.ID,
		Amount:        100,
	}

	// Instanciando a transaction eventos
	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()

	ctx := context.Background()

	// Criando a Transação com UseCase
	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	output, err := uc.Execute(ctx, inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
