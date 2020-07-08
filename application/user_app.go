package application

import (
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers(c *gin.Context) ([]entity.User, interface{}, error)
	GetUser(UUID string) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(UUID string) (*entity.User, error) {
	return u.us.GetUser(UUID)
}

func (u *userApp) GetUsers(c *gin.Context) ([]entity.User, interface{}, error) {
	return u.us.GetUsers(c)
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}