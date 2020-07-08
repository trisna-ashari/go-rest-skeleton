package middleware

import (
	"github.com/gin-gonic/gin"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		accept := c.GetHeader("Accept")

		if accept == "application/vnd.api.v1+json" {
			c.Set("version", "v1.0")
		} else {
			c.Set("version", "v2.0")
		}

		c.Next()
	}
}
