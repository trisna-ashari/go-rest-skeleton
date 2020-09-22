package persistence

import (
	"errors"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	// Expiry of token for reset password.
	ExpiryToken = time.Hour * 24 * 7 * time.Duration(2)
)

// UserForgotPasswordRepo is a struct to store db connection.
type UserForgotPasswordRepo struct {
	db *gorm.DB
}

// NewUserForgotPasswordRepository will initialize persistence.UserForgotPasswordRepo repository.
func NewUserForgotPasswordRepository(db *gorm.DB) *UserForgotPasswordRepo {
	return &UserForgotPasswordRepo{db}
}

// UserForgotPasswordRepo implements the repository.UserForgotPasswordRepository interface.
var _ repository.UserForgotPasswordRepository = &UserForgotPasswordRepo{}

// CreateToken will create new token.
func (u *UserForgotPasswordRepo) CreateToken(user *entity.User) (*entity.UserForgotPassword, map[string]string, error) {
	var userForgotPassword entity.UserForgotPassword
	errDesc := map[string]string{}

	err := u.db.Where("user_uuid = ?", user.UUID).Delete(entity.UserForgotPassword{}).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}

	userForgotPassword.UserUUID = user.UUID
	userForgotPassword.Password = user.Password
	userForgotPassword.Token = uuid.New().String()
	userForgotPassword.ExpiredAt = time.Now().Local().Add(ExpiryToken)
	err = u.db.Create(&userForgotPassword).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return &userForgotPassword, nil, nil
}

// DeactivateToken will deactivate token.
func (u *UserForgotPasswordRepo) DeactivateToken(UUID string) error {
	var userForgotPassword entity.UserForgotPassword
	err := u.db.Where("uuid = ?", UUID).Take(&userForgotPassword).Delete(&userForgotPassword).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextUserForgotPasswordTokenNotFound
		}
		return err
	}
	return nil
}

// GetToken will return token detail.
func (u *UserForgotPasswordRepo) GetToken(token string) (*entity.UserForgotPassword, error) {
	var userForgotPassword entity.UserForgotPassword
	err := u.db.Where("token = ?", token).Take(&userForgotPassword).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextUserForgotPasswordTokenNotFound
		}
		return nil, err
	}
	return &userForgotPassword, nil
}
