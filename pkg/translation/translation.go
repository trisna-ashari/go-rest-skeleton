package translation

import (
	"fmt"
	"go-rest-skeleton/pkg/util"
	"strings"
	"unicode"

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

type Language struct {
	Language string `json:"language" form:"language"`
}

// NewTranslation creates a new translationMessage.
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

	languageFile := fmt.Sprintf("%s/languages/global.%s.yaml", util.RootDir(), translation.Language)
	bundle.MustLoadMessageFile(languageFile)
	translator := i18n.NewLocalizer(bundle, translation.Language, translation.Language)

	translatedMessage := translator.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    translation.Message,
			Other: translation.DefaultMessage,
		},
		TemplateData: messageData,
	})

	return SentenceCase(translatedMessage), translation.Language
}

// GetLanguage gets language from Accept-Language on request header. Default language is "en".
func GetLanguage(c *gin.Context) string {
	accept := c.GetHeader("Accept-Language")
	if accept != "" {
		if IsValidAcceptLanguage(accept) {
			return accept
		}
	}

	return "en"
}

// IsValidAcceptLanguage validates the given string is valid language or not.
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

func SentenceCase(sentence string) string {
	if sentence == "" {
		return ""
	}

	tmpString := []rune(strings.ToLower(sentence))
	tmpString[0] = unicode.ToUpper(tmpString[0])

	return string(tmpString)
}
