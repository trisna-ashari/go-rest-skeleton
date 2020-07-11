package repository

import (
	"go-rest-skeleton/domain/entity"

	"github.com/gin-gonic/gin"
)

// RoleRepository is an interface.
type RoleRepository interface {
	SaveRole(*entity.Role) (*entity.Role, map[string]string)
	UpdateRole(*entity.Role) (*entity.Role, map[string]string)
	DeleteRole(*entity.Role) error
	GetRole(string) (*entity.Role, error)
	GetRolePermissions(string) ([]entity.RolePermission, error)
	GetRoles(c *gin.Context) ([]entity.Role, interface{}, error)
}
