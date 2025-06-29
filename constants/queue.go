package constants

type QueueType string

const (
	QUEUE_EMAIL_SYSTEM QueueType = "email:system"
	QUEUE_SMS          QueueType = "sms"
)
