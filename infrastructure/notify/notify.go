// Package notify perform notification handling. This package uses to generate notification message and send it
// to notification channel (email, sms, firebase, FCM). It also possible to send notification to multiple channels
// at the same time.
// Possible to implement multiple channel.
// Generate message by template and support multilingual usage.
// Design pattern: Template Method - Behavioral Design Pattern.
package notify

// NotificationInterface is an interface.
type NotificationInterface interface {
	SetReceiver(receiver []string)
	SetLanguage(language string)
	SetTemplate(template string)
	SetTemplateData(data interface{})
	GenerateMessage()
	SendNotification() error
}

// Notification represent it self.
type Notification struct {
	NotificationData     NotificationData
	EmailNotification    NotificationInterface
	SMSNotification      NotificationInterface
	FirebaseNotification NotificationInterface
	Notifications        []NotificationInterface
}

// NotificationOptions represent data that can be user to override default options.
type NotificationOptions struct {
	URLParams map[string]interface{}
	URLPath   string
}

// NotificationData represent data needed to initialize notification.
type NotificationData struct {
	receiver     []string
	template     string
	templateData interface{}
	language     string
}

// ToEmail will add email channel to be sent of notification.
func (ni *Notification) ToEmail() *Notification {
	ni.Notifications = append(ni.Notifications, ni.EmailNotification)

	return ni
}

// ToSMS will add sms channel to be sent of notification.
func (ni *Notification) ToSMS() *Notification {
	ni.Notifications = append(ni.Notifications, ni.SMSNotification)

	return ni
}

// ToFirebase will add firebase channel to be sent of notification.
func (ni *Notification) ToFireBase() *Notification {
	ni.Notifications = append(ni.Notifications, ni.FirebaseNotification)

	return ni
}

// Notify will initialize notification.
func (ni *Notification) Notify(
	receiver []string,
	template string,
	templateData interface{},
	language string) *Notification {
	ni.NotificationData = NotificationData{
		receiver:     receiver,
		template:     template,
		templateData: templateData,
		language:     language,
	}

	return ni
}

// Send will send notification to all selected channels.
func (ni *Notification) Send() map[int]error {
	var errors = make(map[int]error)

	for i, n := range ni.Notifications {
		n.SetReceiver(ni.NotificationData.receiver)
		n.SetLanguage(ni.NotificationData.language)
		n.SetTemplate(ni.NotificationData.template)
		n.SetTemplateData(ni.NotificationData.templateData)
		n.GenerateMessage()

		if err := n.SendNotification(); err != nil {
			errors[i] = err
		}
	}

	return errors
}
