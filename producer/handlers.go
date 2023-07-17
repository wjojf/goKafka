package producer

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"goKafka/domain"
	"log"
	"sync"
)

type KafkaProducerService struct {
	conn *kafka.Conn
}

func NewKafkaProducerSerivce(conn *kafka.Conn) KafkaProducerService {
	return KafkaProducerService{conn: conn}
}

func (s KafkaProducerService) SendMessage() {

	var wg sync.WaitGroup

	for i := 0; i <= 1; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("Sending Message #%v", i)

			user, err := json.Marshal(
				domain.User{
					ID:       i,
					Username: fmt.Sprintf("User %v", i),
					Role:     "Admin",
				},
			)

			if err != nil {
				log.Printf("Error serializing message #%v: %v", i, err)
			}

			_, err = s.conn.WriteMessages(
				kafka.Message{Value: user})

			if err != nil {
				log.Printf("Error sending message #%v: %v", i, err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
