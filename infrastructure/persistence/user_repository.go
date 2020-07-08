package persistence

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/security"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

func (r *UserRepo) GetUser(uuid string) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("uuid = ?", uuid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers(c *gin.Context) ([]entity.User, interface{}, error) {
	var total int
	var users []entity.User
	parameters := repository.NewParameters(c)
	errTotal := r.db.Debug().Find(&users).Count(&total).Error
	errList := r.db.Debug().Limit(parameters.Limit).Offset(parameters.Offset).Find(&users).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if gorm.IsRecordNotFoundError(errList) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(parameters, total)
	return users, meta, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["email"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}

func (r *UserRepo) Greeting(c *gin.Context) (string, error) {
	return "", nil
}