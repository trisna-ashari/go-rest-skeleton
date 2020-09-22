package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

type userAuthApp struct {
	us repository.UserRepository
}

// userApp implement the UserAppInterface.
var _ UserAppInterface = &userApp{}

// userAuthApp implement the UserAuthAppInterface.
var _ UserAuthAppInterface = &userApp{}

// UserAppInterface is an interface.
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string, error)
	UpdateUser(string, *entity.User) (*entity.User, map[string]string, error)
	DeleteUser(UUID string) error
	GetUsers(parameters *repository.Parameters) ([]entity.User, interface{}, error)
	GetUser(UUID string) (*entity.User, error)
	GetUserRoles(UUID string) ([]entity.UserRole, error)
	GetUserWithRoles(UUID string) (*entity.User, error)
	GetUserByEmail(*entity.User) (*entity.User, map[string]string, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string, error)
	UpdateUserAvatar(string, *entity.User) (*entity.User, map[string]string, error)
}

type UserAuthAppInterface interface {
	UpdateUser(string, *entity.User) (*entity.User, map[string]string, error)
	GetUser(UUID string) (*entity.User, error)
	GetUserWithRoles(UUID string) (*entity.User, error)
	GetUserByEmail(*entity.User) (*entity.User, map[string]string, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string, error)
}

// SaveUser is an implementation of method SaveUser.
func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.SaveUser(user)
}

// UpdateUser is an implementation of method SaveUser.
func (u *userApp) UpdateUser(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.UpdateUser(UUID, user)
}

// DeleteUser is an implementation of method DeleteUser.
func (u *userApp) DeleteUser(UUID string) error {
	return u.us.DeleteUser(UUID)
}

// GetUser is an implementation of method GetUser.
func (u *userApp) GetUser(UUID string) (*entity.User, error) {
	return u.us.GetUser(UUID)
}

// GetUserRoles is an implementation of method GetUserWithRoles.
func (u *userApp) GetUserRoles(UUID string) ([]entity.UserRole, error) {
	return u.us.GetUserRoles(UUID)
}

// GetUserWithRoles is an implementation of method GetUserWithRoles.
func (u *userApp) GetUserWithRoles(UUID string) (*entity.User, error) {
	return u.us.GetUser(UUID)
}

// GetUsers is an implementation of method GetUsers.
func (u *userApp) GetUsers(p *repository.Parameters) ([]entity.User, interface{}, error) {
	return u.us.GetUsers(p)
}

// GetUserByEmail is an implementation of method GetUserByEmail.
func (u *userApp) GetUserByEmail(user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.GetUserByEmailAndPassword(user)
}

// GetUserByEmailAndPassword is an implementation of method GetUserByEmailAndPassword.
func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.GetUserByEmailAndPassword(user)
}

func (u *userApp) UpdateUserAvatar(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.us.UpdateUserAvatar(UUID, user)
}
