package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

func Produce(ctx context.Context, value string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("DB_HOST")},
		Topic:   os.Getenv("TOPIC"),
	})

	select {
	case <-ctx.Done():
		return
	default:
		message := kafka.Message{Value: []byte(value)}
		if err := w.WriteMessages(ctx, message); err != nil {
			log.Printf("Error writing message: %v\n", err)
		} else {
			log.Printf("Sent message: %s\n", value)
		}

		time.Sleep(time.Second)
	}
}
