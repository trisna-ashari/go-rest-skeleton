package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

func createRolePermission(db *gorm.DB, rolePermission *entity.RolePermission) (*entity.RolePermission, error) {
	var rolePermissionExists entity.RolePermission
	err := db.Where("role_uuid = ? AND permission_uuid = ?", rolePermission.RoleUUID, rolePermission.PermissionUUID).
		Take(&rolePermissionExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(rolePermission).Error
			if err != nil {
				return rolePermission, err
			}
			return rolePermission, err
		}
		return rolePermission, err
	}
	return rolePermission, err
}
