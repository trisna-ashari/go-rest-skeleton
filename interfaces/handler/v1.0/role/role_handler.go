package role

import (
	"errors"
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/message/exception"
	"go-rest-skeleton/infrastructure/message/success"
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

// SaveRole is a function uses to handle create a new role.
func (s *Roles) SaveRole(c *gin.Context) {
	var roleEntity entity.Role
	if err := c.ShouldBindJSON(&roleEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateErr := roleEntity.ValidateSaveRole(c)
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newRole, errDesc, errException := s.ur.SaveRole(&roleEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	middleware.Formatter(c, newRole.DetailRole(), success.RoleSuccessfullyCreateRole, nil)
}

// UpdateUser is a function uses to handle create a new user.
func (s *Roles) UpdateRole(c *gin.Context) {
	var roleEntity entity.Role
	if err := c.ShouldBindUri(&roleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&roleEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := roleEntity.ValidateUpdateRole(c)
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	UUID := c.Param("uuid")
	updatedRole, errDesc, errException := s.ur.UpdateRole(UUID, &roleEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextRoleNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, errException)
			return
		}
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusOK)
	middleware.Formatter(c, updatedRole.DetailRole(), success.RoleSuccessfullyUpdateRole, nil)
}

// DeleteRole is a function uses to handle delete role by UUID.
func (s *Roles) DeleteRole(c *gin.Context) {
	var roleEntity entity.Role
	if err := c.ShouldBindUri(&roleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.ur.DeleteRole(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextUserNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, nil, success.RoleSuccessfullyDeleteRole, nil)
}

// GetRoles is a function uses to handle get role list.
func (s *Roles) GetRoles(c *gin.Context) {
	var roles entity.Roles
	var err error
	parameters := repository.NewParameters(c)
	roles, meta, err := s.ur.GetRoles(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, roles.DetailRoles(), success.RoleSuccessfullyGetRoleList, meta)
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
	middleware.Formatter(c, role.DetailRole(), success.RoleSuccessfullyGetRoleDetail, nil)
}
