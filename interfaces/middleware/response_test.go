package middleware

import (
	"go-rest-skeleton/infrastructure/exception"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseOptions_Handler(t *testing.T) {
	conf := InitConfig()
	optResponse := ResponseOptions{
		Environment:     conf.AppEnvironment,
		DebugMode:       conf.DebugMode,
		DefaultLanguage: conf.AppLanguage,
		DefaultTimezone: conf.AppTimezone,
	}

	expectedResponse := "{\"code\":500,\"data\":null,\"message\":\"Internal Server Error\"}"
	var actualResponse string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(New(optResponse).Handler())
	r.GET("/test", func(c *gin.Context) {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
	})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)
	actualResponse = w.Body.String()
	assert.Equal(t, expectedResponse, actualResponse)
}
