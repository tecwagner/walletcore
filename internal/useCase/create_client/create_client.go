package createClient

import (
	"github.com/tecwagner/walletcore-service/internal/entity"
	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
	"golang.org/x/crypto/bcrypt"
)

func NewCreateClientUseCase(clientGateway clientGateway.IClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {

	if uc.ClientGateway.IsEmailExists(input.Email) {

		return nil, &JSONError{Message: "email is not unique"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}

	client, err := entity.NewClient(input.Name, input.Email, string(hashedPassword))	
	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)	
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
