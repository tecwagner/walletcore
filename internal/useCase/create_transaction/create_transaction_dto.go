package createTransaction


type CreateTransactionInputDTO struct {
	AccountIDFrom  string
	AccountIDTo string
	Amount float64
}

type CreateTransactionOutputDTO struct {
	ID        string	
}
