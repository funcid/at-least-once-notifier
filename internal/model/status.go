package model

type NotificationStatus string

const (
	// StatusPending указывает, что сообщение ожидает отправки.
	StatusPending NotificationStatus = "pending"

	// StatusSent указывает, что сообщение было успешно отправлено.
	StatusSent NotificationStatus = "sent"

	// StatusFailed указывает, что отправка сообщения не удалась.
	StatusFailed NotificationStatus = "failed"
)
