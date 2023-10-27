package web

import (
	"encoding/json"
	"net/http"

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
	var dto authenticationUser.AuthenticationInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	output, err := h.AuthenticationUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
