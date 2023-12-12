package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
	authenticationUser "github.com/tecwagner/walletcore-service/internal/useCase/authentication_user"
)

type WebAuthenticationHandler struct {
	AuthenticationUseCase authenticationUser.AuthenticationUseCase
}

func NewWebAuthenticationHandler(authenticationUseCase authenticationUser.AuthenticationUseCase) *WebAuthenticationHandler {
	return &WebAuthenticationHandler{
		AuthenticationUseCase: authenticationUseCase,
	}
}

func (h *WebAuthenticationHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	var dto authenticationUser.AuthenticationInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		oplog.Error().Msg(err.Error())
		return
	}

	output, err := h.AuthenticationUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		oplog.Error().Msg(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		oplog.Error().Msg(err.Error())
		return
	}

}
