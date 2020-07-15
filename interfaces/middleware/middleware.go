package middleware

import (
	"bytes"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSOptions is a struct to store cors option.
type CORSOptions struct {
	AllowSetting bool
}

// AuthMiddleware is a middleware function uses to handle request only from authorized user.
func AuthMiddleware(g *authorization.Gateway) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := authorization.AuthGateway(g, c)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
			return
		}
		c.Next()
	}
}

// CORSMiddleware is a middleware function uses to inject CORS header to response's header.
func CORSMiddleware(options CORSOptions) gin.HandlerFunc {
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

// MaxSizeAllowed is a middleware function uses to handle max limit size of received file by configured size.
func MaxSizeAllowed(n int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n)
		buff, errRead := c.GetRawData()
		if errRead != nil {
			_ = c.AbortWithError(http.StatusRequestEntityTooLarge, exception.ErrorTextFileTooLarge)
			return
		}
		buf := bytes.NewBuffer(buff)
		c.Request.Body = ioutil.NopCloser(buf)
	}
}
