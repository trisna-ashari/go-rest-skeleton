package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createPermission(db *gorm.DB, permission *entity.Permission) (*entity.Permission, error) {
	var permissionExists entity.Permission
	err := db.Where("module_key = ? AND permission_key = ?", permission.ModuleKey, permission.PermissionKey).
		Take(&permissionExists).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err := db.Create(permission).Error
			if err != nil {
				return permission, err
			}
			return permission, err
		}
		return permission, err
	}
	return permission, err
}
