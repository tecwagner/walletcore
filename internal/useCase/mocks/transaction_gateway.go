package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/walletcore-service/internal/entity"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
