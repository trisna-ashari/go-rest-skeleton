package entity

import (
	"time"

	"github.com/google/uuid"
)

type StorageFile struct {
	UUID         string     `gorm:"size:36;not null;unique_index;primary_key;" json:"uuid"`
	CategoryUUID string     `gorm:"size:36;not null;" json:"category_uui"`
	OriginalName string     `gorm:"size:255;not null;" json:"original_name"`
	Name         string     `gorm:"size:255;not null;" json:"name"`
	Path         string     `gorm:"size:255;not null;" json:"path"`
	Type         string     `gorm:"size:36;not null;" json:"type"`
	Size         int64      `gorm:"size:11;not null;" json:"size"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type File struct {
	URL string
}

// BeforeSave handle uuid generation.
func (sf *StorageFile) BeforeSave() error {
	generateUUID := uuid.New()
	if sf.UUID == "" {
		sf.UUID = generateUUID.String()
	}
	return nil
}
