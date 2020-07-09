package middleware

import (
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/infrastructure/persistence"
)

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
