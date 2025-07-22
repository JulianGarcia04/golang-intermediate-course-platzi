package factory

import (
	"errors"
	"log"
)

// SMS & Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// sms
type SMSNotification struct {
}

func (s *SMSNotification) SendNotification() {
	log.Println("Enviando notificación via SMS")
}

func (s *SMSNotification) GetSender() ISender {
	return &SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (s *SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (s *SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// email

type EmailNotification struct{}

func (e *EmailNotification) SendNotification() {
	log.Println("Enviando notificación via Email")
}

func (e *EmailNotification) GetSender() ISender {
	return &EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (e *EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (e *EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "Email":
		return &EmailNotification{}, nil
	default:
		return nil, errors.New("No such notification type")
	}
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func GetMethod(f INotificationFactory) {
	notification := f.GetSender()

	log.Println(notification.GetSenderMethod())
}

func Init() {
	smsFactory, _ := GetNotificationFactory("SMS")

	emailFactory, _ := GetNotificationFactory("Email")

	SendNotification(smsFactory)

	GetMethod(smsFactory)

	GetMethod(emailFactory)

	SendNotification(emailFactory)
}
