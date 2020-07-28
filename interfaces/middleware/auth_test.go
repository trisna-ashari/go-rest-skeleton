package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuth_WithJWTAuth_Success(t *testing.T) {
	dp := Setup()
	conn, connErr := DBConnSetup(dp.cf.DBTestConfig)
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, _, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}

	jwtAuth, _ := dp.at.CreateToken(user.UUID)
	jwtAccessToken := jwtAuth.AccessToken
	authErr := dp.rd.Auth.CreateAuth(user.UUID, jwtAuth)
	if authErr != nil {
		t.Fatalf("want non error, got %#v", authErr)
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", Auth(dp.ag), func(c *gin.Context) {})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtAccessToken))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestAuth_WithJWTAuth_Failed(t *testing.T) {
	dp := Setup()
	conn, connErr := DBConnSetup(dp.cf.DBTestConfig)
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, _, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}

	jwtAuth, _ := dp.at.CreateToken(user.UUID)
	jwtAccessToken := jwtAuth.AccessToken

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", Auth(dp.ag), func(c *gin.Context) {})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtAccessToken))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}

func TestAuth_WithBasicAuth_Success(t *testing.T) {
	dp := Setup()
	conn, connErr := DBConnSetup(dp.cf.DBTestConfig)
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, userFaker, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.Email, userFaker.Password)))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", Auth(dp.ag), func(c *gin.Context) {})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestAuth_WithBasicAuth_Failed(t *testing.T) {
	dp := Setup()
	conn, connErr := DBConnSetup(dp.cf.DBTestConfig)
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, userFaker, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.Email, "Wrong"+userFaker.Password)))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", Auth(dp.ag), func(c *gin.Context) {})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}

func TestAuth_WithoutAuth_Failed(t *testing.T) {
	dp := Setup()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", Auth(dp.ag), func(c *gin.Context) {})

	c.Request, _ = http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}
