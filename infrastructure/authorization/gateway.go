package authorization

import (
	"encoding/base64"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	lenOfAuthorization = 2
)

// Gateway is a struct.
type Gateway struct {
	ba *BasicAuth
	ja *JWTAuth
	US repository.UserRepository
	RS repository.RoleRepository
}

// NewAuthGateway is a constructor.
func NewAuthGateway(ba *BasicAuth, ja *JWTAuth, us repository.UserRepository, rs repository.RoleRepository) *Gateway {
	return &Gateway{
		ba: ba,
		ja: ja,
		US: us,
		RS: rs,
	}
}

// AuthGateway is authentication gateway based on supported authentication type.
func AuthGateway(g *Gateway, c *gin.Context) (*entity.User, error) {
	var userAuth *entity.User
	auth := c.Request.Header.Get("Authorization")
	authType := strings.SplitN(auth, " ", 2)

	if len(authType) != lenOfAuthorization {
		c.Set("errorTracingCode", exception.ErrorCodeIFAUGA001)
		return userAuth, exception.ErrorTextUnauthorized
	}

	if !(authType[0] == "Basic" || authType[0] == "Bearer") {
		c.Set("errorTracingCode", exception.ErrorCodeIFAUGA002)
		return userAuth, exception.ErrorTextUnauthorized
	}

	if authType[0] == "Basic" {
		payload, _ := base64.StdEncoding.DecodeString(authType[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) == lenOfAuthorization {
			email := pair[0]
			password := pair[1]
			user := entity.User{
				Email:    email,
				Password: password,
			}
			userFound, errBasic, _ := g.ba.us.GetUserByEmailAndPassword(&user)
			if errBasic != nil {
				c.Set("errorTracingCode", exception.ErrorCodeIFAUGA004)
				return userAuth, exception.ErrorTextUnauthorized
			}

			c.Set("UUID", userFound.UUID)
			userAuth = userFound
		} else {
			c.Set("errorTracingCode", exception.ErrorCodeIFAUGA003)
			return userAuth, exception.ErrorTextUnauthorized
		}
	}

	if authType[0] == "Bearer" {
		bearerToken := strings.Split(auth, " ")
		if len(bearerToken) == authorizationLen {
			accessDetails, errJWT := TokenValid(c, g.ja.tk)
			if errJWT != nil {
				c.Set("errorTracingCode", exception.ErrorCodeIFAUGA005)
				return userAuth, exception.ErrorTextUnauthorized
			}

			c.Set("UUID", accessDetails.UUID)
		}
	}

	return userAuth, nil
}
