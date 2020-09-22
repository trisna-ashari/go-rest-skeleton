package handler

import (
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/security"
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

// @Summary Generate a secret
// @Description Retrieve base64 encoded string of private key and public key through rest api.
// @Tags development
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Success 200 {object} response.successOutput
// @Failure 400 {string} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/secret [get]
// GenerateSecret will return base64 encoded string of private key and public key through rest api.
func (s *SecretHandler) GenerateSecret(c *gin.Context) {
	secretPriPubKey, err := security.GenerateSecret()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	response.NewSuccess(c, secretPriPubKey, success.DevSuccessfullyGenerateRSAKey).JSON()
}
