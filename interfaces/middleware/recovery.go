package middleware

import (
	"go-rest-skeleton/infrastructure/message/exception"
	"net/http"

	"github.com/rollbar/rollbar-go"

	"github.com/gin-gonic/gin"
)

// Recovery is a middleware function uses to handle recovery.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				rollbar.Critical(err)
				rollbar.Wait()

				_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextAnErrorOccurred)
				return
			}
		}()

		c.Next()
	}
}
