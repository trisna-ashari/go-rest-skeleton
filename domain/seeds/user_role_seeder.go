package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createUserRole(db *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error) {
	err := db.Create(userRole).Error
	return userRole, err
}
