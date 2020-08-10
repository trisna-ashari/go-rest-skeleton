package persistence

import (
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/domain/seeds"
	"go-rest-skeleton/infrastructure/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // for mysql driver (optional)
	_ "github.com/jinzhu/gorm/dialects/postgres" // for postgres driver (optional)
)

const (
	driverMysql    = "mysql"
	driverPostgres = "postgres"
)

// Repositories represent it self.
type Repositories struct {
	Permission      repository.PermissionRepository
	Role            repository.RoleRepository
	StorageFile     repository.StorageFileRepository
	StorageCategory repository.StorageCategoryRepository
	User            repository.UserRepository
	DB              *gorm.DB
}

// NewDBConnection will initialize db connection.
func NewDBConnection(config config.DBConfig) (*gorm.DB, error) {
	dbURL := ""
	switch config.DBDriver {
	case driverPostgres:
		dbURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBName,
			config.DBPassword,
		)
	case driverMysql:
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
	db.LogMode(config.DBLog)

	return db, nil
}

// NewDBService will initialize db connection and return repositories.
func NewDBService(config config.DBConfig) (*Repositories, error) {
	db, err := NewDBConnection(config)
	if err != nil {
		return nil, err
	}
	db.LogMode(config.DBLog)

	return &Repositories{
		Permission:      NewPermissionRepository(db),
		Role:            NewRoleRepository(db),
		StorageFile:     NewStorageFileRepository(db),
		StorageCategory: NewStorageCategoryRepository(db),
		User:            NewUserRepository(db),
		DB:              db,
	}, nil
}

// Close will closes the database connection.
func (s *Repositories) Close() error {
	return s.DB.Close()
}

// AutoMigrate will migrate all tables.
func (s *Repositories) AutoMigrate() error {
	return s.DB.AutoMigrate(
		&entity.Module{},
		&entity.Permission{},
		&entity.Role{},
		&entity.RolePermission{},
		&entity.StorageCategory{},
		&entity.StorageFile{},
		&entity.User{},
		&entity.UserRole{},
	).Error
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
