package handler_test

import (
	"encoding/json"
	"go-rest-skeleton/infrastructure/security"
	"go-rest-skeleton/infrastructure/util"
	"go-rest-skeleton/interfaces/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSecret_Success(t *testing.T) {
	var secretData security.SecretKey
	secretHandler := handler.NewSecretHandler()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/")
	v1.GET("/secret", secretHandler.GenerateSecret)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/secret", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, req)

	response := util.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &secretData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotNil(t, secretData.PublicKey)
	assert.NotNil(t, secretData.PrivateKey)
}
