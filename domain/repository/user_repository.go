package repository

import (
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(string) (*entity.User, error)
	GetUsers(c *gin.Context) ([]entity.User, interface{}, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}