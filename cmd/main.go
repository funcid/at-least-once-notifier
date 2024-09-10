package main

import (
	"log"
	"net/http"
	"time"

	"at-least-once-notifier/internal/notifier"
	"at-least-once-notifier/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env.example file: %v", err)
	}

	log.Println("Starting notifier service...")

	// Setup database connection
	db, err := notifier.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v\n", err)
	}

	// Create notification service instance
	notifyService := notifier.NewNotificationService(db)

	// Run a ticker for periodically sending notifications from the outbox
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Start REST server
	srv := server.NewServer(notifyService)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", srv.Router()))
	}()

	for {
		select {
		case <-ticker.C:
			err := notifyService.ProcessOutboxNotifications()
			if err != nil {
				log.Printf("Error processing notifications: %v\n", err)
			}
		}
	}
}
