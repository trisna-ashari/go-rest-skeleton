package notify

import (
	"bytes"
	"fmt"
	"go-rest-skeleton/pkg/util"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailChannel struct {
	EmailClient  *gomail.Dialer
	EmailMessage *gomail.Message
	emailData
	emailRequest
}

type emailData struct {
	receiver     []string
	template     string
	templateData interface{}
	language     string
}

type emailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

// SetReceiver sets a value to the receiver.
func (e *EmailChannel) SetReceiver(receiver []string) {
	e.emailData.receiver = receiver
}

// SetLanguage sets a value to the language.
func (e *EmailChannel) SetLanguage(language string) {
	e.emailData.language = language
}

// SetTemplate sets a value to the template.
func (e *EmailChannel) SetTemplate(template string) {
	e.emailData.template = template
}

// SetTemplateData sets a value to the templateData.
func (e *EmailChannel) SetTemplateData(data interface{}) {
	e.emailData.templateData = data
}

func (e *EmailChannel) setSender() {
	e.emailRequest.from = "no-reply@trivaapps.com"
}

func (e *EmailChannel) setReceiver() {
	var receivers = make([]string, len(e.emailData.receiver))
	for i, r := range e.emailData.receiver {
		receivers[i] = r
	}
	e.emailRequest.to = receivers
}

func (e *EmailChannel) setSubject() {
	e.emailRequest.subject = e.emailData.template
}

func (e *EmailChannel) setBody() {
	templateName := fmt.Sprintf("%s_%s", e.emailData.language, e.emailData.template)
	templatePath := fmt.Sprintf("%s/infrastructure/notify/template/%s.html", util.RootDir(), templateName)
	t, errParsing := template.ParseFiles(templatePath)
	if errParsing != nil {
		log.Println(errParsing)
	}

	buf := new(bytes.Buffer)
	if errBind := t.Execute(buf, e.emailData.templateData); errBind != nil {
		log.Println(errParsing)
	}

	e.emailRequest.body = buf.String()
}

// GenerateMessage sets a value to the message.
func (e *EmailChannel) GenerateMessage() {
	e.setSender()
	e.setReceiver()
	e.setSubject()
	e.setBody()

	message := gomail.NewMessage()
	message.SetHeader("From", e.emailRequest.from)
	message.SetHeader("To", e.emailRequest.to...)
	message.SetHeader("Subject", e.emailRequest.subject)
	message.SetBody("text/html", e.emailRequest.body)
	e.EmailMessage = message
}

// SendNotification will send email notification.
func (e *EmailChannel) SendNotification() error {
	if errSend := e.EmailClient.DialAndSend(e.EmailMessage); errSend != nil {
		return errSend
	}

	return nil
}
