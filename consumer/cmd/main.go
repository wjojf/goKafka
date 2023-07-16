package main

import (
	"goKafka/config"
	"goKafka/consumer"
)

func main() {
	kafkaReader := consumer.NewKafkaConsumer(config.TEST_TOPIC_NAME, 0)

	service := consumer.NewKafkaConsumerService(kafkaReader)
	service.HandleMessages()
}
