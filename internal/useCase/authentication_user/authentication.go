package authenticationUser

import (
	"context"
	"errors"

	clientGateway "github.com/tecwagner/walletcore-service/internal/gateway/client_gateway"
	"github.com/tecwagner/walletcore-service/internal/telemetry"
	"github.com/tecwagner/walletcore-service/pkg/security"
	"go.opentelemetry.io/otel/codes"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthenticationUseCase(authGateway clientGateway.IClientGateway, telemetry telemetry.Telemetry) *AuthenticationUseCase {
	return &AuthenticationUseCase{
		AuthGateway: authGateway,
		telemetry:   telemetry,
	}
}

func (uc *AuthenticationUseCase) Execute(input AuthenticationInputDTO) (*AuthenticationOutputDTO, error) {

	ctx := context.Background()
	ctx, span := uc.telemetry.Start(ctx, "useCaseAuth")
	defer span.End()

	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email or password must be provided")
	}

	user, err := uc.AuthGateway.FindByClient(input.Email)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, errors.New("invalid password")
	}

	token, err := security.NewJWTToken(user)

	output := &AuthenticationOutputDTO{
		Token: token,
	}

	return output, nil
}
