package storage

import (
	"mime/multipart"
	"path"
	"time"

	"github.com/google/uuid"
)

const (
	MaxSize    = 2048000
	ReqExpired = 24 * 60 * 60 * time.Second

	CategoryAvatar = "avatar"
)

type FileName struct {
	OriginalFileName string
	NewFileName      string
	Extension        string
}

type FileStorage interface {
	UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error)
	GetFile(UUID string) (interface{}, error)
}

type FileStorageClient struct {
	FileStorage FileStorage
}

func (c *FileStorageClient) SetDriver(com FileStorage) *FileStorageClient {
	return &FileStorageClient{FileStorage: com}
}

func FormatFileName(fileName string) *FileName {
	ext := path.Ext(fileName)
	UUID := uuid.New().String()

	return &FileName{
		OriginalFileName: fileName,
		NewFileName:      UUID,
		Extension:        ext,
	}
}

func (fn *FileName) String() string {
	return fn.NewFileName + fn.Extension
}
