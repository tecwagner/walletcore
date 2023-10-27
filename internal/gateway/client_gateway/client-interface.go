package clientGateway

import "github.com/tecwagner/walletcore-service/internal/entity"

type IClientGateway interface {
	Get(id string) (*entity.Client, error)
	FindByClient(email string) (*entity.Client, error)
	Save(client *entity.Client) error
	IsEmailExists(email string) bool
}
