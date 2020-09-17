package authorization

import (
	"go-rest-skeleton/config"

	"github.com/go-redis/redis/v8"
)

// JWTAuth is represent needed dependencies.
type JWTAuth struct {
	tk *Token
}

// JWTRefreshToken is a struct.
type JWTRefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// NewJWTAuth is a constructor.
func NewJWTAuth(ck config.KeyConfig, rd *redis.Client) *JWTAuth {
	tk := NewToken(ck, rd)

	return &JWTAuth{
		tk: tk,
	}
}
