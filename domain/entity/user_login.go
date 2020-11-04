package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserLogin represent schema of table user_login.
type UserLogin struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	UserUUID  string    `gorm:"size:36;not null;index:user_uuid;" json:"user_uuid"`
	UserAgent string    `gorm:"size:255;" json:"user_agent"`
	Platform  string    `gorm:"size:32;not null;index:platform;" json:"platform"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (ul *UserLogin) TableName() string {
	return "user_login"
}

// BeforeCreate handle uuid generation.
func (ul *UserLogin) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if ul.UUID == "" {
		ul.UUID = generateUUID.String()
	}
	return nil
}
