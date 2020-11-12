package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Document represent schema of table modules.
type Document struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	UserUUID  string    `gorm:"size:36;not null;index;" json:"user_uuid"`
	Title     string    `gorm:"size:255;not null;index;" json:"title"`
	Workflow  string    `gorm:"size:16;not null;index;" json:"workflow"`
	Type      string    `gorm:"size:16;not null;index;" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// Documents represent multiple Document.
type Documents []Document

// TableName return name of table.
func (d *Document) TableName() string {
	return "documents"
}

// BeforeCreate handle uuid generation.
func (d *Document) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if d.UUID == "" {
		d.UUID = generateUUID.String()
	}
	return nil
}
