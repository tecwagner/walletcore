package createAccount

import (
	accountGateway "github.com/tecwagner/walletcore-service/internal/gateway/account_gateway"
	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
)

type CreateAccountUseCase struct {
	AccountGateway accountGateway.IAccountGateway
	ClientGateway  clientGateway.IClientGateway
}
