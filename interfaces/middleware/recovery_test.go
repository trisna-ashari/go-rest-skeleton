package middleware_test

import (
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rollbar/rollbar-go"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestRecovery_InternalServerError(t *testing.T) {
	conf := InitConfig()
	rollbar.SetToken(conf.RollbarConfig.Token)

	optResponse := middleware.ResponseOptions{
		Environment:     conf.AppEnvironment,
		DebugMode:       conf.DebugMode,
		DefaultLanguage: conf.AppLanguage,
		DefaultTimezone: conf.AppTimezone,
	}

	expectedResponse := "{\"code\":500,\"data\":null,\"message\":\"An error occurred\"}"
	var actualResponse string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.Use(middleware.NewResponse(optResponse).Handler())
	r.Use(middleware.Recovery())
	r.GET("/test", func(c *gin.Context) {
		var timer *time.Timer = nil
		timer.Reset(10)
	})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	actualResponse = w.Body.String()
	assert.Equal(t, expectedResponse, actualResponse)
}
