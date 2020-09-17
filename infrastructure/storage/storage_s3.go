package storage

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type S3Adapter struct {
	S3Driver *S3Driver
}

func NewS3Driver(client *s3.S3, bucket string, db *gorm.DB) *S3Driver {
	return &S3Driver{client: client, bucket: bucket, db: db}
}

func (w S3Driver) TestMode() *S3Driver {
	w.testMode = true
	return &w
}

func (w *S3Adapter) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error, interface{}) {
	return w.S3Driver.UploadFile(file, category)
}

func (w *S3Adapter) GetFile(UUID string) (interface{}, error) {
	return w.S3Driver.GetFile(UUID)
}

type S3Driver struct {
	client   *s3.S3
	bucket   string
	db       *gorm.DB
	testMode bool
}

func (w *S3Driver) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error, interface{}) {
	return uuid.New().String(), nil, nil, nil
}

func (w *S3Driver) GetFile(UUID string) (interface{}, error) {
	return UUID, nil
}
