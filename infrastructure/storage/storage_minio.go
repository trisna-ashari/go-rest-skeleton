package storage

import (
	"bytes"
	"context"
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/message/exception"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v7"
)

// MinioDriver represents it self.
type MinioDriver struct {
	client   *minio.Client
	bucket   string
	db       *gorm.DB
	testMode bool
}

// MinioDriver creates new MinioDriver.
func NewMinioDriver(client *minio.Client, bucket string, db *gorm.DB) *MinioDriver {
	return &MinioDriver{client: client, bucket: bucket, db: db}
}

// TestMode sets current MinioDriver run as test mode. Uses when running unit test or integration test.
func (c MinioDriver) TestMode() *MinioDriver {
	c.testMode = true
	return &c
}

// UploadFile uploads the given file to minio server.
func (c *MinioDriver) UploadFile(file *multipart.FileHeader, category string) (string, map[string]string, error) {
	var fileEntity entity.StorageFile
	var fileCategory entity.StorageCategory
	var fileAllowed bool

	fileOpen, err := file.Open()
	if err != nil {
		return "", nil, exception.ErrorTextUploadCannotOpenFile
	}
	defer fileOpen.Close()

	fileSize := file.Size
	if fileSize > int64(MaxSize) {
		return "", nil, exception.ErrorTextUploadInvalidSize
	}

	buffer := make([]byte, fileSize)
	_, _ = fileOpen.Read(buffer)
	fileType := http.DetectContentType(buffer)

	err = c.db.Where("slug = ?", category).Take(&fileCategory).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", nil, exception.ErrorTextStorageCategoryNotFound
		}
		return "", nil, err
	}

	fileMimeTypes := strings.Split(fileCategory.MimeTypes, ",")
	fileAllowed = false
	for _, v := range fileMimeTypes {
		if strings.HasPrefix(fileType, v) {
			fileAllowed = true
		}
	}
	if !fileAllowed {
		return "", nil, exception.ErrorTextUploadInvalidFileType
	}

	fileOriginalName := file.Filename
	fileName := FormatFileName(file.Filename).String()
	filePath := c.FormatFilePath(fileCategory.Path, fileName)
	fileBytes := bytes.NewReader(buffer)
	cacheControl := "max-age=31536000"
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	fileMetaData := minio.PutObjectOptions{ContentType: fileType, CacheControl: cacheControl, UserMetadata: userMetaData}
	_, err = c.client.PutObject(context.Background(), c.bucket, filePath, fileBytes, fileSize, fileMetaData)
	if err != nil {
		return "", nil, err
	}

	fileEntity.CategoryUUID = fileCategory.UUID
	fileEntity.OriginalName = fileOriginalName
	fileEntity.Name = fileName
	fileEntity.Type = fileType
	fileEntity.Size = fileSize
	fileEntity.Path = filePath
	err = c.db.Create(&fileEntity).Error
	if err != nil {
		return "", nil, err
	}

	return fileEntity.UUID, nil, nil
}

// GetFile gets file from minio server and return signed URL.
func (c *MinioDriver) GetFile(UUID string) (interface{}, error) {
	var fileEntity entity.StorageFile
	var fileCategory entity.StorageCategory

	err := c.db.Where("uuid = ?", UUID).Take(&fileEntity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	err = c.db.Where("uuid = ?", fileEntity.CategoryUUID).Take(&fileCategory).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, exception.ErrorTextStorageCategoryNotFound
		}
		return nil, err
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("inline; filename=%s", fileEntity.Name))

	filePath := c.FormatFilePath(fileCategory.Path, fileEntity.Name)
	fileURL, err := c.client.PresignedGetObject(context.Background(), c.bucket, filePath, ReqExpired, reqParams)
	if err != nil {
		return nil, err
	}

	return fileURL.String(), nil
}

// FormatFilePath formats file path by given fileName and category path.
func (c MinioDriver) FormatFilePath(fileCategoryPath string, fileName string) string {
	filePath := fileCategoryPath + "/" + fileName
	if c.testMode {
		return "test/" + filePath
	}

	return filePath
}
