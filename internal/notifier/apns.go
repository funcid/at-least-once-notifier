package notifier

import (
	_ "crypto/tls"
	"fmt"
	"log"
	"os"

	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

func initAPNsClient() *apns.Client {
	certPath := os.Getenv("APNS_CERT_PATH")
	certPassword := os.Getenv("APNS_CERT_PASSWORD")

	cert, err := certificate.FromP12File(certPath, certPassword)
	if err != nil {
		log.Fatalf("error loading APNs certificate: %v", err)
	}

	client := apns.NewClient(cert).Production()
	return client
}

func (svc *NotificationService) sendAPNsNotification(entry OutboxEntry) error {
	topic := os.Getenv("APNS_BUNDLE_ID")

	notification := &apns.Notification{
		DeviceToken: entry.Recipient,
		Topic:       topic,
		Payload: payload.
			NewPayload().
			AlertTitle(os.Getenv("APNS_NOTIFICATION_TITLE")).
			AlertBody(entry.Message),
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
