package repository

import "go-rest-skeleton/domain/entity"

// UserForgotPasswordRepository is an interface.
type UserForgotPasswordRepository interface {
	CreateToken(user *entity.User) (*entity.UserForgotPassword, map[string]string, error)
	DeactivateToken(UUID string) error
	GetToken(Token string) (*entity.UserForgotPassword, error)
}
