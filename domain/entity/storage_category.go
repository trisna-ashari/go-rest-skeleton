package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// StorageCategory represent schema of table storage_categories.
type StorageCategory struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	Slug      string    `gorm:"size:100;not null;uniqueIndex:slug;" json:"slug" form:"slug"`
	Path      string    `gorm:"size:100;not null;" json:"path" form:"path"`
	Name      string    `gorm:"size:100;not null;" json:"name" form:"name"`
	MimeTypes string    `gorm:"size:255;not null;" json:"mime_types" form:"mime_types"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// TableName return name of table.
func (sc *StorageCategory) TableName() string {
	return "storage_categories"
}

// BeforeCreate handle uuid generation.
func (sc *StorageCategory) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if sc.UUID == "" {
		sc.UUID = generateUUID.String()
	}
	return nil
}
