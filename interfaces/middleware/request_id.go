package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestIDOptions struct {
	AllowSetting bool
}

func SetRequestID(options RequestIDOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		if options.AllowSetting {
			requestID = c.Request.Header.Get("Set-Request-Id")
		}

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
