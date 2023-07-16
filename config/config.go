package config

import "fmt"

const (
	kafkaHost = "localhost"
	kafkaPort = "29092"

	TEST_TOPIC_NAME = "test-topic"
)

var KafkaURI string = fmt.Sprintf("%v:%v", kafkaHost, kafkaPort)
