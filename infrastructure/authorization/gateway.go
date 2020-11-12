package authorization

import (
	"encoding/base64"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	lenOfAuthorization  = 2
	authTypeBasic       = "Basic"
	authTypeBearer      = "Bearer"
	authTypeApiKey      = "ApiKey"
	authTypeAccessToken = "AccessToken"
)

// Gateway is a struct.
type Gateway struct {
	ba *BasicAuth
	ja *JWTAuth
	oa *OauthAuth
	US repository.UserRepository
	RS repository.RoleRepository
}

// NewAuthGateway is a constructor.
func NewAuthGateway(
	ba *BasicAuth,
	ja *JWTAuth,
	oa *OauthAuth,
	us repository.UserRepository,
	rs repository.RoleRepository) *Gateway {
	return &Gateway{
		ba: ba,
		ja: ja,
		oa: oa,
		US: us,
		RS: rs,
	}
}

// AuthGateway is authentication gateway based on supported authentication type.
func AuthGateway(g *Gateway, c *gin.Context) (*entity.User, error) {
	var userAuth *entity.User
	headerAuth := c.Request.Header.Get("Authorization")
	headerAuthType := strings.SplitN(headerAuth, " ", 2)
	queryApiKeyAuth := c.DefaultQuery("api_key", "")
	queryAccessTokenAuth := c.DefaultQuery("access_token", "")

	if len(headerAuthType) != lenOfAuthorization && queryApiKeyAuth == "" && queryAccessTokenAuth == "" {
		c.Set("errorTracingCode", exception.ErrorCodeIFAUGA001)
		return userAuth, exception.ErrorTextUnauthorized
	}

	var authType string
	if len(headerAuthType) == lenOfAuthorization {
		authType = headerAuthType[0]
	}

	if queryApiKeyAuth != "" {
		authType = authTypeApiKey
	}

	if queryAccessTokenAuth != "" {
		authType = authTypeAccessToken
	}

	if !(authType == authTypeBasic ||
		authType == authTypeBearer ||
		authType == authTypeApiKey ||
		authType == authTypeAccessToken) {
		c.Set("errorTracingCode", exception.ErrorCodeIFAUGA002)
		return userAuth, exception.ErrorTextUnauthorized
	}

	if authType == authTypeBasic {
		payload, _ := base64.StdEncoding.DecodeString(headerAuthType[1])
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

	if authType == authTypeBearer || authType == authTypeAccessToken {
		accessDetails, errJWT := TokenValid(c, g.ja.tk)
		if errJWT != nil {
			c.Set("errorTracingCode", exception.ErrorCodeIFAUGA005)
			return userAuth, exception.ErrorTextUnauthorized
		}

		if accessDetails.UUID == "" {
			c.Set("errorTracingCode", exception.ErrorCodeIFAUGA006)
			return userAuth, exception.ErrorTextUnauthorized
		}

		c.Set("UUID", accessDetails.UUID)
	}

	if authType == authTypeApiKey {
		panic("implement me")
	}

	return userAuth, nil
}
