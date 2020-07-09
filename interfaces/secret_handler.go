package interfaces

import (
	"encoding/base64"
	"go-rest-skeleton/infrastructure/security"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type secretKey struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// SecretHandler is a struct.
type SecretHandler struct {
}

// NewSecretHandler will initialize secret handler.
func NewSecretHandler() *SecretHandler {
	return &SecretHandler{}
}

// GenerateSecret will return base64 encoded string of private key and public key.
func (s *SecretHandler) GenerateSecret(c *gin.Context) {
	privateKey, publicKey, err := security.GenerateKey64()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	secretKey := secretKey{
		PrivateKey: base64.StdEncoding.EncodeToString([]byte(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString([]byte(publicKey)),
	}

	middleware.Formatter(c, secretKey, "api.msg.success.successfully_generate_rsa_key", nil)
}
