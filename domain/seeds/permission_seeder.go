package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createPermission(db *gorm.DB, permission *entity.Permission) (*entity.Permission, error) {
	err := db.Create(permission).Error
	return permission, err
}
