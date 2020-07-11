package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"

	"github.com/gin-gonic/gin"
)

type roleApp struct {
	us repository.RoleRepository
}

var _ RoleAppInterface = &roleApp{}

// RoleAppInterface is an interface.
type RoleAppInterface interface {
	SaveRole(*entity.Role) (*entity.Role, map[string]string)
	UpdateRole(*entity.Role) (*entity.Role, map[string]string)
	DeleteRole(*entity.Role) error
	GetRoles(c *gin.Context) ([]entity.Role, interface{}, error)
	GetRole(UUID string) (*entity.Role, error)
	GetRolePermissions(UUID string) ([]entity.RolePermission, error)
}

// SaveRole is implementation of method SaveRole.
func (u *roleApp) SaveRole(Role *entity.Role) (*entity.Role, map[string]string) {
	return u.us.SaveRole(Role)
}

// UpdateRole is implementation of method UpdateRole.
func (u *roleApp) UpdateRole(Role *entity.Role) (*entity.Role, map[string]string) {
	return u.us.SaveRole(Role)
}

// DeleteRole is implementation of method SaveRole.
func (u *roleApp) DeleteRole(Role *entity.Role) error {
	return u.us.DeleteRole(Role)
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
func (u *roleApp) GetRoles(c *gin.Context) ([]entity.Role, interface{}, error) {
	return u.us.GetRoles(c)
}
