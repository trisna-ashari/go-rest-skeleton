package main

import (
	"flag"
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/grpc/dialer"
	"go-rest-skeleton/grpc/services"
	"go-rest-skeleton/infrastructure/persistence"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

type rpcServer interface {
	Run(int) error
}

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

	// Init DB Migrate
	errAutoMigrate := dbService.AutoMigrate()
	if errAutoMigrate != nil {
		panic(errAutoMigrate)
	}

	// Init Vars
	var (
		port     = flag.Int("port", 8080, "The services port")
		userAddr = flag.String("user", "localhost:4040", "User server addr")
	)

	flag.Parse()

	// Init gRPC server
	var server rpcServer
	switch os.Args[1] {
	case "user":
		server = services.NewUser(dbService)
	case "service":
		server = services.NewServices(initGRPCConn(*userAddr))
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}

	// Run Server
	errRun := server.Run(*port)
	if errRun != nil {
		log.Fatal(errRun)
	}
}

func initGRPCConn(addr string) *grpc.ClientConn {
	conn, err := dialer.Dial(addr)
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}

	return conn
}
