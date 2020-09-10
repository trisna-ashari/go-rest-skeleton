package notify

type FirebaseChannel struct {
}

func (f FirebaseChannel) SetReceiver(receiver []string) {
	panic("implement me")
}

// SetLanguage sets a value to the language.
func (f *FirebaseChannel) SetLanguage(language string) {
	panic("implement me")
}

func (f FirebaseChannel) SetTemplate(template string) {
	panic("implement me")
}

// SetTemplateData sets a value to the templateData.
func (f FirebaseChannel) SetTemplateData(data interface{}) {
	panic("implement me")
}

func (f FirebaseChannel) GenerateMessage() {
	panic("implement me")
}

func (f FirebaseChannel) SendNotification() error {
	panic("implement me")
}
