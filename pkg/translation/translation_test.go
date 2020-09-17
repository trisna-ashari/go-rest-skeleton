package translation_test

import (
	"go-rest-skeleton/pkg/translation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslation(t *testing.T) {

}

func TestIsValidAcceptLanguage(t *testing.T) {
	selectedLang := "en"
	validLang := translation.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, true, validLang)

	selectedLang = "es"
	invalidLang := translation.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, false, invalidLang)
}
