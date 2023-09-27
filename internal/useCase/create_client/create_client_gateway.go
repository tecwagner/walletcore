package createClient

import "github.com/tecwagner/walletcore-service/internal/gateway"

type CreateClientUseCase struct {
	ClientGateway gateway.IClientGateway
}
