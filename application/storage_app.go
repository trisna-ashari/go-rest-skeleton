package application

import (
	"go-rest-skeleton/infrastructure/storage"
	"mime/multipart"
)

type storageApp struct {
	ss storage.FileStorageInterface
}

// storageApp implement the StorageAppInterface.
var _ StorageAppInterface = &storageApp{}

// StorageAppInterface is an interface.
type StorageAppInterface interface {
	UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error)
	GetFile(UUID string) (interface{}, error)
}

// UploadFile is an implementation of method UploadFile.
func (s storageApp) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error) {
	return s.ss.UploadFile(file, category)
}

// GetFile is an implementation of method GetFile.
func (s storageApp) GetFile(UUID string) (interface{}, error) {
	return s.ss.GetFile(UUID)
}
