package repository

import (
	"go-rest-skeleton/domain/entity"
)

// RoleRepository is an interface.
type RoleRepository interface {
	SaveRole(*entity.Role) (*entity.Role, map[string]string, error)
	UpdateRole(string, *entity.Role) (*entity.Role, map[string]string, error)
	DeleteRole(string) error
	GetRole(string) (*entity.Role, error)
	GetRolePermissions(string) ([]entity.RolePermission, error)
	GetRoleWithPermissions(string) (*entity.Role, error)
	GetRoles(parameters *Parameters) ([]entity.Role, interface{}, error)
}
