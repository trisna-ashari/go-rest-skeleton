package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type translationLanguage struct {
	Available []string
}

type translationMessage struct {
	DefaultMessage string
	Message        string
	Language       string
	Type           string
}

// NewTranslation will translate message.
func NewTranslation(
	c *gin.Context,
	messageType string,
	messageString string,
	messageData map[string]interface{}) (string, string) {
	bundle := i18n.NewBundle(language.Indonesian)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	accept := c.GetHeader("Accept-Language")
	translation := translationMessage{Message: messageString, Language: "en", Type: messageType}

	if accept != "" {
		if IsValidAcceptLanguage(accept) {
			translation.Language = accept
		}
	}
	if messageType == "success" {
		translation.DefaultMessage = "OK"
	}
	if messageType == "error" {
		translation.DefaultMessage = messageString
	}

	languageFile := fmt.Sprintf("languages/global.%s.yaml", translation.Language)
	bundle.MustLoadMessageFile(languageFile)
	localizer := i18n.NewLocalizer(bundle, translation.Language, translation.Language)

	fmt.Println(translation)

	translatedMessage := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    translation.Message,
			Other: translation.DefaultMessage,
		},
		TemplateData: messageData,
	})

	return translatedMessage, translation.Language
}

// IsValidAcceptLanguage will validate given string is valid Accept-Language or not.
func IsValidAcceptLanguage(x string) bool {
	a := translationLanguage{
		Available: []string{"en", "id"},
	}
	for _, n := range a.Available {
		if x == n {
			return true
		}
	}
	return false
}
