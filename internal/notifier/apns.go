package notifier

import (
	_ "crypto/tls"
	"fmt"
	"log"

	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

func initAPNsClient() *apns.Client {
	cert, err := certificate.FromP12File("path/to/your/cert.p12", "your_password")
	if err != nil {
		log.Fatalf("error loading APNs certificate: %v", err)
	}

	client := apns.NewClient(cert).Production()
	return client
}

func (svc *NotificationService) sendAPNsNotification(entry OutboxEntry) error {
	notification := &apns.Notification{
		DeviceToken: entry.Recipient,
		Topic:       "your.bundle.id",
		Payload:     payload.NewPayload().AlertTitle("Notification").AlertBody(entry.Message),
	}

	res, err := svc.apnsClient.Push(notification)
	if err != nil {
		log.Printf("APNs error: %v\n", err)
		return err
	}
	if res.StatusCode != 200 {
		log.Printf("APNs failed: %v\n", res.Reason)
		return fmt.Errorf("APNs failed with status: %d - %s", res.StatusCode, res.Reason)
	}

	log.Println("APNs notification sent successfully")
	return nil
}
