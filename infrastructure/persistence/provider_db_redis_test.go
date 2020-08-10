package persistence_test

import (
	"context"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRedisConnection_Success(t *testing.T) {
	now := time.Now()

	conf := InitConfig()
	conf.RedisConfig.RedisHost = conf.RedisTestConfig.RedisHost
	conf.RedisConfig.RedisDB = conf.RedisTestConfig.RedisDB
	conf.RedisConfig.RedisPassword = conf.RedisTestConfig.RedisPassword
	conf.RedisConfig.RedisPort = conf.RedisTestConfig.RedisPort
	conn, errConn := persistence.NewRedisConnection(conf.RedisConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	errSet := conn.Set(context.Background(), "ping", "pong", time.Since(now)).Err()
	errGet := conn.Get(context.Background(), "ping").Err()
	assert.NoError(t, errConn)
	assert.NoError(t, errSet)
	assert.NoError(t, errGet)
}

func TestNewRedisConnection_Failed(t *testing.T) {
	conf := InitConfig()
	conf.RedisConfig.RedisHost = "invalid host"
	conf.RedisConfig.RedisDB = conf.RedisTestConfig.RedisDB
	conf.RedisConfig.RedisPassword = conf.RedisTestConfig.RedisPassword
	conf.RedisConfig.RedisPort = conf.RedisTestConfig.RedisPort
	_, errConn := persistence.NewRedisConnection(conf.RedisConfig)
	assert.Error(t, errConn)
}

func TestNewRedisService_Success(t *testing.T) {
	conf := InitConfig()
	conf.RedisConfig.RedisHost = conf.RedisTestConfig.RedisHost
	conf.RedisConfig.RedisDB = conf.RedisTestConfig.RedisDB
	conf.RedisConfig.RedisPassword = conf.RedisTestConfig.RedisPassword
	conf.RedisConfig.RedisPort = conf.RedisTestConfig.RedisPort
	redisService, errRedisService := persistence.NewRedisService(conf.RedisConfig)
	if errRedisService != nil {
		t.Fatalf("want non error, got %#v", errRedisService)
	}

	var typeRedisService *persistence.RedisService
	assert.NoError(t, errRedisService)
	assert.IsType(t, typeRedisService, redisService)
}

func TestNewRedisService_Failed(t *testing.T) {
	conf := InitConfig()
	conf.RedisConfig.RedisHost = "invalid host"
	conf.RedisConfig.RedisDB = conf.RedisTestConfig.RedisDB
	conf.RedisConfig.RedisPassword = conf.RedisTestConfig.RedisPassword
	conf.RedisConfig.RedisPort = conf.RedisTestConfig.RedisPort
	_, errRedisService := persistence.NewRedisService(conf.RedisConfig)
	assert.Error(t, errRedisService)
}
