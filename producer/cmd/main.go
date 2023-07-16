package main

import (
	"goKafka/config"
	"goKafka/producer"
)

func main() {
	conn, err := producer.NewKafkaProducer(config.TEST_TOPIC_NAME, 0)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	service := producer.NewKafkaProducerSerivce(conn)

	service.SendMessage()
}
