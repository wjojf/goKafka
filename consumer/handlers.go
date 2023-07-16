package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConsumerClient struct {
	reader *kafka.Reader
}

func NewKafkaConsumerService(reader *kafka.Reader) KafkaConsumerClient {
	return KafkaConsumerClient{reader: reader}
}

func (cl KafkaConsumerClient) HandleMessages() {
	for {
		m, err := cl.reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		log.Printf("Received message: topic = %v; key=%v; value=%v", m.Topic, m.Key, m.Value)

	}
}
