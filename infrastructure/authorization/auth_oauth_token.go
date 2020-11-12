package authorization

import (
	"context"
	"go-rest-skeleton/config"

	"github.com/go-oauth2/oauth2/v4"

	"github.com/go-redis/redis/v8"
)

// OauthAuth is represent needed dependencies.
type OauthAuth struct {
	tk *Token
	au *ClientData
}

// OauthRefreshToken is a struct.
type OauthRefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// NewOauthAuth is a constructor.
func NewOauthAuth(ck config.KeyConfig, rd *redis.Client) *OauthAuth {
	tk := NewToken(ck, rd)
	au := NewAuth(rd)

	return &OauthAuth{
		tk: tk,
		au: au,
	}
}

func (oa *OauthAuth) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	token, _ := oa.tk.CreateToken(data.UserID, data.Client.GetID())
	err = oa.au.CreateAuth(data.UserID, token)
	if err != nil {
		return
	}

	access = token.AccessToken
	refresh = ""
	if isGenRefresh {
		refresh = token.RefreshToken
	}

	return access, refresh, nil
}
