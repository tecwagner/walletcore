package authenticationUser

import (
	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
	"github.com/tecwagner/walletcore-service/internal/telemetry"
)

type AuthenticationUseCase struct {
	AuthGateway clientGateway.IClientGateway
	telemetry   telemetry.Telemetry
}
