package handler_test

import (
	"encoding/json"
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/interfaces/handler"
	"go-rest-skeleton/pkg/encoder"
	"go-rest-skeleton/pkg/util"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing_Success(t *testing.T) {
	SkipThis(t)

	var pingData handler.PingResponse

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()
	pingHandler := handler.NewPingHandler(conf)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/")
	v1.GET("/ping", pingHandler.Ping)

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &pingData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, pingData.DB, "OK")
	assert.EqualValues(t, pingData.Redis, "OK")
}

func TestPing_Failed(t *testing.T) {
	SkipThis(t)

	var pingData handler.PingResponse

	conf := config.New()
	conf.DBConfig.DBUser = "invalid user"
	conf.RedisConfig.RedisHost = "invalid host"
	pingHandler := handler.NewPingHandler(conf)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/")
	v1.GET("/ping", pingHandler.Ping)

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &pingData)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.EqualValues(t, pingData.DB, "Not OK")
	assert.EqualValues(t, pingData.Redis, "Not OK")
}
