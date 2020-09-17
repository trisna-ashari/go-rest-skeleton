package seeds

import (
	"github.com/jinzhu/gorm"
	"go-rest-skeleton/domain/entity"
)

// createStorageCategory will create fake storageCategory and insert into DB.
func createStorageCategory(db *gorm.DB, storageCategory *entity.StorageCategory) (*entity.StorageCategory, error) {
	var storageCategoryExists entity.StorageCategory
	err := db.Where("slug = ?", storageCategory.Slug).Take(&storageCategoryExists).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err := db.Create(storageCategory).Error
			if err != nil {
				return storageCategory, err
			}
			return storageCategory, err
		}
		return storageCategory, err
	}
	return storageCategory, err
}
