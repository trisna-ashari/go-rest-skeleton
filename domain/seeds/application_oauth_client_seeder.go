package seeds

import (
	"errors"
	"go-rest-skeleton/domain/entity"

	"gorm.io/gorm"
)

// createApplicationOauthClient will create predefined ApplicationOauthClient and insert into DB.
func createApplicationOauthClient(
	db *gorm.DB,
	application *entity.Application,
	applicationOauth *entity.ApplicationOauthClient) (*entity.ApplicationOauthClient, error) {
	var applicationExists entity.Application
	var applicationOauthExists entity.ApplicationOauthClient
	errApplication := db.Where("name = ?", application.Name).Take(&applicationExists).Error
	if errApplication != nil {
		return applicationOauth, nil
	}

	applicationOauth.ApplicationUUID = application.UUID

	errApplicationOauthClient := db.Where("application_uuid = ?", applicationExists.UUID).Take(&applicationOauthExists).Error
	if errApplicationOauthClient != nil {
		if errors.Is(errApplicationOauthClient, gorm.ErrRecordNotFound) {
			err := db.Create(applicationOauth).Error
			if err != nil {
				return applicationOauth, err
			}
			return applicationOauth, err
		}
		return applicationOauth, errApplicationOauthClient
	}
	return applicationOauth, nil
}
