package middleware

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

// APIVersion is a middleware function uses to inject X-Api-Version to response's header.
func APIVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		re, _ := regexp.Compile(`v\d*`)
		version := re.FindStringSubmatch(c.Request.RequestURI)

		if len(version) > 0 {
			c.Header("X-Api-Version", version[0])
		}

		c.Next()
	}
}
