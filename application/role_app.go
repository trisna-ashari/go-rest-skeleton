package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type roleApp struct {
	us repository.RoleRepository
}

var _ RoleAppInterface = &roleApp{}

// RoleAppInterface is an interface.
type RoleAppInterface interface {
	SaveRole(*entity.Role) (*entity.Role, map[string]string, error)
	UpdateRole(string, *entity.Role) (*entity.Role, map[string]string, error)
	DeleteRole(UUID string) error
	GetRoles(p *repository.Parameters) ([]entity.Role, interface{}, error)
	GetRole(UUID string) (*entity.Role, error)
	GetRolePermissions(UUID string) ([]entity.RolePermission, error)
}

// SaveRole is implementation of method SaveRole.
func (u *roleApp) SaveRole(Role *entity.Role) (*entity.Role, map[string]string, error) {
	return u.us.SaveRole(Role)
}

// UpdateRole is implementation of method UpdateRole.
func (u *roleApp) UpdateRole(UUID string, Role *entity.Role) (*entity.Role, map[string]string, error) {
	return u.us.UpdateRole(UUID, Role)
}

// DeleteRole is implementation of method SaveRole.
func (u *roleApp) DeleteRole(UUID string) error {
	return u.us.DeleteRole(UUID)
}

// GetRole is implementation of method GetRole.
func (u *roleApp) GetRole(UUID string) (*entity.Role, error) {
	return u.us.GetRole(UUID)
}

// GetRolePermissions is implementation of method GetRolePermissions.
func (u *roleApp) GetRolePermissions(UUID string) ([]entity.RolePermission, error) {
	return u.us.GetRolePermissions(UUID)
}

// GetRoles is implementation of method GetRoles.
func (u *roleApp) GetRoles(p *repository.Parameters) ([]entity.Role, interface{}, error) {
	return u.us.GetRoles(p)
}
