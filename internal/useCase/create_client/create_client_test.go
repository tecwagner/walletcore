package createClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/useCase/mocks"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("IsEmailExists", "john@example.com").Return(false)
	m.On("Save", mock.Anything).Return(nil).Once()
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:     "John",
		Email:    "john@example.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "John", output.Name)
	assert.Equal(t, "john@example.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
	m.AssertNumberOfCalls(t, "IsEmailExists", 1) 
}

func TestCreateClientUseCase_EmailUnique(t *testing.T) {
	// Crie um mock do ClientGateway
	m := &mocks.ClientGatewayMock{}

	// Defina o comportamento do mock para o método IsEmailExists
	// Primeira chamada retorna true (email já existe), segunda chamada retorna false (email é único)
	m.On("IsEmailExists", "john@example.com").Return(true).Once()
	m.On("IsEmailExists", "jane@example.com").Return(false).Once()

	// Defina o comportamento do mock para o método Save
	// Espera ser chamado com um cliente
	m.On("Save", mock.Anything).Return(nil).Once()

	// Crie uma instância do CreateClientUseCase com o mock
	uc := NewCreateClientUseCase(m)
	// Tente criar um cliente com um email que já existe
	_, err := uc.Execute(CreateClientInputDTO{
		Name:  "John",
		Email: "john@example.com",
	})

	// Verifique se o erro é não nulo e contém a mensagem de erro esperada
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email is not unique")

	// Tente criar um cliente com um email único
	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "Jane",
		Email: "jane@example.com",
	})

	// Verifique se não há erros
	assert.Nil(t, err)
	// Verifique se o resultado não é nulo
	assert.NotNil(t, output)
	// Verifique se os campos do resultado são conforme o esperado
	assert.Equal(t, "Jane", output.Name)
	assert.Equal(t, "jane@example.com", output.Email)
	// Verifique se o método IsEmailExists foi chamado duas vezes
	m.AssertNumberOfCalls(t, "IsEmailExists", 2)
	// Verifique se o método Save foi chamado uma vez
	m.AssertNumberOfCalls(t, "Save", 1)
	// Garanta que todas as expectativas tenham sido atendidas
	m.AssertExpectations(t)
}
