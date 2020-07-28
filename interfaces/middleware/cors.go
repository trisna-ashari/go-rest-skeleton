package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSOptions is a struct to store cors option.
type CORSOptions struct {
	AllowSetting bool
}

// CORS is a middleware function uses to inject CORS header to response's header.
func CORS(options CORSOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if options.AllowSetting {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.Next()
	}
}
