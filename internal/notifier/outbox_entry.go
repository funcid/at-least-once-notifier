package notifier

import (
	"at-least-once-notifier/internal/model"
	"time"
)

type OutboxEntry struct {
	ID        uint `gorm:"primaryKey"`
	Service   model.NotificationProvider
	Message   string
	Recipient string
	Status    model.NotificationStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
