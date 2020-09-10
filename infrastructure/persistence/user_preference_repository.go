package persistence

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"

	"github.com/jinzhu/gorm"
)

// UserPreferenceRepo is a struct to store db connection.
type UserPreferenceRepo struct {
	db *gorm.DB
}

// NewUserPreferenceRepository will initialize persistence.UserPreferenceRepo repository.
func NewUserPreferenceRepository(db *gorm.DB) *UserPreferenceRepo {
	return &UserPreferenceRepo{db}
}

// UserPreferenceRepo implements the repository.UserPreferenceRepository interface.
var _ repository.UserPreferenceRepository = &UserPreferenceRepo{}

func (up UserPreferenceRepo) GetUserPreference(userUUID string) (*entity.UserPreference, error) {
	var userPreference entity.UserPreference
	err := up.db.Where("user_uuid = ?", userUUID).Take(&userPreference).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		userPreference.UserUUID = userUUID
		userPreference.Preference = userPreference.BuildDefaultPreference()
		err := up.db.Create(&userPreference).Error
		if err != nil {
			return nil, exception.ErrorTextAnErrorOccurred
		}
		return &userPreference, nil
	}

	userPreferencePatch := userPreference.BuildPatchUserPreference(userPreference.Preference)
	err = up.db.First(&userPreference, "user_uuid = ?", userUUID).Updates(userPreferencePatch).Error
	if err != nil {
		return &userPreference, err
	}

	return &userPreference, nil
}

func (up UserPreferenceRepo) UpdateUserPreference(userUUID string, p *entity.DetailUserPreference) (
	*entity.UserPreference,
	map[string]string,
	error) {
	errDesc := map[string]string{}
	var userPreference entity.UserPreference
	userPreferenceData := &entity.UserPreference{
		Preference: userPreference.BuildPatchPreference(p),
	}
	err := up.db.First(&userPreference, "user_uuid = ?", userUUID).Updates(userPreferenceData).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		userPreference.UserUUID = userUUID
		err := up.db.Create(&userPreferenceData).Error
		if err != nil {
			return nil, errDesc, exception.ErrorTextAnErrorOccurred
		}
		return &userPreference, nil, nil
	}
	return &userPreference, nil, nil
}

func (up UserPreferenceRepo) ResetUserPreference(userUUID string) (*entity.UserPreference, error) {
	var userPreference entity.UserPreference
	userPreferenceData := &entity.UserPreference{
		Preference: userPreference.BuildDefaultPreference(),
	}
	err := up.db.First(&userPreference, "user_uuid = ?", userUUID).Updates(userPreferenceData).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		userPreference.UserUUID = userUUID
		err := up.db.Create(&userPreferenceData).Error
		if err != nil {
			return nil, exception.ErrorTextAnErrorOccurred
		}
		return &userPreference, nil
	}
	return &userPreference, nil
}
