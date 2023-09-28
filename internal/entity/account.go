package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string    `json:"id"`
	Client    *Client   
	ClientID  string	`json:"client"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

func NewAccount(client *Client) *Account {

	if client == nil {
		return nil
	}

	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) Credit(amount float64) {

	fmt.Println(a.Client.ID)
	fmt.Println(a.Balance)

	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {

	fmt.Println(a.Client.ID)
	fmt.Println(a.Balance)

	a.Balance -= amount
	a.UpdatedAt = time.Now()
}
