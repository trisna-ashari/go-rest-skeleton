package routers

import (
	"context"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
	"go-rest-skeleton/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const (
	authServerURL = "http://localhost:8181"
)

var (
	oauthConfig = oauth2.Config{
		ClientID:     "go-rest-skeleton",
		ClientSecret: "go-rest-skeleton",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:8181/oauth2/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
	otherOauthConfig = oauth2.Config{
		ClientID:     "1",
		ClientSecret: "1",
		Scopes:       []string{"read", "write"},
		RedirectURL:  "http://localhost:8181/other/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
	globalToken *oauth2.Token // Non-concurrent security
)

func oauthClientRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	c := e.Group("oauth2")
	c.GET("/", redirectHandler)
	c.GET("/callback", callbackHandler)
	c.GET("/refresh", refreshHandler)

	o := e.Group("/other")
	o.GET("/", func(c *gin.Context) {
		u := otherOauthConfig.AuthCodeURL("ok")
		http.Redirect(c.Writer, c.Request, u, http.StatusFound)
	})
}

func redirectHandler(c *gin.Context) {
	u := oauthConfig.AuthCodeURL("ok")
	http.Redirect(c.Writer, c.Request, u, http.StatusFound)
}

func callbackHandler(c *gin.Context) {
	r := c.Request
	errParse := r.ParseForm()
	if errParse != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errParse)
		return
	}

	state := r.Form.Get("state")
	if state != "ok" {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextAnErrorOccurred)
		return
	}
	code := r.Form.Get("code")
	if code == "" {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextAnErrorOccurred)
		return
	}
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextAnErrorOccurred)
		return
	}

	globalToken = token

	response.NewSuccess(c, token, success.AuthSuccessfullyLogin).JSON()
}

func refreshHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
	if globalToken == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	globalToken.Expiry = time.Now()
	token, err := oauthConfig.TokenSource(context.Background(), globalToken).Token()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalToken = token

	response.NewSuccess(c, token, success.AuthSuccessfullyRefreshToken).JSON()
}
