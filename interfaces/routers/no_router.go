package routers

import (
	"go-rest-skeleton/infrastructure/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func roRoutes(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		err := exception.ErrorTextNotFound
		_ = c.AbortWithError(http.StatusNotFound, err)
	})
}
