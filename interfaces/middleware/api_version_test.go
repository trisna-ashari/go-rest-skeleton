package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestAPIVersion_URLContainsAPIPrefix(t *testing.T) {
	expectedVersion := "v1"
	var actualVersion string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, router := gin.CreateTestContext(w)
	router.Use(APIVersion())
	v1 := router.Group("/api/v1/external")
	v1.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("GET", "/api/v1/external/test", nil)
	router.ServeHTTP(w, c.Request)
	actualVersion = w.Header().Get("X-Api-Version")
	assert.Equal(t, expectedVersion, actualVersion)
}

func TestAPIVersion_URLDoesNotContainsAPIPrefix(t *testing.T) {
	expectedVersion := ""
	var actualVersion string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, router := gin.CreateTestContext(w)
	router.Use(APIVersion())
	v1 := router.Group("/api/external")
	v1.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest("GET", "/api/external/test", nil)
	router.ServeHTTP(w, c.Request)
	actualVersion = w.Header().Get("X-Api-Version")
	assert.Equal(t, expectedVersion, actualVersion)
}
