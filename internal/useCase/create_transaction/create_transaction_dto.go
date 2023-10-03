package createTransaction


type CreateTransactionInputDTO struct {
	AccountFromID  string
	AccountToID string
	Amount float64
}

type CreateTransactionOutputDTO struct {
	ID        string	
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}