package api

import (
	"context"
	"log"

	"github.com/bd986650/rate-limiter/config"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: config.KafkaBrokers,
			Topic:   config.KafkaTopic,
		}),
	}
}

func (kp *KafkaProducer) Send(userID string, message string) {
	err := kp.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(userID),
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Println("Kafka send error:", err)
	}
}
