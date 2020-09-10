package entity

import (
	"encoding/json"

	"github.com/google/uuid"
)

// UserPreference represent schema of table user_preferences.
type UserPreference struct {
	UUID       string           `gorm:"size:36;not null;unique_index;primary_key;" json:"uuid"`
	UserUUID   string           `gorm:"size:36;not null;unique_index;" json:"user_uuid"`
	Preference *json.RawMessage `gorm:"type:text;not null;" json:"preference"`
}

// BeforeSave handle uuid generation.
func (up *UserPreference) BeforeSave() error {
	generateUUID := uuid.New()
	if up.UUID == "" {
		up.UUID = generateUUID.String()
	}
	return nil
}

// DetailUserPreference represent format of detail user preference.
type DetailUserPreference struct {
	Language string `json:"language"`
	DarkMode bool   `json:"dark_mode"`
	SkipTour bool   `json:"skip_tour"`
}

// DetailUserPreference will return user preference.
func (up *UserPreference) DetailUserPreference() interface{} {
	return up.Preference
}

// BuildDefaultPreference will return default preference.
func (up *UserPreference) BuildDefaultPreference() *json.RawMessage {
	defaultUserPreference, err := json.Marshal(
		&DetailUserPreference{
			Language: "en",
			DarkMode: false,
			SkipTour: false,
		},
	)
	if err != nil {
		return nil
	}

	return (*json.RawMessage)(&defaultUserPreference)
}

// BuildPatchPreference will return preference based on current.
func (up *UserPreference) BuildPatchPreference(preference *DetailUserPreference) *json.RawMessage {
	patchUserPreference, err := json.Marshal(
		&DetailUserPreference{
			Language: preference.Language,
			DarkMode: preference.DarkMode,
			SkipTour: preference.SkipTour,
		},
	)
	if err != nil {
		return nil
	}

	return (*json.RawMessage)(&patchUserPreference)
}

// BuildPatchUserPreference will return user preference based on current preference.
func (up *UserPreference) BuildPatchUserPreference(preference *json.RawMessage) *UserPreference {
	var preferenceDetail DetailUserPreference
	_ = json.Unmarshal(*preference, &preferenceDetail)

	return &UserPreference{
		UUID:       up.UUID,
		UserUUID:   up.UserUUID,
		Preference: up.BuildPatchPreference(&preferenceDetail),
	}
}
