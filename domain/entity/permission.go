package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Permission represent schema of table permissions.
type Permission struct {
	UUID          string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	ModuleKey     string    `gorm:"size:100;not null;index:module_key;" json:"module_key" form:"module_key"`
	PermissionKey string    `gorm:"size:100;not null;uniqueIndex:permission_key;" json:"permission_key" form:"permission_key"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     gorm.DeletedAt
}

// Permissions represent multiple permission.
type Permissions []Permission

// FieldsForPermissionDetail represent fields for permission detail.
type FieldsForPermissionDetail struct {
	UUID          string `gorm:"size:36;not null;unique_index;" json:"uuid"`
	ModuleKey     string `gorm:"size:100;not null;" json:"module_key"`
	PermissionKey string `gorm:"size:100;not null;" json:"permission_key"`
}

// DetailPermission represent format of permission detail.
type DetailPermission struct {
	FieldsForPermissionDetail
}

// DetailPermissionList represent format of permission list.
type DetailPermissionList struct {
	FieldsForPermissionDetail
}

// TableName return name of table.
func (p *Permission) TableName() string {
	return "permissions"
}

// BeforeCreate handle uuid generation.
func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if p.UUID == "" {
		p.UUID = generateUUID.String()
	}
	return nil
}

// DetailPermission will return formatted permission detail.
func (p *Permission) DetailPermission() interface{} {
	return &DetailPermission{
		FieldsForPermissionDetail: FieldsForPermissionDetail{
			UUID:          p.UUID,
			ModuleKey:     p.ModuleKey,
			PermissionKey: p.PermissionKey,
		},
	}
}
