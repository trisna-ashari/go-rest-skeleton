package authorization

import (
	"go-rest-skeleton/infrastructure/config"

	"github.com/go-redis/redis/v8"
)

// JWTAuth is represent needed dependencies.
type JWTAuth struct {
	tk *Token
}

// NewJWTAuth is a constructor.
func NewJWTAuth(ck config.KeyConfig, rd *redis.Client) *JWTAuth {
	tk := NewToken(ck, rd)

	return &JWTAuth{
		tk: tk,
	}
}
