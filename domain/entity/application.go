package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Application represent schema of table modules.
type Application struct {
	UUID                    string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	Name                    string    `gorm:"size:255;not null;index;" json:"name"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               gorm.DeletedAt
	ApplicationApiKeys      []ApplicationApiKey      `gorm:"foreignKey:ApplicationUUID"`
	ApplicationOauth        ApplicationOauth         `gorm:"foreignKey:ApplicationUUID"`
	ApplicationOauthClients []ApplicationOauthClient `gorm:"foreignKey:ApplicationUUID"`
}

// Applications represent multiple Application.
type Applications []Application

// TableName return name of table.
func (a *Application) TableName() string {
	return "applications"
}

// BeforeCreate handle uuid generation.
func (a *Application) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if a.UUID == "" {
		a.UUID = generateUUID.String()
	}
	return nil
}
