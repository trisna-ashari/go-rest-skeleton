package entity

import (
	"time"

	"github.com/google/uuid"
)

// StorageCategory represent schema of table storage_categories.
type StorageCategory struct {
	UUID      string     `gorm:"size:36;not null;unique_index;primary_key;" json:"uuid"`
	Slug      string     `gorm:"size:100;not null;index:slug;" json:"slug"`
	Path      string     `gorm:"size:100;not null;" json:"path"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	MimeTypes string     `gorm:"size:255;not null;" json:"mime_types"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BeforeSave handle uuid generation.
func (sc *StorageCategory) BeforeSave() error {
	generateUUID := uuid.New()
	if sc.UUID == "" {
		sc.UUID = generateUUID.String()
	}
	return nil
}
