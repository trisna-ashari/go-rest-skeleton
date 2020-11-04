package persistence

import (
	"errors"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"

	"gorm.io/gorm"
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
func (r RoleRepo) SaveRole(role *entity.Role) (*entity.Role, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&role).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return role, nil, nil
}

// UpdateRole will update specified role.
func (r RoleRepo) UpdateRole(uuid string, role *entity.Role) (*entity.Role, map[string]string, error) {
	errDesc := map[string]string{}
	roleData := &entity.Role{
		Name: role.Name,
	}
	err := r.db.First(&role, "uuid = ?", uuid).Updates(roleData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextRoleInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextRoleNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return role, nil, nil
}

// DeleteRole will delete role.
func (r RoleRepo) DeleteRole(uuid string) error {
	var role entity.Role
	err := r.db.Where("uuid = ?", uuid).Take(&role).Delete(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextUserNotFound
		}
		return err
	}
	return nil
}

// GetRole will return a role.
func (r RoleRepo) GetRole(uuid string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.Where("uuid = ?", uuid).Take(&role).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrorTextUserNotFound
	}
	return permission, nil
}

// GetRoleWithPermissions will return user detail with roles.
func (r *RoleRepo) GetRoleWithPermissions(uuid string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.Preload("RolePermissions.Permission").Where("uuid = ?", uuid).Take(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserNotFound
		}
		return nil, err
	}
	return &role, nil
}

// GetRoles will return role list.
func (r RoleRepo) GetRoles(p *repository.Parameters) ([]entity.Role, interface{}, error) {
	var total int64
	var roles []entity.Role
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&roles).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&roles).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if errors.Is(errList, gorm.ErrRecordNotFound) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(p, total)
	return roles, meta, nil
}
