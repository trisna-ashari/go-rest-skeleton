package security_test

import (
	"go-rest-skeleton/infrastructure/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hashedPassword, _ := security.Hash("123456")
	expectedHashedPassword := hashedPassword
	assert.Equal(t, hashedPassword, expectedHashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "123456"
	hashedPassword, _ := security.Hash(password)
	assert.Equal(t, security.VerifyPassword(string(hashedPassword), password), nil)
}
