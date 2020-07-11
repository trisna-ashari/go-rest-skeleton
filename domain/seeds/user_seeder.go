package seeds

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/security"

	"github.com/jinzhu/gorm"
)

// CreateUser will create fake user and insert into DB.
func CreateUser(
	db *gorm.DB,
	UUID string,
	firstName string,
	lastName string,
	email string,
	phone string,
	password string) error {
	hashedPassword, _ := security.Hash(password)

	return db.Create(&entity.User{
		UUID:      UUID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Password:  string(hashedPassword),
	}).Error
}
