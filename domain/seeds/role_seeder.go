package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

// createRole will create predefined role and insert into DB.
func createRole(db *gorm.DB, role *entity.Role) (*entity.Role, error) {
	err := db.Create(role).Error
	return role, err
}
