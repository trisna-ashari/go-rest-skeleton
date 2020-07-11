package repository

import (
	"go-rest-skeleton/domain/entity"

	"github.com/gin-gonic/gin"
)

// PermissionRepository is an interface.
type PermissionRepository interface {
	SavePermission(*entity.Permission) (*entity.Permission, map[string]string)
	UpdatePermission(*entity.Permission) (*entity.Permission, map[string]string)
	DeletePermission(*entity.Permission) error
	GetPermission(string) (*entity.Permission, error)
	GetPermissions(c *gin.Context) ([]entity.Permission, interface{}, error)
}
