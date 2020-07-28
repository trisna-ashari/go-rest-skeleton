package routers

import (
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/interfaces"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func devRoutes(e *gin.Engine) {
	secret := interfaces.NewSecretHandler()

	e.GET("/ping", func(c *gin.Context) {
		middleware.Formatter(c, nil, "pong", nil)
	})
	e.GET("/test", func(c *gin.Context) {
		middleware.Formatter(c, nil, "ok", nil)
	})
	e.GET("/secret", func(c *gin.Context) {
		if os.Getenv("APP_ENV") == "production" {
			err := exception.ErrorTextNotFound
			_ = c.AbortWithError(http.StatusNotFound, err)
		}
	}, secret.GenerateSecret)
}
