package producer

import (
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaProducerService struct {
	conn *kafka.Conn
}

func NewKafkaProducerSerivce(conn *kafka.Conn) KafkaProducerService {
	return KafkaProducerService{conn: conn}
}

func (s KafkaProducerService) SendMessage() {
	_, err := s.conn.WriteMessages(
		kafka.Message{Key: []byte("1"), Value: []byte("Order ID = 1 Updated")},
		kafka.Message{Key: []byte("2"), Value: []byte("Order ID = 2 Updated")},
	)

	if err != nil {
		log.Fatalf("Failed to send messages: %v", err)
	}
}
