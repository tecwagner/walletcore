package authenticationUser

import (
	"errors"

	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
	"github.com/tecwagner/walletcore-service/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthenticationUseCase(authGateway clientGateway.IClientGateway) *AuthenticationUseCase {
	return &AuthenticationUseCase{
		AuthGateway: authGateway,
	}
}

func (uc *AuthenticationUseCase) Execute(input AuthenticationInputDTO) (*AuthenticationOutputDTO, error) {

	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email or password must be provided")
	}
	
	user, err := uc.AuthGateway.FindByClient(input.Email)
	if err != nil {
		return nil, err
	}

	err= bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	
	if err != nil {
		return nil, errors.New("invalid password")
	}
	
	token, err := security.NewJWTToken(user)

	output := &AuthenticationOutputDTO{
		Token: token,
	}

	return output, nil
}
