package persistence

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
	Default string
	Message string
	Language string
	Type string
}

func NewTranslation(c *gin.Context, messageType string, messageString string) (string, string) {
	bundle := i18n.NewBundle(language.Indonesian)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	accept := c.GetHeader("Accept-Language")
	translation := translationMessage{Message: messageString, Language: "en", Type: messageType}

	if accept != "" {
		if IsValidAcceptLanguage(accept) == true {
			translation.Language = accept
		}
	}
	if messageType == "success" {
		translation.Default = "OK"
	}
	if messageType == "error" {
		translation.Default = messageString
	}

	languageFile := fmt.Sprintf("languages/global.%s.yaml", translation.Language)
	bundle.MustLoadMessageFile(languageFile)
	localizer := i18n.NewLocalizer(bundle, translation.Language, translation.Language)

	fmt.Println(translation)

	translatedMessage := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    translation.Message,
			Other: translation.Default,
		},
	})

	return translatedMessage, translation.Language
}

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