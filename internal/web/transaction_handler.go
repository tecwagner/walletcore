package web

import (
	"encoding/json"
	"net/http"

	createTransaction "github.com/tecwagner/walletcore-service/internal/useCase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase createTransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase createTransaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createTransaction.CreateTransactionInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ctx := r.Context()

	output, err := h.CreateTransactionUseCase.Execute(ctx, dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
