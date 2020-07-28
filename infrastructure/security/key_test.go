package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSecret(t *testing.T) {
	_, genErr := GenerateSecret()
	assert.Nil(t, genErr)
}
