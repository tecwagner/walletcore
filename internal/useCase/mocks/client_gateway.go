package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type ClientGatewayMock struct {
	mock.Mock
}

// Mock Interface Gateway
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) FindByClient(email string) (*entity.Client, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) IsEmailExists(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}