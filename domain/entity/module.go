package entity

import "github.com/google/uuid"

// Module represent schema of table modules.
type Module struct {
	UUID string `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	Key  string `gorm:"size:100;not null;index:module_key;" json:"module_key"`
}

// Modules represent multiple Module.
type Modules []Module

// BeforeSave handle uuid generation.
func (m *Module) BeforeSave() error {
	generateUUID := uuid.New()
	if m.UUID == "" {
		m.UUID = generateUUID.String()
	}
	return nil
}
