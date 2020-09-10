package notification

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/notify"
)

type ForgotPassword struct {
	Notification        application.NotifyAppInterface
	NotificationOptions notify.NotificationOptions
	Receiver            string
	Template            string
	TemplateData        interface{}
	Language            string
}

func NewForgotPassword(
	receiver *entity.User,
	notification application.NotifyAppInterface,
	language string,
	options notify.NotificationOptions) *ForgotPassword {
	template := "forgot_password"
	templateData := struct {
		Name string
		URL  string
	}{
		Name: receiver.FirstName,
		URL:  "http://trivaapps.com" + options.URLPath,
	}

	return &ForgotPassword{
		Notification:        notification,
		NotificationOptions: options,
		Receiver:            receiver.Email,
		Template:            template,
		TemplateData:        templateData,
		Language:            language,
	}
}

func (n *ForgotPassword) Send() map[int]error {
	return n.Notification.Notify([]string{n.Receiver}, n.Template, n.TemplateData, n.Language).ToEmail().Send()
}
