package mock

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

// UserAppInterface is a mock of application.UserAppInterface.
type UserAppInterface struct {
	SaveUserFn                  func(*entity.User) (*entity.User, map[string]string, error)
	UpdateUserFn                func(string, *entity.User) (*entity.User, map[string]string, error)
	DeleteUserFn                func(UUID string) error
	GetUsersFn                  func(params *repository.Parameters) ([]entity.User, interface{}, error)
	GetUserFn                   func(UUID string) (*entity.User, error)
	GetUserRolesFn              func(UUID string) ([]entity.UserRole, error)
	GetUserWithRolesFn          func(UUID string) (*entity.User, error)
	GetUserByEmailFn            func(*entity.User) (*entity.User, map[string]string, error)
	GetUserByEmailAndPasswordFn func(*entity.User) (*entity.User, map[string]string, error)
	UpdateUserAvatarFn          func(string, *entity.User) (*entity.User, map[string]string, error)
}

// SaveUser calls the SaveUserFn.
func (u *UserAppInterface) SaveUser(user *entity.User) (*entity.User, map[string]string, error) {
	return u.SaveUserFn(user)
}

// UpdateUser calls the UpdateUserFn.
func (u *UserAppInterface) UpdateUser(uuid string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.UpdateUserFn(uuid, user)
}

// DeleteUser calls the DeleteUserFn.
func (u *UserAppInterface) DeleteUser(uuid string) error {
	return u.DeleteUserFn(uuid)
}

// GetUsers calls the GetUsersFn.
func (u *UserAppInterface) GetUsers(params *repository.Parameters) ([]entity.User, interface{}, error) {
	return u.GetUsersFn(params)
}

// GetUser calls the GetUserFn.
func (u *UserAppInterface) GetUser(uuid string) (*entity.User, error) {
	return u.GetUserFn(uuid)
}

// GetUserRoles calls the GetUserRolesFn.
func (u *UserAppInterface) GetUserRoles(uuid string) ([]entity.UserRole, error) {
	return u.GetUserRolesFn(uuid)
}

// GetUserWithRoles calls the GetUserWithRolesFn.
func (u *UserAppInterface) GetUserWithRoles(uuid string) (*entity.User, error) {
	return u.GetUserWithRolesFn(uuid)
}

// GetUserByEmail calls the GetUserByEmailFn.
func (u *UserAppInterface) GetUserByEmail(user *entity.User) (*entity.User, map[string]string, error) {
	return u.GetUserByEmailFn(user)
}

// GetUserByEmailAndPassword calls the GetUserByEmailAndPasswordFn.
func (u *UserAppInterface) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string, error) {
	return u.GetUserByEmailAndPasswordFn(user)
}

// UpdateUserAvatar calls the UpdateUserAvatarFn.
func (u *UserAppInterface) UpdateUserAvatar(uuid string, user *entity.User) (*entity.User, map[string]string, error) {
	return u.UpdateUserFn(uuid, user)
}
