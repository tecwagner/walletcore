package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

// Mock Interface Gateway
func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}
