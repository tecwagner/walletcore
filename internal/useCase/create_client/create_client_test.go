package createClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type ClientGatewayMocks struct {
	mock.Mock
}

// Mock Intercafe Gateway
func (m *ClientGatewayMocks) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}
func (m *ClientGatewayMocks) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMocks{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "John",
		Email: "john@example.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "John", output.Name)
	assert.Equal(t, "john@example.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
