package security_test

import (
	"go-rest-skeleton/pkg/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailOrigin(t *testing.T) {
	emailOrigin, _ := security.EmailOrigin("trisna.x2+github@gmail.com")
	assert.Equal(t, emailOrigin, "trisnax2@gmail.com")
}
