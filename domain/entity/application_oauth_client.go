package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ApplicationOauthClient represent schema of table modules.
type ApplicationOauthClient struct {
	UUID            string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	ApplicationUUID string    `gorm:"size:36;not null;index;" json:"application_uuid"`
	Name            string    `gorm:"size:255;not null;index;" json:"name"`
	ClientID        string    `gorm:"size:255;not null;index;" json:"client_id"`
	ClientSecret    string    `gorm:"size:255;not null;index;" json:"client_secret"`
	Referrers       string    `gorm:"size:255;not null;" json:"referrers"`
	Callbacks       string    `gorm:"size:255;not null;" json:"callbacks"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt
}

// ApplicationOauthClients represent multiple ApplicationOauthClient.
type ApplicationOauthClients []ApplicationOauthClient

// TableName return name of table.
func (aoc *ApplicationOauthClient) TableName() string {
	return "application_oauth_clients"
}

// BeforeCreate handle uuid generation.
func (aoc *ApplicationOauthClient) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if aoc.UUID == "" {
		aoc.UUID = generateUUID.String()
	}
	return nil
}
