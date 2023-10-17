package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedNewClient(t *testing.T) {
	client, err := NewClient("Joarez", "joarez@tekpixel.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Joarez", client.Name)
	assert.Equal(t, "joarez@tekpixel.com", client.Email)

}


func TestCreatedNewClientWhenArgsInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdatedClient(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com")
	err := client.Update("Joares", "joares@tekpixel.com")
	assert.Nil(t, err)
	assert.Equal(t, "Joares", client.Name)
	assert.Equal(t, "joares@tekpixel.com", client.Email)
}
func TestUpdatedClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com")
	err := client.Update("", "joares@tekpixel.com")
	assert.Contains(t, err.Error(), "name is required")
}
func TestCreatedNewSellerClient(t *testing.T) {
	client, err := NewClient("Vendedor", "vendedor@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Vendedor", client.Name)
	assert.Equal(t, "vendedor@example.com", client.Email)
}


func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
