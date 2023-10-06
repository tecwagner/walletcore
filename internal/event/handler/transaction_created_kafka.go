package handler

import (
	"fmt"
	"sync"

	"github.com/tecwagner/walletcore-service/pkg/events"
	"github.com/tecwagner/walletcore-service/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

// Recebe a mensagem e envia para o kafka e finaliza o waitGroup
func (h *TransactionCreatedKafkaHandler) Handle(message events.IEventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
}
