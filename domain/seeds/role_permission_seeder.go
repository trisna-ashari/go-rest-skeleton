package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createRolePermission(db *gorm.DB, rolePermission *entity.RolePermission) (*entity.RolePermission, error) {
	var rolePermissionExists entity.RolePermission
	err := db.Where("role_uuid = ? AND permission_uuid = ?", rolePermission.RoleUUID, rolePermission.PermissionUUID).
		Take(&rolePermissionExists).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
