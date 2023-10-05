package createTransaction

type CreateTransactionInputDTO struct {
	AccountFromID string  `json:"account_id_from"`
	AccountToID   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountFromID string  `json:"account_id_from"`
	AccountToID   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountFromID        string  `json:"account_id_from"`
	AccountToID          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}
