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
	err := client.Update("", "joares@tekpixel.com", "123")
	assert.Contains(t, err.Error(), "name is required")
}
func TestCreateCustomerWithEmptyEmailField(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	err := client.Update("Joarez", "", "123")
	assert.Contains(t, err.Error(), "email is required")
}
func TestCreateCustomerWithEmptyPasswordField(t *testing.T) {
	client, _ := NewClient("Joarez", "joarez@tekpixel.com", "123")
	err := client.Update("Joarez", "joarez@tekpixel.com", "")
	assert.Contains(t, err.Error(), "password is required")
}

func TestAddAccountBelongsToClient(t *testing.T) {

	client := &Client{ID: "123"}

	anotherClientAccount := &Account{Client: &Client{ID: "321"}}
	err := client.AddAccount(anotherClientAccount)

	if err == nil {
		t.Errorf("account does not belong to this client")
	}
	clientAccount := &Account{Client: client}
	err = client.AddAccount(clientAccount)

	// Não deve haver erro ao adicionar a conta correta ao cliente
	if err != nil {
		t.Errorf("Erro ao adicionar conta correta ao cliente: %s", err)
	}
}
