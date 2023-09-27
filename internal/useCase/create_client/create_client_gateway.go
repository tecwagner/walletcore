package createClient

import clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"

type CreateClientUseCase struct {
	ClientGateway clientGateway.IClientGateway
}
