package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createRolePermission(db *gorm.DB, rolePermission *entity.RolePermission) (*entity.RolePermission, error) {
	err := db.Create(rolePermission).Error
	return rolePermission, err
}
