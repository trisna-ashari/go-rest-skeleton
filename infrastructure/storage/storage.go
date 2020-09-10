// Package storage performs file handling (get, upload, and delete file).
// This package support multiple storage clients, such as minio, s3, google cloud storage.
// Storage clients called 'driver'.
// Possible to switch storage driver
// Design pattern: Adapter - Behavioral Design Pattern.
package storage

import (
	"mime/multipart"
	"path"
	"time"

	"github.com/google/uuid"
)

const (
	// Max allowed file size.
	MaxSize = 2048000

	// Expiry of signed URL.
	ReqExpired = 24 * 60 * 60 * time.Second

	// Collections of category.
	CategoryAvatar    = "avatar"
	CategoryDocument  = "document"
	CategoryFile      = "file"
	CategoryThumbnail = "thumbnail"
)

// FileName represents it self.
type FileName struct {
	OriginalFileName string
	NewFileName      string
	Extension        string
}

// FileStorageInterface is an interface. Needs to be implemented in StorageDriver.
type FileStorageInterface interface {
	UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error)
	GetFile(UUID string) (interface{}, error)
}

// FileStorageClient represents it self.
type FileStorageClient struct {
	FileStorage FileStorageInterface
}

// SetDriver sets the given driver to the FileStorageClient.
func (c *FileStorageClient) SetDriver(com FileStorageInterface) *FileStorageClient {
	return &FileStorageClient{FileStorage: com}
}

// FormatFileName formats the given filename and return FileName.
func FormatFileName(fileName string) *FileName {
	ext := path.Ext(fileName)
	UUID := uuid.New().String()

	return &FileName{
		OriginalFileName: fileName,
		NewFileName:      UUID,
		Extension:        ext,
	}
}

// String formats FileName.NewFileName + FileName.Extension and return as string.
func (fn *FileName) String() string {
	return fn.NewFileName + fn.Extension
}
