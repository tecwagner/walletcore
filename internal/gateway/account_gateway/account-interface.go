package accountGateway

import "github.com/tecwagner/walletcore-service/internal/entity"

type IAccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
