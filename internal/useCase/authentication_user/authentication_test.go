package authenticationUser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tecwagner/walletcore-service/internal/entity"
	"github.com/tecwagner/walletcore-service/internal/useCase/mocks"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthUserUserCase_Execute(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	user, _ := entity.NewClient("John", "john@example.com", string(hashedPassword))
	m := &mocks.ClientGatewayMock{}
	m.On("FindByClient", user.Email).Return(user, nil)

	uc := NewAuthenticationUseCase(m)
	inputDTO := AuthenticationInputDTO{Email: user.Email, Password: "123"}

	fmt.Println("Auth:", inputDTO)
	output, err := uc.Execute(inputDTO)
	fmt.Println("Auth out:", output)
	assert.Nil(t, err)
	assert.NotNil(t, output.Token)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindByClient", 1)
}

func TestAuthUserUserCaseFaile_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	uc := NewAuthenticationUseCase(m)

	inputEmailEmptyDTO := AuthenticationInputDTO{Email: "", Password: "123"}

	outputEmailEmpty, err := uc.Execute(inputEmailEmptyDTO)

	assert.NotNil(t, err)
	assert.Nil(t, outputEmailEmpty)

	inputPasswordEmptyDTO := AuthenticationInputDTO{Email: "john@example.com", Password: ""}

	outputPasswordEmpty, err := uc.Execute(inputPasswordEmptyDTO)

	assert.NotNil(t, err)
	assert.Nil(t, outputPasswordEmpty)

}
