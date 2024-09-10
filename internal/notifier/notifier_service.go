package notifier

import (
	"at-least-once-notifier/internal/model"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/sideshow/apns2"
	"github.com/twilio/twilio-go"
	_ "github.com/twilio/twilio-go/rest/api/v2010"
	"gorm.io/gorm"
)

type NotificationService struct {
	db           *gorm.DB
	fcmClient    *messaging.Client
	apnsClient   *apns2.Client
	twilioClient *twilio.RestClient
}

func NewNotificationService(db *gorm.DB) *NotificationService {
	return &NotificationService{
		db:           db,
		fcmClient:    initFCMClient(),
		apnsClient:   initAPNsClient(),
		twilioClient: initTwilioClient(),
	}
}

func (svc *NotificationService) ProcessOutboxNotifications() error {
	outboxEntries, err := svc.getOutboxEntries()
	if err != nil {
		return err
	}

	for _, entry := range outboxEntries {
		if sendErr := svc.sendNotification(entry); sendErr == nil {
			err := svc.markAsSent(entry)
			if err != nil {
				log.Printf("Failed to mark entry as sent: %v\n", err)
			}
		} else {
			log.Printf("Failed to send notification: %v\n", sendErr)
		}
	}

	return nil
}

func (svc *NotificationService) getOutboxEntries() ([]OutboxEntry, error) {
	var outboxEntries []OutboxEntry
	// Извлечение записей со статусом "pending"
	if err := svc.db.Where("status = ?", model.StatusPending).Find(&outboxEntries).Error; err != nil {
		return nil, err
	}
	return outboxEntries, nil
}

func (svc *NotificationService) sendNotification(entry OutboxEntry) error {
	log.Printf("Sending notification to %s via %s: %s", entry.Recipient, entry.Service, entry.Message)
	var err error

	switch entry.Service {
	case model.ServiceFCM:
		err = svc.sendFCMNotification(entry)
	case model.ServiceAPNs:
		err = svc.sendAPNsNotification(entry)
	case model.ServiceSMS:
		err = svc.sendSMSNotification(entry)
	default:
		err = fmt.Errorf("unknown notification service: %s", entry.Service)
	}

	return err
}

func (svc *NotificationService) markAsSent(entry OutboxEntry) error {
	entry.Status = model.StatusSent
	return svc.db.Save(&entry).Error
}
