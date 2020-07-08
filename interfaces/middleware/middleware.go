package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/infrastructure/authorization"
	"golang.org/x/exp/errors"
	"io/ioutil"
	"net/http"
)

type CORSOptions struct {
	AllowSetting bool
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := authorization.TokenValid(c.Request)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New(err.Error()))
			return
		}
		c.Next()
	}
}

func CORSMiddleware(options CORSOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if options.AllowSetting {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
		}
		c.Next()
	}
}

//Avoid a large file from loading into memory
//If the file size is greater than 8MB dont allow it to even load into memory and waste our time.
func MaxSizeAllowed(n int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n)
		buff, errRead := c.GetRawData()
		if errRead != nil {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"status":     http.StatusRequestEntityTooLarge,
				"upload_err": "too large: upload an image less than 8MB",
			})
			c.Abort()
			return
		}
		buf := bytes.NewBuffer(buff)
		c.Request.Body = ioutil.NopCloser(buf)
	}
}