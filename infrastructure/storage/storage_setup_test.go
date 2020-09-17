package storage_test

import (
	"fmt"
	"go-rest-skeleton/config"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/infrastructure/storage"
	"go-rest-skeleton/pkg/util"
	"log"
	"testing"

	"github.com/google/uuid"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	driverMysql    = "mysql"
	driverPostgres = "postgres"
)

type entities struct {
	entity interface{}
}

// SkipThis is a function.
func SkipThis(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test")
	}
}

// InitConfig will initialize config.
func InitConfig() *config.Config {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	return config.New()
}

// DBConnSetup will initialize db connection, run auto migrate, and truncate all specified tables.
func DBConnSetup(config config.DBTestConfig) (*gorm.DB, error) {
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

	entities := []entities{
		{entity: &entity.Module{}},
		{entity: &entity.Permission{}},
		{entity: &entity.Role{}},
		{entity: &entity.RolePermission{}},
		{entity: &entity.StorageCategory{}},
		{entity: &entity.StorageFile{}},
		{entity: &entity.User{}},
		{entity: &entity.UserRole{}},
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
		&entity.StorageCategory{},
		&entity.StorageFile{},
		&entity.User{},
		&entity.UserRole{},
	).Error
	if err != nil {
		return nil, err
	}

	for _, model := range entities {
		err := db.Exec(fmt.Sprintf("TRUNCATE %s", db.NewScope(model.entity).TableName())).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	return db, nil
}

// DBServiceSetup will initialize db connection and return repositories.
func DBServiceSetup(config config.DBTestConfig) (*persistence.Repositories, error) {
	db, err := DBConnSetup(config)
	if err != nil {
		return nil, err
	}

	return &persistence.Repositories{
		Permission:      persistence.NewPermissionRepository(db),
		Role:            persistence.NewRoleRepository(db),
		StorageCategory: persistence.NewStorageCategoryRepository(db),
		StorageFile:     persistence.NewStorageFileRepository(db),
		User:            persistence.NewUserRepository(db),
		DB:              db,
	}, nil
}

// MinioConnSetup will initialize connection to minio server.
func MinioConnSetup(config config.MinioConfig) *minio.Client {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

// StorageServiceSetup will initialize connection to minio server and construct storage service.
func StorageServiceSetup(config config.MinioConfig, db *gorm.DB) (*persistence.StorageService, error) {
	minioClient, err := persistence.NewMinioConnection(config)
	if err != nil {
		return nil, err
	}
	minioDriver := storage.NewMinioDriver(minioClient, config.Bucket, db)
	fileStorageClient := storage.FileStorageClient{FileStorage: minioDriver}

	return &persistence.StorageService{
		Storage: fileStorageClient.FileStorage,
	}, nil
}

// SeedStorageCategories will create dummy storage categories and save it into database.
func SeedStorageCategories(db *gorm.DB) ([]entity.StorageCategory, error) {
	categories := []entity.StorageCategory{
		{
			UUID:      uuid.New().String(),
			Slug:      "avatar",
			Name:      "Avatar",
			Path:      "test/avatar",
			MimeTypes: "image/jpg,image/jpeg,image/png,image/gif",
		},
	}

	for _, v := range categories {
		category := v
		err := db.Create(&category).Error
		if err != nil {
			return nil, err
		}
	}

	return categories, nil
}
