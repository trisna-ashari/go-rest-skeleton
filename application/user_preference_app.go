package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type userPreferenceApp struct {
	up repository.UserPreferenceRepository
}

// userPreferenceApp implement the UserPreferenceAppInterface.
var _ UserPreferenceAppInterface = &userPreferenceApp{}

// UserPreferenceAppInterface is interface.
type UserPreferenceAppInterface interface {
	GetUserPreference(UUID string) (*entity.UserPreference, error)
	UpdateUserPreference(UUID string, p *entity.DetailUserPreference) (*entity.UserPreference, map[string]string, error)
	ResetUserPreference(UUID string) (*entity.UserPreference, error)
}

// GetUserPreference is an implementation of method GetUserPreference.
func (up userPreferenceApp) GetUserPreference(userUUID string) (*entity.UserPreference, error) {
	return up.up.GetUserPreference(userUUID)
}

// UpdateUserPreference is an implementation of method UpdateUserPreference.
func (up userPreferenceApp) UpdateUserPreference(userUUID string, p *entity.DetailUserPreference) (
	*entity.UserPreference,
	map[string]string,
	error) {
	return up.up.UpdateUserPreference(userUUID, p)
}

// ResetUserPreference is an implementation of method ResetUserPreference.
func (up userPreferenceApp) ResetUserPreference(userUUID string) (*entity.UserPreference, error) {
	return up.up.ResetUserPreference(userUUID)
}
