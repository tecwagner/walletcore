package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, float64(0), account.Balance)
	assert.Equal(t, client.ID, account.Client.ID)
}
func TestCreateNewAccountWhilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account, "client is requested")
}

func TestCreditAccount(t *testing.T) {
	account := &Account{}

	account.Credit(100.0)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	account := &Account{Balance: 100}

	account.Debit(50.0)
	assert.Equal(t, float64(50), account.Balance)
}

// func TestAccountCreditWithInvalidAmount(t *testing.T) {
// 	account := &Account{}

// 	account.Credit(0)
// 	assert.Equal(t, float64(0), account.Balance)
// }

// func TestAccountDebitWithInvalidAmount(t *testing.T) {
// 	account := &Account{}

// 	err := account.Debit(0)
// 	assert.Error(t, err)
// 	assert.Equal(t, float64(0), account.Balance)
// }
