package web

import (
	"encoding/json"
	"net/http"

	createClient "github.com/tecwagner/walletcore-service/internal/useCase/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase createClient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createClient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto createClient.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		if emailUniqueError, ok := err.(*createClient.JSONError); ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			responseJSON := createClient.JSONError{Message: emailUniqueError.Error()}
			jsonErrorBytes, _ := json.Marshal(responseJSON)
			w.Write(jsonErrorBytes)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
