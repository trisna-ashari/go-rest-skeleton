package response_test

import (
	"go-rest-skeleton/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	expectedResponse := `{"code":200,"data":null,"message":"Ok"}`
	var actualResponse string

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", func(c *gin.Context) {
		response.NewSuccess(c, nil, "api.msg.success.common.ok").JSON()
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
