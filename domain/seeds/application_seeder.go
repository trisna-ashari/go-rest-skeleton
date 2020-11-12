package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

// createApplication will create predefined application and insert into DB.
func createApplication(db *gorm.DB, application *entity.Application) (*entity.Application, error) {
	var applicationExists entity.Application
	err := db.Where("name = ?", application.Name).Take(&applicationExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(application).Error
			if err != nil {
				return application, err
			}
			return application, err
		}
		return application, err
	}
	return application, err
}
