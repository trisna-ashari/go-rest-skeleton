package entity

import "github.com/google/uuid"

// UserRole represent schema of table user_roles.
type UserRole struct {
	UUID     string `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	UserUUID string `gorm:"size:36;not null;index:user_uuid;" json:"user_uuid"`
	RoleUUID string `gorm:"size:100;not null;index:role_uuid;" json:"role_uuid"`
	Role     Role   `gorm:"foreignKey:UUID;association_foreignKey:RoleUUID"`
}

// UserRoleFaker represent content when generate fake data of user role.
type UserRoleFaker struct {
	UUID     string `faker:"uuid_hyphenated"`
	UserUUID string `faker:"first_name"`
	RoleUUID string `faker:"first_name"`
}

// UserRoles represent multiple user_role.
type UserRoles []UserRole

// BeforeSave handle uuid generation.
func (ur *UserRole) BeforeSave() error {
	generateUUID := uuid.New()
	if ur.UUID == "" {
		ur.UUID = generateUUID.String()
	}
	return nil
}

// GetUserRole will return multiple role detail.
func (ur UserRoles) GetUserRole() []interface{} {
	result := make([]interface{}, len(ur))
	for index, role := range ur {
		result[index] = role.Role.DetailRole()
	}
	return result
}
