package createTransaction


type CreateTransactionInputDTO struct {
	AccountFromID  string
	AccountToID string
	Amount float64
}

type CreateTransactionOutputDTO struct {
	ID        string	
}
