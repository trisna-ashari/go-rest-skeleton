package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	hashedPassword, _ := Hash("123456")
	expectedHashedPassword := hashedPassword
	assert.Equal(t, hashedPassword, expectedHashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "123456"
	hashedPassword, _ := Hash(password)
	assert.Equal(t, VerifyPassword(string(hashedPassword), password), nil)
}
