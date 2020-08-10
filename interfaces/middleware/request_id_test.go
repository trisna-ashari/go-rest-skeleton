package middleware_test

import (
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetRequestID_WithRequestHeader(t *testing.T) {
	expectedRequestID := uuid.New().String()
	var actualRequestID string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(middleware.SetRequestID(middleware.RequestIDOptions{AllowSetting: true}))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Set-Request-Id", expectedRequestID)
	r.ServeHTTP(w, c.Request)
	actualRequestID = w.Header().Get("X-Request-Id")
	assert.Equal(t, expectedRequestID, actualRequestID)
}

func TestSetRequestID_WithoutRequestHeader(t *testing.T) {
	var actualRequestID string
	var actualRequestIDExists bool

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(middleware.SetRequestID(middleware.RequestIDOptions{AllowSetting: true}))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)
	actualRequestID = w.Header().Get("X-Request-Id")
	if actualRequestID != "" {
		actualRequestIDExists = true
	}
	assert.Equal(t, true, actualRequestIDExists)
}
