package createTransaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/event"
	"github.com/tecwagner/walletcore-service/pkg/events"
)

type ClientGatewayMock struct {
	mock.Mock
}

// Mock Interface Gateway
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}
func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

// UpdateBalance implements accountGateway.IAccountGateway.
func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

// Mock Interface Gateway
func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	// Create a new account
	payer, _ := entity.NewClient("John", "john@example.com")
	accountFrom := entity.NewAccount(payer)
	accountFrom.Credit(1000)

	payee, _ := entity.NewClient("Joh Dou", "johdou@example.com")
	accountTo := entity.NewAccount(payee)
	accountTo.Credit(1000)

	// Consultando as contas
	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", payer.ID).Return(accountFrom, nil)
	mockAccount.On("FindByID", payee.ID).Return(accountTo, nil)

	// Criando transação
	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	//Nova Transação com os mock
	inputDTO := CreateTransactionInputDTO{
		AccountFromID: payer.ID,
		AccountToID:   payee.ID,
		Amount:        100,
	}

	// Instanciando a transaction eventos
	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()

	// Criando a Transação com UseCase
	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount, dispatcher, eventTransaction)
	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
