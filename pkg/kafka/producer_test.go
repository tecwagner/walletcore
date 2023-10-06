package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	type TransactionDtoOutput struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		ErrorMessage string `json:"error_message"`
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "123",
		Status:       "rejected",
		ErrorMessage: "you dont have limit for this transaction",
	}

	// outputJson, _ := json.Marshal(expectedOutput)

	consfigMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewProducer(&consfigMap)
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
