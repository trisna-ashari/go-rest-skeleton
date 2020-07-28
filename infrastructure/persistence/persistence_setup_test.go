package persistence

import (
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/config"
	"go-rest-skeleton/infrastructure/util"
	"log"

	"github.com/bxcodec/faker"

	"github.com/go-redis/redis/v8"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DBConn is a function.
func DBConn() (*gorm.DB, error) {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()
	return DBConnSetup(conf.DBTestConfig)
}

// DBServices is a function.
func DBServices() (*Repositories, error) {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()
	return DBServicesSetup(conf.DBTestConfig)
}

// RedisConn is a function.
func RedisConn() (*gorm.DB, error) {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()
	return DBConnSetup(conf.DBTestConfig)
}

// RedisServices is a function.
func RedisServices() (*RedisService, error) {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()
	return RedisServicesSetup(conf.RedisConfig)
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

	db, err := gorm.Open(config.DBDriver, dbURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(false)

	err = db.AutoMigrate(
		&entity.Module{},
		&entity.Permission{},
		&entity.Role{},
		&entity.RolePermission{},
		&entity.User{},
		&entity.UserRole{},
	).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DBServicesSetup will initialize db connection and return repositories.
func DBServicesSetup(config config.DBTestConfig) (*Repositories, error) {
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

	db, err := gorm.Open(config.DBDriver, dbURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(false)

	return &Repositories{
		Permission: NewPermissionRepository(db),
		Role:       NewRoleRepository(db),
		User:       NewUserRepository(db),
		db:         db,
	}, nil
}

// RedisConnSetup will initialize connection to redis server.
func RedisConnSetup(config config.RedisConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       10,
	})
	return redisClient, nil
}

// RedisServiceSetup will initialize connection to redis server.
func RedisServicesSetup(config config.RedisConfig) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       10,
	})
	return &RedisService{
		Auth:   authorization.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}

func seedUser(db *gorm.DB) (*entity.User, *entity.UserFaker, error) {
	userFaker := entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user := entity.User{
		FirstName: userFaker.FirstName,
		LastName:  userFaker.LastName,
		Email:     userFaker.Email,
		Phone:     userFaker.Phone,
		Password:  userFaker.Password,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, nil, err
	}
	return &user, &userFaker, nil
}

func seedRoles(db *gorm.DB) ([]entity.Role, error) {
	roles := []entity.Role{
		{
			UUID: uuid.New().String(),
			Name: "Administrator",
		},
		{
			UUID: uuid.New().String(),
			Name: "User",
		},
	}
	var role entity.Role
	for _, v := range roles {
		role = v
		err := db.Create(role).Error
		if err != nil {
			return nil, err
		}
	}
	return roles, nil
}
