package service

import (
	"go-rest-skeleton/application"
	"go-rest-skeleton/domain/entity"
)

type AuthService struct {
	us application.UserAuthAppInterface
	uf application.UserForgotPasswordAppInterface
}

func (a AuthService) UpdateUser(UUID string, user *entity.User) (*entity.User, map[string]string, error) {
	return a.us.UpdateUser(UUID, user)
}

func (a AuthService) GetUser(UUID string) (*entity.User, error) {
	return a.us.GetUser(UUID)
}

func (a AuthService) CreateToken(user *entity.User) (*entity.UserForgotPassword, map[string]string, error) {
	return a.uf.CreateToken(user)
}

func (a AuthService) DeactivateToken(UUID string) error {
	return a.uf.DeactivateToken(UUID)
}

func (a AuthService) GetUserWithRoles(UUID string) (*entity.User, error) {
	return a.us.GetUserWithRoles(UUID)
}

func (a AuthService) GetUserByEmail(user *entity.User) (*entity.User, map[string]string, error) {
	return a.us.GetUserByEmail(user)
}

func (a AuthService) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string, error) {
	return a.us.GetUserByEmailAndPassword(user)
}

func (a AuthService) GetToken(Token string) (*entity.UserForgotPassword, error) {
	return a.uf.GetToken(Token)
}

var _ AuthServiceInterface = &AuthService{}

type AuthServiceInterface interface {
	application.UserAuthAppInterface
	application.UserForgotPasswordAppInterface
}

func NewAuthService(us application.UserAuthAppInterface, uf application.UserForgotPasswordAppInterface) AuthService {
	return AuthService{
		us: us,
		uf: uf,
	}
}
