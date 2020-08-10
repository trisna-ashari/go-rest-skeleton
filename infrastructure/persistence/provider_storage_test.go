package persistence_test

import (
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/minio/minio-go/v7"

	"github.com/stretchr/testify/assert"
)

func TestNewMinioConnection_Success(t *testing.T) {
	conf := InitConfig()
	minioConn, errConn := persistence.NewMinioConnection(conf.MinioConfig)

	var typeMinioClient *minio.Client
	assert.NoError(t, errConn)
	assert.IsType(t, minioConn, typeMinioClient)
}

func TestNewMinioConnection_Failed(t *testing.T) {
	conf := InitConfig()
	conf.MinioConfig.Endpoint = "invalid host"
	_, errConn := persistence.NewMinioConnection(conf.MinioConfig)
	assert.Error(t, errConn)
}

func TestNewStorageService_Success(t *testing.T) {
	conf := InitConfig()
	dbConn, errDBConn := DBConnSetup(conf.DBTestConfig)
	if errDBConn != nil {
		t.Fatalf("want non error, got %#v", errDBConn)
	}
	storageService, errStorageService := persistence.NewStorageService(conf.MinioConfig, dbConn)

	var typeStorageService *persistence.StorageService
	assert.NoError(t, errStorageService)
	assert.IsType(t, storageService, typeStorageService)
}

func TestNewStorageService_Failed(t *testing.T) {
	conf := InitConfig()
	conf.MinioConfig.Endpoint = "invalid host"
	dbConn, _ := DBConnSetup(conf.DBTestConfig)
	_, errStorageService := persistence.NewStorageService(conf.MinioConfig, dbConn)

	assert.Error(t, errStorageService)
}
