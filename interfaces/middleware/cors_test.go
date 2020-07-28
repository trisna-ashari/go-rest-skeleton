package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS_WithAcceptedHTTPMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(CORS(CORSOptions{AllowSetting: true}))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Header().Get("Access-Control-Allow-Origin"), "*")
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Credentials"), "true")
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Methods"), "POST, OPTIONS, GET, PUT, PATCH, DELETE")
}

func TestCORS_OptionsHTTPMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(CORS(CORSOptions{AllowSetting: true}))
	r.OPTIONS("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("OPTIONS", "/test", nil)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNoContent)
}
