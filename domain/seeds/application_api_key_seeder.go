package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

// createApplicationApiKey will create predefined ApplicationApiKey and insert into DB.
func createApplicationApiKey(db *gorm.DB, application *entity.Application, applicationApiKey *entity.ApplicationApiKey) (
	*entity.ApplicationApiKey, error) {
	var applicationExists entity.Application
	var applicationApiKeyExists entity.ApplicationApiKey
	errApplication := db.Where("name = ?", application.Name).Take(&applicationExists).Error
	if errApplication != nil {
		return applicationApiKey, nil
	}

	applicationApiKey.ApplicationUUID = application.UUID

	errApplicationApiKey := db.Where("application_uuid = ?", applicationExists.UUID).Take(&applicationApiKeyExists).Error
	if errApplicationApiKey != nil {
		if errors.Is(errApplicationApiKey, gorm.ErrRecordNotFound) {
			err := db.Create(applicationApiKey).Error
			if err != nil {
				return applicationApiKey, err
			}
			return applicationApiKey, err
		}
		return applicationApiKey, errApplicationApiKey
	}
	return applicationApiKey, nil
}
