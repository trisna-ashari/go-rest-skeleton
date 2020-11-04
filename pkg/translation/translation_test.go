package translation_test

import (
	"go-rest-skeleton/pkg/encoder"
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/translation"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslation_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external")
	v1.GET("/translation", func(c *gin.Context) {
		response.NewSuccess(c, nil, "api.msg.success.common.ok").JSON()
	})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/translation", nil)
	c.Request.Header.Set("Accept-Language", "en")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	responseDecoder := encoder.ResponseDecoder(w.Body)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotNil(t, responseDecoder["message"])
	assert.Equal(t, responseDecoder["message"], "Ok")
}

func TestNewTranslation_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external")
	v1.GET("/translation", func(c *gin.Context) {
		response.NewSuccess(c, nil, "api.msg.error.common.not_found").JSON()
	})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/translation", nil)
	c.Request.Header.Set("Accept-Language", "en")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	responseDecoder := encoder.ResponseDecoder(w.Body)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotNil(t, responseDecoder["message"])
	assert.Equal(t, responseDecoder["message"], "Not found")
}

func TestGetLanguage_WithHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external")
	v1.GET("/translation", func(c *gin.Context) {
		response.NewSuccess(c, nil, "api.msg.success.common.ok").JSON()
	})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/translation/external", nil)
	c.Request.Header.Set("Accept-Language", "en")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, "en", translation.GetLanguage(c))
}

func TestGetLanguage_WithoutHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external")
	v1.GET("/translation", func(c *gin.Context) {
		response.NewSuccess(c, nil, "api.msg.success.common.ok").JSON()
	})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/translation/external", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, "en", translation.GetLanguage(c))
}

func TestIsValidAcceptLanguage(t *testing.T) {
	selectedLang := "en"
	validLang := translation.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, true, validLang)

	selectedLang = "es"
	invalidLang := translation.IsValidAcceptLanguage(selectedLang)
	assert.Equal(t, false, invalidLang)
}
