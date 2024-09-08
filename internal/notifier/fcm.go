package notifier

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func initFCMClient() *messaging.Client {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "your_project_id"}
	opt := option.WithCredentialsFile("/app/serviceAccountKey.json")
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

// sendFCMNotification отправляет уведомление через Firebase Cloud Messaging.
func (svc *NotificationService) sendFCMNotification(entry OutboxEntry) error {
	message := &messaging.Message{
		Token: entry.Recipient,
		Notification: &messaging.Notification{
			Title: "Notification",
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
