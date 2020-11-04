package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

// createRolePermission will create fake rolePermission and insert into DB.
func createRolePermission(
	db *gorm.DB,
	role *entity.Role,
	permission *entity.Permission,
	rolePermission *entity.RolePermission) (*entity.RolePermission, error) {
	var roleExists entity.Role
	var permissionExists entity.Permission
	var rolePermissionExists entity.RolePermission

	errRole := db.Where("name = ?", role.Name).Take(&roleExists).Error
	errPermission := db.Where("module_key = ? AND permission_key = ?", permission.ModuleKey, permission.PermissionKey).Take(&permissionExists).Error
	if errRole != nil && errPermission != nil {
		return rolePermission, nil
	}

	rolePermission.RoleUUID = roleExists.UUID
	rolePermission.PermissionUUID = permissionExists.UUID

	errRolePermission := db.Where("role_uuid = ? AND permission_uuid = ?", roleExists.UUID, permissionExists.UUID).
		Take(&rolePermissionExists).Error
	if errRolePermission != nil {
		if errors.Is(errRolePermission, gorm.ErrRecordNotFound) {
			err := db.Create(rolePermission).Error
			if err != nil {
				return rolePermission, err
			}
			return rolePermission, err
		}
		return rolePermission, errRolePermission
	}
	return rolePermission, errRolePermission
}
