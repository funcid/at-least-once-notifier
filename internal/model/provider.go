package model

type NotificationProvider string

const (
	// ServiceFCM Firebase Cloud Messaging.
	ServiceFCM NotificationProvider = "FCM"

	// ServiceAPNs Apple Push Notification Service.
	ServiceAPNs NotificationProvider = "APNs"

	// ServiceSMS Twilio.
	ServiceSMS NotificationProvider = "SMS"
)
