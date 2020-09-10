package entity

import "github.com/google/uuid"

// RolePermission represent schema of table role_permissions.
type RolePermission struct {
	UUID           string     `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	RoleUUID       string     `gorm:"size:100;not null;index:role_uuid;" json:"role_uuid"`
	PermissionUUID string     `gorm:"size:100;not null;index:permission_uuid;" json:"permission_uuid"`
	Permission     Permission `gorm:"foreignKey:UUID;association_foreignKey:PermissionUUID"`
}

// RolePermissionFaker represent content when generate fake data of user role.
type RolePermissionFaker struct {
	UUID     string `faker:"uuid_hyphenated"`
	UserUUID string `faker:"first_name"`
	RoleUUID string `faker:"first_name"`
}

// RolePermissions represent multiple user_role.
type RolePermissions []RolePermission

// BeforeSave handle uuid generation.
func (rp *RolePermission) BeforeSave() error {
	generateUUID := uuid.New()
	if rp.UUID == "" {
		rp.UUID = generateUUID.String()
	}
	return nil
}

// GetRolePermission will return multiple role detail.
func (rp RolePermissions) GetRolePermission() []interface{} {
	result := make([]interface{}, len(rp))
	for index, role := range rp {
		result[index] = role.Permission.DetailPermission()
	}
	return result
}
