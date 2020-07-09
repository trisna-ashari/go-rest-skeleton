package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDOptions is a struct to store option of SetRequestID.
type RequestIDOptions struct {
	AllowSetting bool
}

// SetRequestID is a middleware function uses to inject X-Request-Id to response's header.
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
