package notifier

import (
	"log"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func initTwilioClient() *twilio.RestClient {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func (svc *NotificationService) sendSMSNotification(entry OutboxEntry) error {
	from := os.Getenv("TWILIO_NUMBER")

	params := &openapi.CreateMessageParams{}
	params.SetTo(entry.Recipient)
	params.SetFrom(from)
	params.SetBody(entry.Message)

	_, err := svc.twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Twilio error: %v\n", err)
		return err
	}

	log.Println("SMS notification sent successfully")
	return nil
}
