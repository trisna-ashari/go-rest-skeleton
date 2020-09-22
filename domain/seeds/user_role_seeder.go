package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

func createUserRole(db *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error) {
	var userExists entity.User
	var userRoleExists entity.UserRole
	err := db.Where("uuid = ?", userRole.UserUUID).Take(&userExists).Error
	if err == nil {
		err = db.Where("user_uuid = ?", userRole.UUID).Take(&userRoleExists).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
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
