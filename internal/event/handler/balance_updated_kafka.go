package handler

import (
	"fmt"
	"sync"

	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/kafka"
)

type UpdatedBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewUpdatedBalanceKafkaHandler(kafka *kafka.Producer) *UpdatedBalanceKafkaHandler {
	return &UpdatedBalanceKafkaHandler{
		Kafka: kafka,
	}
}

// Recebe a mensagem e envia para o kafka e finaliza o waitGroup
func (h *UpdatedBalanceKafkaHandler) Handle(message events.IEventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("UpdatedBalanceKafkaHandler: ", message.GetPayload())
}
