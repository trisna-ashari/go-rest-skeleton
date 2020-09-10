package repository

import "go-rest-skeleton/domain/entity"

// UserPreferenceRepository is an interface.
type UserPreferenceRepository interface {
	GetUserPreference(string) (*entity.UserPreference, error)
	UpdateUserPreference(string, *entity.DetailUserPreference) (*entity.UserPreference, map[string]string, error)
	ResetUserPreference(string) (*entity.UserPreference, error)
}
