package config

import (
	"os"
	"strconv"
	"strings"
)

// DBConfig represent db config keys.
type DBConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
}

// RedisConfig represent redis config keys.
type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

// KeyConfig represent key config keys.
type KeyConfig struct {
	AppPrivateKey string
	AppPublicKey  string
}

// Config represent config keys.
type Config struct {
	DBConfig
	RedisConfig
	KeyConfig
	AppEnvironment  string
	AppLanguage     string
	AppTimezone     string
	EnableCors      bool
	EnableLogger    bool
	EnableRequestID bool
	DebugMode       bool
}

// New returns a new Config struct.
func New() *Config {
	return &Config{
		DBConfig: DBConfig{
			DBDriver:   getEnv("DB_DRIVER", "mysql"),
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_POST", "3306"),
			DBUser:     getEnv("DB_USER", "root"),
			DBName:     getEnv("DB_NAME", "go-rest-skeleton"),
			DBPassword: getEnv("DB_PASSWORD", ""),
		},
		RedisConfig: RedisConfig{
			RedisHost:     getEnv("REDIS_HOST", "127.0.0.1"),
			RedisPort:     getEnv("REDIS_PORT", "6379"),
			RedisPassword: getEnv("REDIS_PASSWORD", ""),
		},
		KeyConfig: KeyConfig{
			AppPrivateKey: getEnv("APP_PRIVATE_KEY", "default-private-key"),
			AppPublicKey:  getEnv("APP_PUBLIC_KEY", "default-public-key"),
		},
		AppEnvironment:  getEnv("APP_ENV", "local"),
		AppLanguage:     getEnv("APP_LANG", "en"),
		AppTimezone:     getEnv("APP_TIMEZONE", "Asia/Jakarta"),
		EnableCors:      getEnvAsBool("ENABLE_CORS", true),
		EnableLogger:    getEnvAsBool("ENABLE_LOGGER", true),
		EnableRequestID: getEnvAsBool("ENABLE_REQUEST_ID", true),
		DebugMode:       getEnv("APP_ENV", "local") != "production",
	}
}

// Simple helper function to read an environment or return a default value.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value.
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value.
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
