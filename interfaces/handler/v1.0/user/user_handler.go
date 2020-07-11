package userv1point00

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/authorization"
	"go-rest-skeleton/infrastructure/exception"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const moduleName = "user"

// Users is a struct defines the dependencies that will be used.
type Users struct {
	rs application.RoleAppInterface
	us application.UserAppInterface
	rd authorization.AuthInterface
	tk authorization.TokenInterface
	pl middleware.PolicyInterface
}

// NewUsers is constructor will initialize user handler.
func NewUsers(
	ds *persistence.Repositories,
	rd authorization.AuthInterface,
	tk authorization.TokenInterface) *Users {
	return &Users{
		rs: ds.Role,
		us: ds.User,
		rd: rd,
		tk: tk,
		pl: middleware.NewPolicy(ds.User, ds.Role, tk),
	}
}

// SaveUser is a function uses to handle create a new user.
func (s *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	validateErr := user.ValidateSaveUser(c)
	if len(validateErr) > 0 {
		c.Set("data", validateErr)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	middleware.Formatter(c, newUser.DetailUser(), "api.msg.success.successfully_create_a_new_user", nil)
}

// GetUsers is a function uses to handle get user list.
func (s *Users) GetUsers(c *gin.Context) {
	if !s.pl.Can(moduleName, "read", c) {
		_ = c.AbortWithError(http.StatusForbidden, exception.ErrorTextForbidden)
		return
	}
	var users entity.Users
	var err error
	users, meta, err := s.us.GetUsers(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, users.DetailUsers(), "api.msg.success.successfully_get_user_list", meta)
}

// GetUser is a function uses to handle get user detail by UUID.
func (s *Users) GetUser(c *gin.Context) {
	var userEntity entity.User
	if err := c.ShouldBindUri(&userEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	user, err := s.us.GetUser(UUID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextUserNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	middleware.Formatter(c, user.DetailUser(), "api.msg.success.successfully_get_user_detail", nil)
}
