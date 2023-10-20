package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
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
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	account := NewAccount(client)

	account.Credit(100.0)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	account := NewAccount(client)
	account.Credit(100.0)
	account.Debit(50.0)
	assert.Equal(t, float64(50), account.Balance)
}

