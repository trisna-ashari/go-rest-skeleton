package middleware

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		re, _ := regexp.Compile("v\\d*")
		version := re.FindStringSubmatch(c.Request.RequestURI)

		if len(version) > 0 {
			c.Header("X-Api-Version", version[0])
		}

		c.Next()
	}
}
