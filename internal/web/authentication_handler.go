package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/tecwagner/walletcore-service/internal/telemetry"
	authenticationUser "github.com/tecwagner/walletcore-service/internal/useCase/authentication_user"
	"go.opentelemetry.io/otel/codes"
)

type WebAuthenticationHandler struct {
	AuthenticationUseCase authenticationUser.AuthenticationUseCase
	telemetry             telemetry.Telemetry
}

func NewWebAuthenticationHandler(authenticationUseCase authenticationUser.AuthenticationUseCase, otel telemetry.Telemetry) *WebAuthenticationHandler {
	return &WebAuthenticationHandler{
		AuthenticationUseCase: authenticationUseCase,
		telemetry:             otel,
	}
}

func (h *WebAuthenticationHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	ctx := context.Background()

	ctx, span := h.telemetry.Start(ctx, "auth-handler")
	defer span.End()

	var dto authenticationUser.AuthenticationInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		oplog.Error().Msg(err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	output, err := h.AuthenticationUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		oplog.Error().Msg(err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		oplog.Error().Msg(err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

}
