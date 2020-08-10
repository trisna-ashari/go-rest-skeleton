package mock

import (
	"mime/multipart"
)

type StorageAppInterface struct {
	UploadFileFn func(file *multipart.FileHeader, category string) (string, map[string]string, error)
	GetFileFn    func(UUID string) (interface{}, error)
}

func (s *StorageAppInterface) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error) {
	return s.UploadFileFn(file, category)
}

func (s *StorageAppInterface) GetFile(UUID string) (interface{}, error) {
	return s.GetFileFn(UUID)
}
