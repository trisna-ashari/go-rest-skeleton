package util

import (
	"strings"
	"unicode"
)

func SliceContains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
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
