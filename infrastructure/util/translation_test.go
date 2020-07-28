package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslation(t *testing.T) {

}

func TestIsValidAcceptLanguage(t *testing.T) {
	selectedLang := "en"
	validLang := IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, true, validLang)

	selectedLang = "es"
	invalidLang := IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, false, invalidLang)
}
