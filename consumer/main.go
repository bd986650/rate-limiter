package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "requests",
		GroupID: "logger",
	})

	log.Println("Consumer running...")
	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		log.Printf("Kafka log: %s = %s\n", string(msg.Key), string(msg.Value))
	}
}
