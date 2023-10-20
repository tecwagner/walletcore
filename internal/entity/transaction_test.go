package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("john", "john@example.com", "123")
	accountFrom := NewAccount(client1)

	client2, _ := NewClient("john Doe", "john-doe@example.com", "123")
	accountTo := NewAccount(client2)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, float64(100))
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(900), accountFrom.Balance)
	assert.Equal(t, float64(1100), accountTo.Balance)
}

func TestCreateTransactionWithInsuficientBalanca(t *testing.T) {
	client1, _ := NewClient("john", "john@example.com", "123")
	accountFrom := NewAccount(client1)

	client2, _ := NewClient("john Doe", "john-doe@example.com", "123")
	accountTo := NewAccount(client2)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient balance")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, accountFrom.Balance)
	assert.Equal(t, 1000.0, accountTo.Balance)
}
func TestCreateTransactionWithForTheSameAccount(t *testing.T) {
	client1, _ := NewClient("john", "john@example.com", "123")
	accountFrom := NewAccount(client1)

	accountTo := NewAccount(client1)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "cannot transfer to the same account")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, accountFrom.Balance)
	assert.Equal(t, 1000.0, accountTo.Balance)
}
