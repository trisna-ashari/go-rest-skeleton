package persistence

import (
	"context"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/config"

	"github.com/go-redis/redis/v8"
)

// RedisService represent it self.
type RedisService struct {
	Auth   authorization.AuthInterface
	Client *redis.Client
}

// NewRedisConnection will initialize connection to redis server.
func NewRedisConnection(config config.RedisConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

// NewRedisService will initialize connection to redis server and construct redis service.
func NewRedisService(config config.RedisConfig) (*RedisService, error) {
	redisClient, err := NewRedisConnection(config)
	if err != nil {
		return nil, err
	}

	return &RedisService{
		Auth:   authorization.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
