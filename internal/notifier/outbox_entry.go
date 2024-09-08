package notifier

import "time"

type OutboxEntry struct {
	ID        uint   `gorm:"primaryKey"`
	Service   string // FCM, APNs, SMS
	Message   string // Сообщение
	Recipient string // Получатель сообщения
	Status    string // Например, pending, sent, failed
	CreatedAt time.Time
	UpdatedAt time.Time
}
