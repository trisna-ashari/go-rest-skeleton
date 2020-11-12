package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

func createUserRole(db *gorm.DB, user *entity.User, role *entity.Role, userRole *entity.UserRole) (*entity.UserRole, error) {
	var userExists entity.User
	var roleExists entity.Role
	var userRoleExists entity.UserRole
	errUser := db.Where("email = ?", user.Email).Take(&userExists).Error
	errRole := db.Where("name = ?", role.Name).Take(&roleExists).Error
	if errUser != nil && errRole != nil {
		return userRole, nil
	}

	userRole.UserUUID = userExists.UUID
	userRole.RoleUUID = roleExists.UUID

	errUserRole := db.Where("user_uuid = ? AND role_uuid = ?", userExists.UUID, roleExists.UUID).
		Take(&userRoleExists).Error
	if errUserRole != nil {
		if errors.Is(errUserRole, gorm.ErrRecordNotFound) {
			err := db.Create(userRole).Error
			if err != nil {
				return userRole, err
			}
			return userRole, err
		}
		return userRole, errUserRole
	}
	return userRole, nil
}
