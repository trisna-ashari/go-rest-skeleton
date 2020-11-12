package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ApplicationApiKey represent schema of table modules.
type ApplicationApiKey struct {
	UUID            string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	ApplicationUUID string    `gorm:"size:36;not null;index;" json:"application_uuid"`
	Name            string    `gorm:"size:255;not null;index;" json:"name"`
	ApiKey          string    `gorm:"size:255;not null;index;" json:"api_key"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt
}

// ApplicationApiKeys represent multiple ApplicationApiKey.
type ApplicationApiKeys []ApplicationApiKey

// TableName return name of table.
func (aak *ApplicationApiKey) TableName() string {
	return "application_api_keys"
}

// BeforeCreate handle uuid generation.
func (aak *ApplicationApiKey) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if aak.UUID == "" {
		aak.UUID = generateUUID.String()
	}
	return nil
}
