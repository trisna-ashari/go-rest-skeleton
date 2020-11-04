package persistence

import (
	"go-rest-skeleton/config"
	"go-rest-skeleton/infrastructure/notify"

	"gopkg.in/gomail.v2"
)

type SMTPClient struct {
	Client *gomail.Dialer
}

type NotificationService struct {
	Notification *notify.Notification
	smtpClient   *SMTPClient
}

func NewSMTPClient(config config.SMTPConfig) (*SMTPClient, error) {
	dialer := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUsername, config.SMTPPassword)
	_, err := dialer.Dial()
	if err != nil {
		return nil, err
	}

	return &SMTPClient{Client: dialer}, nil
}

func NewNotificationService(config *config.Config) (*NotificationService, error) {
	smtpClient, errSMTP := NewSMTPClient(config.SMTPConfig)
	if errSMTP != nil {
		return &NotificationService{}, errSMTP
	}

	emailChannel := &notify.EmailChannel{EmailClient: smtpClient.Client}
	smsChannel := &notify.SMSChannel{}
	firebaseChannel := &notify.FirebaseChannel{}
	notification := notify.Notification{
		EmailNotification:    emailChannel,
		SMSNotification:      smsChannel,
		FirebaseNotification: firebaseChannel,
	}

	return &NotificationService{
		Notification: &notification,
		smtpClient:   smtpClient,
	}, nil
}
