package application

import (
	"go-rest-skeleton/infrastructure/storage"
	"mime/multipart"
)

type storageApp struct {
	ss storage.FileStorage
}

var _ StorageAppInterface = &storageApp{}

type StorageAppInterface interface {
	UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error)
	GetFile(UUID string) (interface{}, error)
}

func (s storageApp) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error) {
	return s.ss.UploadFile(file, category)
}

func (s storageApp) GetFile(UUID string) (interface{}, error) {
	return s.ss.GetFile(UUID)
}
