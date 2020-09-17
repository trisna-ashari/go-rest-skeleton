package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

func createUserRole(db *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error) {
	var userExists entity.User
	var userRoleExists entity.UserRole
	err := db.Where("uuid = ?", userRole.UserUUID).Take(&userExists).Error
	if err == nil {
		err = db.Where("user_uuid = ?", userRole.UUID).Take(&userRoleExists).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				err := db.Create(userRole).Error
				if err != nil {
					return userRole, err
				}
				return userRole, err
			}
			return userRole, err
		}
		return userRole, err
	}
	return userRole, nil
}
