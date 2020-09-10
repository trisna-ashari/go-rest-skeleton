package notify

type SMSChannel struct {
}

// SetReceiver sets a value to the receiver.
func (s SMSChannel) SetReceiver(receiver []string) {
	panic("implement me")
}

// SetLanguage sets a value to the language.
func (s *SMSChannel) SetLanguage(language string) {
	panic("implement me")
}

// SetTemplate sets a value to the template.
func (s SMSChannel) SetTemplate(template string) {
	panic("implement me")
}

// SetTemplateData sets a value to the templateData.
func (s SMSChannel) SetTemplateData(data interface{}) {
	panic("implement me")
}

func (s SMSChannel) GenerateMessage() {
	panic("implement me")
}

func (s SMSChannel) SendNotification() error {
	panic("implement me")
}
