package seeds

import (
	"go-rest-skeleton/domain/entity"

	"github.com/jinzhu/gorm"
)

// CreateRole will create predefined role and insert into DB.
func CreateRole(db *gorm.DB, UUID string, name string) error {
	return db.Create(&entity.Role{
		UUID: UUID,
		Name: name,
	}).Error
}
