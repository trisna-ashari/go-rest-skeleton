package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Module represent schema of table modules.
type Module struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	Key       string    `gorm:"size:100;not null;index:module_key;" json:"module_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// Modules represent multiple Module.
type Modules []Module

// TableName return name of table.
func (m *Module) TableName() string {
	return "modules"
}

// BeforeCreate handle uuid generation.
func (m *Module) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if m.UUID == "" {
		m.UUID = generateUUID.String()
	}
	return nil
}
