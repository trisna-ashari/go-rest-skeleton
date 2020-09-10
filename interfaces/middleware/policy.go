package middleware

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Policy represent required dependencies.
type Policy struct {
	au repository.UserRepository
	ar repository.RoleRepository
}

var _ PolicyInterface = &Policy{}

// PolicyInterface is an interface.
type PolicyInterface interface {
	Can(action string, c *gin.Context) bool
}

// NewPolicy is constructor will initialize policy.
func NewPolicy(
	au repository.UserRepository,
	ar repository.RoleRepository) *Policy {
	return &Policy{
		au: au,
		ar: ar,
	}
}

// Can is a function uses to verify access.
func (p *Policy) Can(action string, c *gin.Context) bool {
	UUID, exists := c.Get("UUID")
	if !exists {
		c.Set("ErrorTracingCode", exception.ErrorCodeITMIPO001)
		_ = c.AbortWithError(http.StatusUnauthorized, exception.ErrorTextUnauthorized)
		return false
	}

	userRoles, _ := p.au.GetUserRoles(UUID.(string))
	var hasPermission bool
	var role entity.Role
	hasPermission = false
	for _, userRole := range userRoles {
		role = userRole.Role
		rolePermissions, _ := p.ar.GetRolePermissions(role.UUID)
		var permission entity.Permission
		for _, rolePermission := range rolePermissions {
			permission = rolePermission.Permission
			if permission.PermissionKey == action {
				hasPermission = true
				return hasPermission
			}
		}
	}
	if !hasPermission {
		c.Set("errorTracingCode", exception.ErrorCodeITMIPO002)
	}
	return hasPermission
}
