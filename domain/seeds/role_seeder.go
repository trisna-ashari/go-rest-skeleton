package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

// createRole will create predefined role and insert into DB.
func createRole(db *gorm.DB, UUID string, name string) error {
	return db.Create(&entity.Role{
		UUID: UUID,
		Name: name,
	}).Error
}
