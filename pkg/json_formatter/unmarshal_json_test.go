package json_formatter_test

import (
	"go-rest-skeleton/pkg/json_formatter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseDecoder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    nil,
			"message": "OK",
		})
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)
	response := json_formatter.ResponseDecoder(w.Body)

	assert.EqualValues(t, response["code"], http.StatusOK)
	assert.Nil(t, response["data"])
	assert.EqualValues(t, response["message"], "OK")
}
