package middleware

import (
	"go-rest-skeleton/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

// Formatter is a middleware function uses to generalize response format of RESTful api.
func Formatter(c *gin.Context, data interface{}, message string, meta interface{}) {
	response := successOutput{Code: c.Writer.Status(), Message: "OK"}
	response.Data = data
	response.Meta = meta

	if message != "" {
		response.Message = message
	}

	translatedMessage, language := persistence.NewTranslation(c, "error", response.Message)
	response.Message = translatedMessage

	c.Header("Accept-Language", language)
	c.JSON(c.Writer.Status(), response)
}
