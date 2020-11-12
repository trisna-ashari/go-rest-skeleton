package persistence

import (
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/domain/registry"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/domain/seeds"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	driverMysql    = "mysql"
	driverPostgres = "postgres"
)

// Repositories represent it self.
type Repositories struct {
	Document           repository.DocumentRepository
	Permission         repository.PermissionRepository
	Role               repository.RoleRepository
	StorageFile        repository.StorageFileRepository
	StorageCategory    repository.StorageCategoryRepository
	Tour               repository.TourRepository
	User               repository.UserRepository
	UserForgotPassword repository.UserForgotPasswordRepository
	UserPreference     repository.UserPreferenceRepository
	DB                 *gorm.DB
}

// NewDBConnection will initialize db connection.
func NewDBConnection(config config.DBConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if config.DBLog {
		gormConfig.Logger = newLogger
	}

	var dbURL string
	var db *gorm.DB
	switch config.DBDriver {
	case driverPostgres:
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			config.DBHost,
			config.DBUser,
			config.DBPassword,
			config.DBName,
			config.DBPort,
			config.DBTimeZone,
		)

		fmt.Println(dbURL)

		db, err := gorm.Open(postgres.Open(dbURL), gormConfig)
		if err != nil {
			return nil, err
		}

		return db, nil
	case driverMysql:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBName,
		)

		db, err := gorm.Open(mysql.Open(dbURL), gormConfig)
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return db, nil
}

// NewDBService will initialize db connection and return repositories.
func NewDBService(config config.DBConfig) (*Repositories, error) {
	db, err := NewDBConnection(config)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Document:           NewDocumentRepository(db),
		Permission:         NewPermissionRepository(db),
		Role:               NewRoleRepository(db),
		StorageFile:        NewStorageFileRepository(db),
		StorageCategory:    NewStorageCategoryRepository(db),
		Tour:               NewTourRepository(db),
		User:               NewUserRepository(db),
		UserForgotPassword: NewUserForgotPasswordRepository(db),
		UserPreference:     NewUserPreferenceRepository(db),
		DB:                 db,
	}, nil
}

// AutoMigrate will migrate all tables.
func (s *Repositories) AutoMigrate() error {
	var err error
	entities := registry.CollectEntities()
	for _, model := range entities {
		err = s.DB.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

// Seeds will run all seeders.
func (s *Repositories) Seeds() error {
	db := s.DB
	var err error
	for _, seed := range seeds.All() {
		errSeed := seed.Run(db)
		if errSeed != nil {
			err = errSeed
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	return err
}

// InitialSeeds will seeds predefined initial seeders.
func (s *Repositories) InitialSeeds() error {
	db := s.DB
	var err error
	for _, seed := range seeds.Init() {
		errSeed := seed.Run(db)
		if errSeed != nil {
			err = errSeed
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	return err
}
