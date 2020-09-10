package middleware

import (
	"bytes"
	"go-rest-skeleton/infrastructure/message/exception"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
