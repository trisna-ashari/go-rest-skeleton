package persistence

import (
	"errors"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"

	"gorm.io/gorm"
)

// StorageFileRepo is a struct to store db connection.
type StorageFileRepo struct {
	db *gorm.DB
}

// NewStorageFileRepository will initialize StorageFile repository.
func NewStorageFileRepository(db *gorm.DB) *StorageFileRepo {
	return &StorageFileRepo{db}
}

// StorageFileRepo implements the repository.StorageFileRepository interface.
var _ repository.StorageFileRepository = &StorageFileRepo{}

func (r StorageFileRepo) SaveFile(file *entity.StorageFile) (*entity.StorageFile, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&file).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return file, nil, nil
}

// GetFile will return file detail.
func (r *StorageFileRepo) GetFile(uuid string) (*entity.StorageFile, error) {
	var file entity.StorageFile
	err := r.db.Where("uuid = ?", uuid).Take(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserNotFound
		}
		return nil, err
	}
	return &file, nil
}
