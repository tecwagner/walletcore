package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            string `json:"id"`
	AccountFrom   *Account
	AccountFromID string `json:"accountFromID"`
	AccountTo     *Account
	AccountToID   string    `json:"accountToID"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	fmt.Println(transaction)
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	transaction.Commit()
	return transaction, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)

}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("value must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient balance")
	}
	return nil
}
