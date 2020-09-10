package routers

import (
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/interfaces/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func devRoutes(e *gin.Engine, r *Router) {
	ping := handler.NewPingHandler(r.conf)
	secret := handler.NewSecretHandler()

	e.GET("/api/ping", ping.Ping)
	e.GET("/api/secret", func(c *gin.Context) {
		if os.Getenv("APP_ENV") == "production" {
			err := exception.ErrorTextNotFound
			_ = c.AbortWithError(http.StatusNotFound, err)
		}
	}, secret.GenerateSecret)
}
