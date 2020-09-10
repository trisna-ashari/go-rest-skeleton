package entity

import "github.com/google/uuid"

// UserLogin represent schema of table user_login.
type UserLogin struct {
	UUID      string `gorm:"size:36;not null;unique_index;primary_key;" json:"uuid"`
	UserUUID  string `gorm:"size:36;not null;index:user_uuid;" json:"user_uuid"`
	UserAgent string `gorm:"size:255;" json:"user_agent"`
	Platform  string `gorm:"size:32;not null;index:platform;" json:"platform"`
}

// BeforeSave handle uuid generation.
func (ul *UserLogin) BeforeSave() error {
	generateUUID := uuid.New()
	if ul.UUID == "" {
		ul.UUID = generateUUID.String()
	}
	return nil
}
