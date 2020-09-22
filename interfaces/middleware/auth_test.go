package middleware_test

import (
	"encoding/base64"
	"fmt"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuth_WithJWTAuth_Success(t *testing.T) {
	SkipThis(t)

	dp := Setup()
	conn, errConn := DBConnSetup(dp.cf.DBTestConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, _, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
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
	r.GET("/test", middleware.Auth(dp.ag), func(c *gin.Context) {})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtAccessToken))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestAuth_WithJWTAuth_Failed(t *testing.T) {
	SkipThis(t)

	dp := Setup()
	conn, errConn := DBConnSetup(dp.cf.DBTestConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, _, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}

	jwtAuth, _ := dp.at.CreateToken(user.UUID)
	jwtAccessToken := jwtAuth.AccessToken

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", middleware.Auth(dp.ag), func(c *gin.Context) {})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtAccessToken))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}

func TestAuth_WithBasicAuth_Success(t *testing.T) {
	SkipThis(t)

	dp := Setup()
	conn, errConn := DBConnSetup(dp.cf.DBTestConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, userFaker, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.Email, userFaker.Password)))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", middleware.Auth(dp.ag), func(c *gin.Context) {})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	c.Request.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestAuth_WithBasicAuth_Failed(t *testing.T) {
	SkipThis(t)

	dp := Setup()
	conn, errConn := DBConnSetup(dp.cf.DBTestConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, userFaker, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.Email, "Wrong"+userFaker.Password)))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", middleware.Auth(dp.ag), func(c *gin.Context) {})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	c.Request.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}

func TestAuth_WithoutAuth_Failed(t *testing.T) {
	SkipThis(t)

	dp := Setup()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", middleware.Auth(dp.ag), func(c *gin.Context) {})

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)
	assert.EqualValues(t, http.StatusUnauthorized, w.Code)
}
