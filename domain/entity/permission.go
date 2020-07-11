package entity

// Permission represent schema of table permission.
type Permission struct {
	UUID          string `gorm:"size:36;not null;unique_index;" json:"uuid"`
	ModuleKey     string `gorm:"size:100;not null;" json:"module_key"`
	PermissionKey string `gorm:"size:100;not null;" json:"permission_key"`
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
