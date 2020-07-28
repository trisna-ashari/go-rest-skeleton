package middleware

import (
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationGateway struct {
	gw *authorization.Gateway
}

func Guard(gw *authorization.Gateway) *AuthenticationGateway {
	return &AuthenticationGateway{
		gw: gw,
	}
}

func (ag *AuthenticationGateway) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := authorization.AuthGateway(ag.gw, c)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
			return
		}
		c.Next()
	}
}

func (ag *AuthenticationGateway) Authorize(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		policy := NewPolicy(ag.gw.US, ag.gw.RS)
		if !policy.Can(action, c) {
			_ = c.AbortWithError(http.StatusForbidden, exception.ErrorTextForbidden)
			return
		}
		c.Next()
	}
}

// Auth is a middleware function uses to handle request only from authorized user.
func Auth(g *authorization.Gateway) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := authorization.AuthGateway(g, c)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
			return
		}
		c.Next()
	}
}
