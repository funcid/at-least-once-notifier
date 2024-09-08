package notifier

import (
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func initTwilioClient() *twilio.RestClient {
	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "your_twilio_account_sid",
		Password: "your_twilio_auth_token",
	})
}

func (svc *NotificationService) sendSMSNotification(entry OutboxEntry) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(entry.Recipient)
	params.SetFrom("your_twilio_number")
	params.SetBody(entry.Message)

	_, err := svc.twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Twilio error: %v\n", err)
		return err
	}

	log.Println("SMS notification sent successfully")
	return nil
}
