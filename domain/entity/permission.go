package entity

import "github.com/google/uuid"

// Permission represent schema of table permissions.
type Permission struct {
	UUID          string `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	ModuleKey     string `gorm:"size:100;not null;index:module_key;" json:"module_key"`
	PermissionKey string `gorm:"size:100;not null;index:permission_key;" json:"permission_key"`
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

// BeforeSave handle uuid generation.
func (p *Permission) BeforeSave() error {
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
