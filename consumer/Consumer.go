package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("BROKER")},
		Topic:   os.Getenv("TOPIC"),
		GroupID: os.Getenv("GROUP_ID"),
	})

	defer r.Close()
	log.Println("Kafka consumer started.")

	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka consumer stopped.")
			return
		default:
			msg, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v\n", err)
				continue
			}
			fmt.Println("Received:", string(msg.Value))
		}
	}
}
