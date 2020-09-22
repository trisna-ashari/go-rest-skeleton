package middleware_test

import (
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/registry"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/pkg/util"
	"log"
	"testing"

	"github.com/bxcodec/faker"
	"gorm.io/driver/mysql"

	"github.com/go-redis/redis/v8"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Repositories represent it self.
type Repositories struct {
	Permission repository.PermissionRepository
	Role       repository.RoleRepository
	User       repository.UserRepository
	db         *gorm.DB
}

// RedisService represent it self.
type RedisService struct {
	Auth   authorization.AuthInterface
	Client *redis.Client
}

// Dependencies represent it self.
type Dependencies struct {
	db *gorm.DB
	rd *RedisService
	ag *authorization.Gateway
	at *authorization.Token
	cf *config.Config
}

func SkipThis(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test")
	}
}

func InitConfig() *config.Config {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	return config.New()
}

func Setup() *Dependencies {
	conf := InitConfig()
	db, _ := DBConnSetup(conf.DBTestConfig)
	rd, _ := RedisConnSetup(conf.RedisTestConfig)
	dbService, _ := DBServiceSetup(db)
	redisService, _ := RedisServiceSetup(rd)

	authBasic := authorization.NewBasicAuth(dbService.User)
	authJWT := authorization.NewJWTAuth(conf.KeyConfig, redisService.Client)
	authGateway := authorization.NewAuthGateway(authBasic, authJWT, dbService.User, dbService.Role)
	authToken := authorization.NewToken(conf.KeyConfig, redisService.Client)

	return &Dependencies{
		db: dbService.db,
		rd: redisService,
		ag: authGateway,
		at: authToken,
		cf: conf,
	}
}

// DBConnSetup is a function.
func DBConnSetup(config config.DBTestConfig) (*gorm.DB, error) {
	dbURL := ""
	switch config.DBDriver {
	case "postgres":
		dbURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBName,
			config.DBPassword,
		)
	case "mysql":
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBName,
		)
	}

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(mysql.Open(dbURL), gormConfig)

	if err != nil {
		return nil, err
	}

	tables := registry.CollectTableNames()
	for _, table := range tables {
		err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name)).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	entities := registry.CollectEntities()
	for _, model := range entities {
		err := db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db, nil
}

// DBServiceSetup will initialize db connection and return repositories.
func DBServiceSetup(db *gorm.DB) (*Repositories, error) {
	return &Repositories{
		Permission: persistence.NewPermissionRepository(db),
		Role:       persistence.NewRoleRepository(db),
		User:       persistence.NewUserRepository(db),
		db:         db,
	}, nil
}

// RedisConnSetup will initialize redis connection and return redis client.
func RedisConnSetup(config config.RedisTestConfig) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	}), nil
}

// RedisServiceSetup will initialize connection to redis server.
func RedisServiceSetup(rc *redis.Client) (*RedisService, error) {
	return &RedisService{
		Auth:   authorization.NewAuth(rc),
		Client: rc,
	}, nil
}

func seedUser(db *gorm.DB) (*entity.User, *entity.UserFaker, error) {
	userFaker := entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user := entity.User{
		Name:     userFaker.Name,
		Email:    userFaker.Email,
		Phone:    userFaker.Phone,
		Password: userFaker.Password,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, nil, err
	}
	return &user, &userFaker, nil
}
