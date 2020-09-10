package persistence

import (
	"go-rest-skeleton/infrastructure/config"
	"go-rest-skeleton/infrastructure/storage"

	"github.com/jinzhu/gorm"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// StorageService represent it self.
type StorageService struct {
	Storage storage.FileStorageInterface
}

// NewMinioConnection will initialize connection to minio server.
func NewMinioConnection(config config.MinioConfig) (*minio.Client, error) {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

// NewStorageService will initialize connection to minio server and construct storage service.
func NewStorageService(config config.MinioConfig, db *gorm.DB) (*StorageService, error) {
	minioClient, err := NewMinioConnection(config)
	if err != nil {
		return nil, err
	}
	minioDriver := storage.NewMinioDriver(minioClient, config.Bucket, db)
	fileStorageClient := storage.FileStorageClient{FileStorage: minioDriver}

	return &StorageService{
		Storage: fileStorageClient.FileStorage,
	}, nil
}
