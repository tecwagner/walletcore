package createClient

import (
	"github.com/tecwagner/walletcore-service/internal/entity"
	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
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

	client, err := entity.NewClient(input.Name, input.Email)

	// Criar uma tratativa de error
	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)
	// Criar uma tratativa de error
	if err != nil {
		return nil, err
	}

	// Saida do return
	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
