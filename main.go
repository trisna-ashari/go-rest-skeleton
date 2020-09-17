package main

import (
	"go-rest-skeleton/config"
	_ "go-rest-skeleton/docs"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/cmd"
	"go-rest-skeleton/interfaces/routers"
	"log"
	"os"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/urfave/cli/v2"

	"github.com/joho/godotenv"
)

// @title Go-Rest-Skeleton API Example
// @version 1.0
// @description This is a sample of RESTful api made by go-rest-skeleton
// @contact.name Trisna Novi Ashari
// @contact.url https://github.com/trisna-ashari/go-rest-skeleton
// @contact.email trisna.x2@gmail.com
// @license.name MIT

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_key

// @host localhost:8888
// @schemes http
// main init the go-rest-skeleton.
func main() {
	// Check .env file
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	// Init Config
	conf := config.New()
	timeLoc, _ := time.LoadLocation(conf.AppTimezone)
	time.Local = timeLoc

	// Connect to DB
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		panic(errDBService)
	}
	defer dbService.Close()

	// Init DB Migrate
	errAutoMigrate := dbService.AutoMigrate()
	if errAutoMigrate != nil {
		panic(errAutoMigrate)
	}

	// Connect to redis
	redisService, errRedis := persistence.NewRedisService(conf.RedisConfig)
	if errRedis != nil {
		panic(errRedis)
	}

	// Connect to storage service
	storageService, _ := persistence.NewStorageService(conf.MinioConfig, dbService.DB)

	// Init notification service
	notificationService, _ := persistence.NewNotificationService(conf)

	// Init App
	app := cmd.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		router := routers.NewRouter(conf, dbService, redisService, storageService, notificationService).Start()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Run app at defined port
		appPort := os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "8888"
		}
		log.Println(router.Run(":" + appPort))
		return nil
	}

	// Init Cli
	cliCommands := cmd.NewCommand(dbService)
	app.Commands = cliCommands
	err := app.Run(os.Args)
	if err != nil {
		panic(app)
	}
}
