package util_test

import (
	"go-rest-skeleton/infrastructure/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslation(t *testing.T) {

}

func TestIsValidAcceptLanguage(t *testing.T) {
	selectedLang := "en"
	validLang := util.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, true, validLang)

	selectedLang = "es"
	invalidLang := util.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, false, invalidLang)
}
