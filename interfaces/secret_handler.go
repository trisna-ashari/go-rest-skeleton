package interfaces

import (
	"go-rest-skeleton/infrastructure/security"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SecretHandler is a struct.
type SecretHandler struct {
}

// NewSecretHandler will initialize secret handler.
func NewSecretHandler() *SecretHandler {
	return &SecretHandler{}
}

// GenerateSecret will return base64 encoded string of private key and public key through rest api.
func (s *SecretHandler) GenerateSecret(c *gin.Context) {
	secretPriPubKey, err := security.GenerateSecret()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	middleware.Formatter(c, secretPriPubKey, "api.msg.success.successfully_generate_rsa_key", nil)
}
