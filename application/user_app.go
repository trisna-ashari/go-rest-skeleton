package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"

	"github.com/gin-gonic/gin"
)

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

// UserAppInterface is an interface.
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers(c *gin.Context) ([]entity.User, interface{}, error)
	GetUser(UUID string) (*entity.User, error)
	GetUserRoles(UUID string) ([]entity.UserRole, error)
	GetUserWithRoles(UUID string) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

// SaveUser is implementation of method SaveUser.
func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

// GetUser is implementation of method GetUser.
func (u *userApp) GetUser(UUID string) (*entity.User, error) {
	return u.us.GetUser(UUID)
}

// GetUserRoles is implementation of method GetUserWithRoles.
func (u *userApp) GetUserRoles(UUID string) ([]entity.UserRole, error) {
	return u.us.GetUserRoles(UUID)
}

// GetUserWithRoles is implementation of method GetUserWithRoles.
func (u *userApp) GetUserWithRoles(UUID string) (*entity.User, error) {
	return u.us.GetUser(UUID)
}

// GetUsers is implementation of method GetUsers.
func (u *userApp) GetUsers(c *gin.Context) ([]entity.User, interface{}, error) {
	return u.us.GetUsers(c)
}

// GetUserByEmailAndPassword is implementation of method GetUserByEmailAndPassword.
func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}
