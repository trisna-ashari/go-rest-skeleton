package repository

import (
	"go-rest-skeleton/domain/entity"

	"github.com/gin-gonic/gin"
)

// UserRepository is an interface.
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(string) (*entity.User, error)
	GetUserRoles(string) ([]entity.UserRole, error)
	GetUserWithRoles(string) (*entity.User, error)
	GetUsers(c *gin.Context) ([]entity.User, interface{}, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}
