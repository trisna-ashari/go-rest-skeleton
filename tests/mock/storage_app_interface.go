package mock

import (
	"mime/multipart"
)

// StorageAppInterface is a mock of application.StorageAppInterface.
type StorageAppInterface struct {
	UploadFileFn func(file *multipart.FileHeader, category string) (string, map[string]string, error, interface{})
	GetFileFn    func(UUID string) (interface{}, error)
}

// UploadFile calls the UploadFileFn.
func (s *StorageAppInterface) UploadFile(file *multipart.FileHeader, c string) (string, map[string]string, error, interface{}) {
	return s.UploadFileFn(file, c)
}

// GetFile calls the GetFileFn.
func (s *StorageAppInterface) GetFile(UUID string) (interface{}, error) {
	return s.GetFileFn(UUID)
}
