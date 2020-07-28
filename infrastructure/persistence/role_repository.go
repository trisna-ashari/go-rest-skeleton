package persistence

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/exception"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RoleRepo is a struct to store db connection.
type RoleRepo struct {
	db *gorm.DB
}

// NewRoleRepository will initialize Role repository.
func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

// RoleRepo implements the repository.RoleRepository interface.
var _ repository.RoleRepository = &RoleRepo{}

// SaveRole will create a new role.
func (r RoleRepo) SaveRole(role *entity.Role) (*entity.Role, map[string]string) {
	panic("implement me")
}

// UpdateRole will update specified role.
func (r RoleRepo) UpdateRole(role *entity.Role) (*entity.Role, map[string]string) {
	panic("implement me")
}

// DeleteRole will delete role.
func (r RoleRepo) DeleteRole(role *entity.Role) error {
	panic("implement me")
}

// GetRole will return a role.
func (r RoleRepo) GetRole(uuid string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.Where("uuid = ?", uuid).Take(&role).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, exception.ErrorTextRoleNotFound
	}
	return &role, nil
}

// GetRolePermissions will return user roles.
func (r *RoleRepo) GetRolePermissions(uuid string) ([]entity.RolePermission, error) {
	var permission []entity.RolePermission
	err := r.db.Preload("Permission").Where("role_uuid = ?", uuid).Find(&permission).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, exception.ErrorTextUserNotFound
	}
	return permission, nil
}

// GetRoles will return role list.
func (r RoleRepo) GetRoles(c *gin.Context) ([]entity.Role, interface{}, error) {
	var total int
	var roles []entity.Role
	parameters := repository.NewParameters(c)
	errTotal := r.db.Find(&roles).Count(&total).Error
	errList := r.db.Limit(parameters.Limit).Offset(parameters.Offset).Find(&roles).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if gorm.IsRecordNotFoundError(errList) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(parameters, total)
	return roles, meta, nil
}
