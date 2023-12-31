package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            string `json:"id"`
	AccountFrom   *Account
	AccountFromID string `json:"account_from_id"`
	AccountTo     *Account
	AccountToID   string    `json:"account_to_id"`
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
	if t.AccountFrom.ID == t.AccountTo.ID {
		return errors.New("cannot transfer to the same account")
	}
	if t.AccountFrom == nil || t.AccountTo == nil {
		return errors.New("source and destination accounts must be specified")
	}
	return nil
}
