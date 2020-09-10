package middleware_test

import (
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	expectedResponse := `{"code":200,"data":null,"message":"OK"}`
	var actualResponse string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", func(c *gin.Context) {
		middleware.Formatter(c, nil, "api.msg.success.common.ok", nil)
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)
	actualResponse = w.Body.String()

	assert.Equal(t, expectedResponse, actualResponse)
}
