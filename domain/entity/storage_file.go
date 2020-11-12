package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// StorageFile represent schema of table storage_files.
type StorageFile struct {
	UUID         string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	CategoryUUID string    `gorm:"size:36;not null;index;" json:"category_uui"`
	OriginalName string    `gorm:"size:255;not null;" json:"original_name"`
	Name         string    `gorm:"size:255;not null;" json:"name"`
	Path         string    `gorm:"size:255;not null;" json:"path"`
	Type         string    `gorm:"size:36;not null;" json:"type"`
	Size         int64     `gorm:"not null;" json:"size"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}

// TableName return name of table.
func (sf *StorageFile) TableName() string {
	return "storage_files"
}

// BeforeCreate handle uuid generation.
func (sf *StorageFile) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if sf.UUID == "" {
		sf.UUID = generateUUID.String()
	}
	return nil
}
