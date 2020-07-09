package authorization

import (
	"github.com/go-redis/redis"
)

// RedisService represent it self.
type RedisService struct {
	Auth   AuthInterface
	Client *redis.Client
}

// NewRedisDB will initialize connection to redis server.
func NewRedisDB(host, port, password string) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
