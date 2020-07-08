package seeds

import (
	"github.com/jinzhu/gorm"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/security"
)

func CreateUser(db *gorm.DB, UUID string, firstName string, lastName string, email string, phone string, password string) error {
	hashedPassword, _ := security.Hash("123456")
	return db.Create(&entity.User{
		UUID:      UUID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Password:  string(hashedPassword),
	}).Error
}
