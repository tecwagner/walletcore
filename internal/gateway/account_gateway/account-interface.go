package accountGateway

import "github.com/tecwagner/walletcore-service/internal/entity"

type IAccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
