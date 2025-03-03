package constants

// enum queue

type QueueMail string
type QueueSms string

const (
	QUEUE_EMAIL_SYSTEM QueueMail = "email:system"
	QUEUE_SMS          QueueSms  = "sms"
)
