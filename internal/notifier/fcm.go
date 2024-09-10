package notifier

import (
	"context"
	"log"
	"os"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func initFCMClient() *messaging.Client {
	ctx := context.Background()
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")

	conf := &firebase.Config{ProjectID: projectID}
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
	fcmClient, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error initializing fcm client: %v", err)
	}
	return fcmClient
}

func (svc *NotificationService) sendFCMNotification(entry OutboxEntry) error {
	message := &messaging.Message{
		Token: entry.Recipient,
		Notification: &messaging.Notification{
			Title: os.Getenv("FIREBASE_NOTIFICATION_TITLE"),
			Body:  entry.Message,
		},
	}

	_, err := svc.fcmClient.Send(context.Background(), message)
	if err != nil {
		log.Printf("FCM error: %v\n", err)
		return err
	}

	log.Println("FCM notification sent successfully")
	return nil
}
