package persistence

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PermissionRepo is a struct to store db connection.
type PermissionRepo struct {
	db *gorm.DB
}

// NewPermissionRepository will initialize Permission repository.
func NewPermissionRepository(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db}
}

// PermissionRepo implements the repository.PermissionRepository interface.
var _ repository.PermissionRepository = &PermissionRepo{}

// SavePermission will create a new permission.
func (p PermissionRepo) SavePermission(permission *entity.Permission) (*entity.Permission, map[string]string) {
	panic("implement me")
}

// UpdatePermission will update specified permission.
func (p PermissionRepo) UpdatePermission(permission *entity.Permission) (*entity.Permission, map[string]string) {
	panic("implement me")
}

// DeletePermission will delete specified permission.
func (p PermissionRepo) DeletePermission(permission *entity.Permission) error {
	panic("implement me")
}

// GetPermission will return a permission.
func (p PermissionRepo) GetPermission(s string) (*entity.Permission, error) {
	panic("implement me")
}

// GetPermissions will return a permission list.
func (p PermissionRepo) GetPermissions(c *gin.Context) ([]entity.Permission, interface{}, error) {
	panic("implement me")
}
