package persistence

import (
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/config"

	"github.com/go-redis/redis"
)

// RedisService represent it self.
type RedisService struct {
	Auth   authorization.AuthInterface
	Client *redis.Client
}

// NewRedisDB will initialize connection to redis server.
func NewRedisDB(config config.RedisConfig) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       0,
	})
	return &RedisService{
		Auth:   authorization.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
