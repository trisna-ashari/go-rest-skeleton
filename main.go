package main

import (
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/config"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	roleV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/role"
	userV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/user"
	welcomeV1Point00 "go-rest-skeleton/interfaces/handler/v1.0/welcome"
	userV2Point00 "go-rest-skeleton/interfaces/handler/v2.0/user"
	welcomeV2Point00 "go-rest-skeleton/interfaces/handler/v2.0/welcome"
)

func main() {
	// Check .env file
	if err := godotenv.Load(); err != nil {
		panic("no .env file provided")
	}

	// Init Config
	conf := config.New()
	timeLoc, _ := time.LoadLocation(conf.AppTimezone)
	time.Local = timeLoc

	// Connect to DB
	dbServices, errDBServices := persistence.NewRepositories(conf.DBConfig)
	if errDBServices != nil {
		panic(errDBServices)
	}
	defer dbServices.Close()

	// Init DB Migrate
	errAutoMigrate := dbServices.AutoMigrate()
	if errAutoMigrate != nil {
		panic(errAutoMigrate)
	}

	// Init DB Seeds
	errDBSeeds := dbServices.Seeds()
	if errDBSeeds != nil {
		panic(errDBSeeds)
	}

	// Connect to redis
	redisServices, errRedis := persistence.NewRedisDB(conf.RedisConfig)
	if errRedis != nil {
		panic(errRedis)
	}

	// Init response options
	optResponse := middleware.ResponseOptions{
		Environment:     conf.AppEnvironment,
		DebugMode:       conf.DebugMode,
		DefaultLanguage: conf.AppLanguage,
		DefaultTimezone: conf.AppTimezone,
	}

	// Init authorization
	authBasic := authorization.NewBasicAuth(dbServices.User)
	authJWT := authorization.NewJWTAuth(conf.KeyConfig, redisServices.Client)
	authToken := authorization.NewToken(conf.KeyConfig, redisServices.Client)
	authGateway := authorization.NewAuthGateway(authBasic, authJWT)
	authenticate := interfaces.NewAuthenticate(dbServices.User, redisServices.Auth, authToken)

	// Init interfaces
	secret := interfaces.NewSecretHandler()
	welcomeApp := interfaces.NewWelcomeHandler(dbServices.User, authToken)
	welcomeV1 := welcomeV1Point00.NewWelcomeHandler()
	welcomeV2 := welcomeV2Point00.NewWelcomeHandler()
	roleV1 := roleV1Point00.NewRoles(dbServices.Role, redisServices.Auth, authToken)
	userV1 := userV1Point00.NewUsers(dbServices, redisServices.Auth, authToken)
	userV2 := userV2Point00.NewUsers(dbServices.User, redisServices.Auth, authToken)

	// Logging
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// Init gin with middleware
	if !conf.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middleware.New(optResponse).Handler())
	router.Use(middleware.SetRequestID(middleware.RequestIDOptions{AllowSetting: conf.EnableRequestID}))
	router.Use(middleware.CORSMiddleware(middleware.CORSOptions{AllowSetting: conf.EnableCors}))
	router.Use(middleware.SetLogger(middleware.LoggerOptions{AllowSetting: conf.EnableLogger}))
	router.Use(middleware.APIVersion())

	// Prepare group routing
	v1 := router.Group("/api/v1/external")
	v2 := router.Group("/api/v2/external")

	// Routes V1
	// Authorization
	v1.GET("/profile", middleware.AuthMiddleware(authGateway), authenticate.Profile)
	v1.POST("/login", authenticate.Login)
	v1.POST("/logout", middleware.AuthMiddleware(authGateway), authenticate.Logout)
	v1.POST("/refresh", authenticate.Refresh)
	v1.POST("/language", middleware.AuthMiddleware(authGateway), authenticate.SwitchLanguage)

	// Roles
	v1.GET("/roles", middleware.AuthMiddleware(authGateway), roleV1.GetRoles)
	v1.GET("/roles/:uuid", middleware.AuthMiddleware(authGateway), roleV1.GetRole)

	// Users
	v1.GET("/users", middleware.AuthMiddleware(authGateway), userV1.GetUsers)
	v1.POST("/users", middleware.AuthMiddleware(authGateway), userV1.SaveUser)
	v1.GET("/users/:uuid", middleware.AuthMiddleware(authGateway), userV1.GetUser)

	// Welcome
	v1.GET("/welcome_app", welcomeApp.Index)
	v1.GET("/welcome", welcomeV1.Index)

	// Routes V2
	// Users
	v2.GET("/users", middleware.AuthMiddleware(authGateway), userV2.GetUsers)
	v2.POST("/users", middleware.AuthMiddleware(authGateway), userV2.SaveUser)
	v2.GET("/users/:uuid", middleware.AuthMiddleware(authGateway), userV2.GetUser)

	// Welcome
	v2.GET("/welcome", welcomeV2.Index)

	// Generate secret & refresh key
	router.GET("/ping", func(c *gin.Context) {
		middleware.Formatter(c, nil, "pong", nil)
	})
	router.GET("/secret", func(c *gin.Context) {
		if os.Getenv("APP_ENV") == "production" {
			err := exception.ErrorTextNotFound
			_ = c.AbortWithError(http.StatusNotFound, err)
		}
	}, secret.GenerateSecret)

	// No route
	router.NoRoute(func(c *gin.Context) {
		err := exception.ErrorTextNotFound
		_ = c.AbortWithError(http.StatusNotFound, err)
	})

	// Run app at defined port
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8888"
	}
	panic(router.Run(":" + appPort))
}
