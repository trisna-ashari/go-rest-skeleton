package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

// createUser will create fake user and insert into DB.
func createUser(db *gorm.DB, user *entity.User) (*entity.User, error) {
	err := db.Create(user).Error
	return user, err
}
