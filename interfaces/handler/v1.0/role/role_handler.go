package role

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Roles is a struct defines the dependencies that will be used.
type Roles struct {
	ur application.RoleAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
}

// NewRoles is constructor will initialize role handler.
func NewRoles(
	ur application.RoleAppInterface,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface) *Roles {
	return &Roles{
		ur: ur,
		rd: rd,
		tk: tk,
	}
}

// GetRoles is a function uses to handle get role list.
func (s *Roles) GetRoles(c *gin.Context) {
	var roles entity.Roles
	var err error
	roles, meta, err := s.ur.GetRoles(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, roles.DetailRoles(), "api.msg.success.successfully_get_user_list", meta)
}

// GetRole is a function uses to handle get role detail by UUID.
func (s *Roles) GetRole(c *gin.Context) {
	var roleEntity entity.Role
	if err := c.ShouldBindUri(&roleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	role, err := s.ur.GetRole(UUID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, role.DetailRole(), "api.msg.success.successfully_get_user_detail", nil)
}
