package consumer

import (
	"github.com/segmentio/kafka-go"
	"goKafka/config"
)

func NewKafkaConsumer(topicName string, partition int) *kafka.Reader {
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:   []string{config.KafkaURI},
			Topic:     topicName,
			Partition: partition,
			MaxBytes:  10e6,
		})
}
