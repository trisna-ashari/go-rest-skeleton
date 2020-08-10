package persistence

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"

	"github.com/jinzhu/gorm"
)

// StorageCategoryRepo is a struct to store db connection.
type StorageCategoryRepo struct {
	db *gorm.DB
}

// NewStorageCategoryRepository will initialize StorageCategory repository.
func NewStorageCategoryRepository(db *gorm.DB) *StorageCategoryRepo {
	return &StorageCategoryRepo{db}
}

// StorageCategoryRepo implements the repository.StorageCategoryRepository interface.
var _ repository.StorageCategoryRepository = &StorageCategoryRepo{}

// SaveCategory will create a new StorageCategory.
func (r StorageCategoryRepo) SaveCategory(category *entity.StorageCategory) (*entity.StorageCategory, map[string]string, error) {
	panic("implement me")
}
