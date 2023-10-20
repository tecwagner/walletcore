package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedNewClient(t *testing.T) {
	client, err := NewClient("Joarez", "joarez@tekpixel.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Joarez", client.Name)
	assert.Equal(t, "joarez@tekpixel.com", client.Email)

}


func TestCreatedNewClientWhenArgsInvalid(t *testing.T) {
	client, err := NewClient("", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdatedClient(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "1234")
	err := client.Update("Joares", "joares@tekpixel.com", "123456")
	assert.Nil(t, err)
	assert.Equal(t, "Joares", client.Name)
	assert.Equal(t, "joares@tekpixel.com", client.Email)
	assert.Equal(t, "123456", client.Password)
}
func TestUpdatedClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	err := client.Update("", "joares@tekpixel.com","123")
	assert.Contains(t, err.Error(), "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
