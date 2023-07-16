package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"goKafka/config"
)

func NewKafkaProducer(topicName string, partition int) (*kafka.Conn, error) {
	return kafka.DialLeader(context.Background(), "tcp", config.KafkaURI, topicName, partition)
}
