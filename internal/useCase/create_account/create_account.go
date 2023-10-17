package createAccount

import (
	"time"

	"github.com/tecwagner/walletcore-service/internal/entity"
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
)

func NewCreateAccountUseCase(account accountGateway.IAccountGateway, client clientGateway.IClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: account,
		ClientGateway:  client,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	// Consulta client
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	// Cria uma conta passando o client por parametro
	account := entity.NewAccount(client)

	// Salva a conta
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	output := &CreateAccountOutputDTO{
		ID:        account.ID,
		ClientID:  input.ClientID,
		Balance:   account.Balance,
		CreatedAt: time.Now(),
	}
	return output, nil
}
