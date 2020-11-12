package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

// createApplicationOauth will create predefined ApplicationOauth and insert into DB.
func createApplicationOauth(db *gorm.DB, application *entity.Application, applicationOauth *entity.ApplicationOauth) (
	*entity.ApplicationOauth, error) {
	var applicationExists entity.Application
	var applicationOauthExists entity.ApplicationOauth
	errApplication := db.Where("name = ?", application.Name).Take(&applicationExists).Error
	if errApplication != nil {
		return applicationOauth, nil
	}

	applicationOauth.ApplicationUUID = application.UUID

	errApplicationOauth := db.Where("application_uuid = ?", applicationExists.UUID).Take(&applicationOauthExists).Error
	if errApplicationOauth != nil {
		if errors.Is(errApplicationOauth, gorm.ErrRecordNotFound) {
			err := db.Create(applicationOauth).Error
			if err != nil {
				return applicationOauth, err
			}
			return applicationOauth, err
		}
		return applicationOauth, errApplicationOauth
	}
	return applicationOauth, nil
}
