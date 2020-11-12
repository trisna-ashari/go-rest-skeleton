package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationOauth struct {
	UUID             string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	ApplicationUUID  string    `gorm:"size:36;not null;index;" json:"application_uuid"`
	Name             string    `gorm:"size:255;not null;index;" json:"name"`
	SupportEmails    string    `gorm:"size:255;not null;" json:"support_emails"`
	DeveloperEmails  string    `gorm:"size:255;not null;" json:"developer_emails"`
	Logo             string    `gorm:"size:255;not null;" json:"logo"`
	HomePageURL      string    `gorm:"size:255;not null;" json:"homepage_url"`
	TosURL           string    `gorm:"size:255;not null;" json:"tos_url"`
	PrivacyPolicyURL string    `gorm:"size:255;not null;" json:"privacy_policy_url"`
	Domains          string    `gorm:"size:255;not null;" json:"domains"`
	Scopes           string    `gorm:"size:255;not null;" json:"scopes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        gorm.DeletedAt
}

// ApplicationOauthList represent multiple ApplicationOauth.
type ApplicationOauthList []ApplicationOauth

// TableName return name of table.
func (ao *ApplicationOauth) TableName() string {
	return "application_oauth"
}

// BeforeCreate handle uuid generation.
func (ao *ApplicationOauth) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if ao.UUID == "" {
		ao.UUID = generateUUID.String()
	}
	return nil
}
