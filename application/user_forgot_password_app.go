package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type userForgotPasswordApp struct {
	uf repository.UserForgotPasswordRepository
}

// userForgotPasswordApp implement the userForgotPasswordAppInterface.
var _ UserForgotPasswordAppInterface = &userForgotPasswordApp{}

type UserForgotPasswordAppInterface interface {
	CreateToken(user *entity.User) (*entity.UserForgotPassword, map[string]string, error)
	DeactivateToken(UUID string) error
	GetToken(Token string) (*entity.UserForgotPassword, error)
}

func (u userForgotPasswordApp) CreateToken(user *entity.User) (*entity.UserForgotPassword, map[string]string, error) {
	return u.uf.CreateToken(user)
}

func (u userForgotPasswordApp) DeactivateToken(UUID string) error {
	return u.uf.DeactivateToken(UUID)
}

func (u userForgotPasswordApp) GetToken(Token string) (*entity.UserForgotPassword, error) {
	return u.uf.GetToken(Token)
}
