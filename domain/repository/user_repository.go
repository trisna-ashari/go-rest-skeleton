package repository

import (
	"go-rest-skeleton/domain/entity"
)

// UserRepository is an interface.
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string, error)
	UpdateUser(string, *entity.User) (*entity.User, map[string]string, error)
	DeleteUser(string) error
	GetUser(string) (*entity.User, error)
	GetUserRoles(string) ([]entity.UserRole, error)
	GetUserWithRoles(string) (*entity.User, error)
	GetUsers(parameters *Parameters) ([]entity.User, interface{}, error)
	GetUserByEmail(*entity.User) (*entity.User, map[string]string, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string, error)
	UpdateUserAvatar(string, *entity.User) (*entity.User, map[string]string, error)
}
