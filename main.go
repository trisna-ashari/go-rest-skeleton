package main

import (
	"go-rest-skeleton/infrastructure/config"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/cmd"
	"go-rest-skeleton/interfaces/routers"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/joho/godotenv"
)

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

	// Connect to redis
	redisServices, errRedis := persistence.NewRedisDB(conf.RedisConfig)
	if errRedis != nil {
		panic(errRedis)
	}

	// Init App
	app := cmd.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		router := routers.NewRouter(conf, dbServices, redisServices).Start()

		// Run app at defined port
		appPort := os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "8888"
		}
		log.Println(router.Run(":" + appPort))
		return nil
	}

	// Init Cli
	cliCommands := cmd.NewCommand(dbServices)
	app.Commands = cliCommands
	err := app.Run(os.Args)
	if err != nil {
		panic(app)
	}
}
