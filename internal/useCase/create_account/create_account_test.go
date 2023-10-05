package createAccount

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/useCase/mocks"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	// Create a new account
	client, _ := entity.NewClient("John", "john@example.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	// Create a new account

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	// UseCase Function Execute
	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDTO := CreateAccountInputDTO{ClientID: client.ID}

	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	assert.NotNil(t, output.ClientID)
	assert.NotNil(t, output.Balance)
	assert.NotNil(t, output.CreatedAt)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
