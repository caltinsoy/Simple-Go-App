package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-crud/consumer"
	"go-crud/controller"
	"go-crud/initializers"
	"log"
	"os"
	"os/signal"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.GET("/health", controller.HealthCheck)
	r.POST("/log", func(context *gin.Context) {
		controller.CreateLog(context, ctx)
	})

	// Start the HTTP server in a goroutine
	go func() {
		if err := r.Run(":4000"); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Start the Kafka consumer in a goroutine
	go func() { consumer.Consume(ctx) }()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down gracefully...")
	// Allow some time for cleanup tasks before exiting (adjust as needed)
	time.Sleep(2 * time.Second)

	log.Println("Shutdown complete.")

}
