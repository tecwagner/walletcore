package authenticationUser

import clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"

type AuthenticationUseCase struct {
	AuthGateway clientGateway.IClientGateway
}