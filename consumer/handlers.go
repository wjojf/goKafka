package consumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"goKafka/domain"
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

		user := domain.User{}
		if err := json.Unmarshal(m.Value, &user); err != nil {
			log.Println("Error parsing Value as User. Skipping...")
			continue
		}

		log.Printf("Handling User ID = %v", user.ID)

	}
}
