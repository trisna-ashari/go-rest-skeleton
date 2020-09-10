package application

import (
	"go-rest-skeleton/infrastructure/notify"
)

type notifyApp struct {
	ni notify.Notification
}

// notifyApp implement the NotifyAppInterface.
var _ NotifyAppInterface = &notifyApp{}

// NotifyAppInterface is an interface.
type NotifyAppInterface interface {
	ToEmail() *notify.Notification
	ToSMS() *notify.Notification
	Notify(receiver []string, template string, templateData interface{}, language string) *notify.Notification
}

// ToEmail is an implementation of method ToEmail.
func (n notifyApp) ToEmail() *notify.Notification {
	return n.ni.ToEmail()
}

// ToSMS is an implementation of method ToSMS.
func (n notifyApp) ToSMS() *notify.Notification {
	return n.ni.ToSMS()
}

// Notify is an implementation of method Notify.
func (n notifyApp) Notify(
	receiver []string,
	template string,
	templateData interface{},
	language string) *notify.Notification {
	return n.ni.Notify(receiver, template, templateData, language)
}
