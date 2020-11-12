package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole represent schema of table user_roles.
type UserRole struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	UserUUID  string    `gorm:"size:36;not null;index;" json:"user_uuid"`
	RoleUUID  string    `gorm:"size:100;not null;index;" json:"role_uuid"`
	Role      Role      `gorm:"foreignKey:RoleUUID;association_foreignKey:RoleUUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// UserRoleFaker represent content when generate fake data of user role.
type UserRoleFaker struct {
	UUID     string `faker:"uuid_hyphenated"`
	UserUUID string `faker:"first_name"`
	RoleUUID string `faker:"first_name"`
}

// UserRoles represent multiple user_role.
type UserRoles []UserRole

// TableName return name of table.
func (ur *UserRole) TableName() string {
	return "user_roles"
}

// BeforeCreate handle uuid generation.
func (ur *UserRole) BeforeCreate(tx *gorm.DB) error {
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
