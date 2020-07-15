package middleware

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Policy represent required dependencies.
type Policy struct {
	au application.UserAppInterface
	ar application.RoleAppInterface
	tk authorization.TokenInterface
}

var _ PolicyInterface = &Policy{}

// PolicyInterface is an interface.
type PolicyInterface interface {
	Can(module string, action string, c *gin.Context) bool
}

// NewPolicy is constructor will initialize policy.
func NewPolicy(
	au application.UserAppInterface,
	ar application.RoleAppInterface,
	tk authorization.TokenInterface) *Policy {
	return &Policy{
		au: au,
		ar: ar,
		tk: tk,
	}
}

// Can is a function uses to verify access.
func (p *Policy) Can(module string, action string, c *gin.Context) bool {
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
			if permission.ModuleKey == module && permission.PermissionKey == action {
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
