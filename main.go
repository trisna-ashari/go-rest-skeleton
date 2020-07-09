package main

import (
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	// Connect to DB: postgres | mysql
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbServices, errDBServices := persistence.NewRepositories(dbDriver, dbUser, dbPassword, dbHost, dbName, dbPort)
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
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisServices, errRedis := authorization.NewRedisDB(redisHost, redisPort, redisPassword)
	if errRedis != nil {
		panic(errRedis)
	}

	// Get options
	optSetLanguage := os.Getenv("APP_LANG")
	optSetDebug := os.Getenv("APP_ENV") != "production"
	optSetRequestID, _ := strconv.ParseBool(os.Getenv("ENABLE_REQUEST_ID"))
	optSetLogger, _ := strconv.ParseBool(os.Getenv("ENABLE_LOGGER"))
	optSetCors, _ := strconv.ParseBool(os.Getenv("ENABLE_CORS"))

	// Init response options
	optResponse := middleware.ResponseOptions{
		Environment:     os.Getenv("APP_ENV"),
		DebugMode:       optSetDebug,
		DefaultLanguage: optSetLanguage,
	}

	// Init authorization
	authToken := authorization.NewToken()
	authenticate := interfaces.NewAuthenticate(dbServices.User, redisServices.Auth, authToken)

	// Init interfaces
	secret := interfaces.NewSecretHandler()
	welcomeApp := interfaces.NewWelcomeHandler(dbServices.User, authToken)
	welcomeV1 := welcomeV1Point00.NewWelcomeHandler()
	welcomeV2 := welcomeV2Point00.NewWelcomeHandler()
	userV1 := userV1Point00.NewUsers(dbServices.User, redisServices.Auth, authToken)
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.New(optResponse).Handler())
	router.Use(middleware.SetRequestID(middleware.RequestIDOptions{AllowSetting: optSetRequestID}))
	router.Use(middleware.CORSMiddleware(middleware.CORSOptions{AllowSetting: optSetCors}))
	router.Use(middleware.SetLogger(middleware.LoggerOptions{AllowSetting: optSetLogger}))
	router.Use(middleware.APIVersion())

	// Prepare group routing
	v1 := router.Group("/api/v1/external")
	v2 := router.Group("/api/v2/external")

	// Routes V1
	// Authorization
	v1.GET("/profile", middleware.AuthMiddleware(), authenticate.Profile)
	v1.POST("/login", authenticate.Login)
	v1.POST("/logout", middleware.AuthMiddleware(), authenticate.Logout)
	v1.POST("/refresh", authenticate.Refresh)
	v1.POST("/language", middleware.AuthMiddleware(), authenticate.SwitchLanguage)

	// Users
	v1.GET("/users", middleware.AuthMiddleware(), userV1.GetUsers)
	v1.POST("/users", middleware.AuthMiddleware(), userV1.SaveUser)
	v1.GET("/users/:uuid", middleware.AuthMiddleware(), userV1.GetUser)

	// Welcome
	v1.GET("/welcome_app", welcomeApp.Index)
	v1.GET("/welcome", welcomeV1.Index)

	// Ping
	v1.GET("/ping", func(c *gin.Context) {
		middleware.Formatter(c, nil, "pong", nil)
	})

	// Routes V2
	// Users
	v2.GET("/users", middleware.AuthMiddleware(), userV2.GetUsers)
	v2.POST("/users", middleware.AuthMiddleware(), userV2.SaveUser)
	v2.GET("/users/:uuid", middleware.AuthMiddleware(), userV2.GetUser)

	// Welcome
	v2.GET("/welcome", welcomeV2.Index)

	// Generate secret & refresh key
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
