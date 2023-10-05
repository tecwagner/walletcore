package createTransaction

type CreateTransactionInputDTO struct {
	AccountFromID string  `json:"account_from_id"`
	AccountToID   string  `json:"account_to_id"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID string `json:"id_transaction"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}
