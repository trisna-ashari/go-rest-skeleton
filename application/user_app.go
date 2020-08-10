package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

// UserAppInterface is an interface.
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string, error)
	UpdateUser(string, *entity.User) (*entity.User, map[string]string, error)
	DeleteUser(UUID string) error
	GetUsers(parameters *repository.Parameters) ([]entity.User, interface{}, error)
	GetUser(UUID string) (*entity.User, error)
	GetUserRoles(UUID string) ([]entity.UserRole, error)
	GetUserWithRoles(UUID string) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string, error)
	UpdateUserAvatar(string, *entity.User) (*entity.User, map[string]string, error)
}

// SaveUser is implementation of method SaveUser.
func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.SaveUser(user)
}

// UpdateUser is implementation of method SaveUser.
func (u *userApp) UpdateUser(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.UpdateUser(UUID, user)
}

// DeleteUser is implementation of method DeleteUser.
func (u *userApp) DeleteUser(UUID string) error {
	return u.us.DeleteUser(UUID)
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
func (u *userApp) GetUsers(p *repository.Parameters) ([]entity.User, interface{}, error) {
	return u.us.GetUsers(p)
}

// GetUserByEmailAndPassword is implementation of method GetUserByEmailAndPassword.
func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.GetUserByEmailAndPassword(user)
}

func (u *userApp) UpdateUserAvatar(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.UpdateUserAvatar(UUID, user)
}
