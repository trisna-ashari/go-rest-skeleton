package entity

import "time"

// Role represent schema of table roles.
type Role struct {
	UUID           string           `gorm:"size:36;not null;unique_index;" json:"uuid"`
	Name           string           `gorm:"size:100;not null;" json:"name"`
	CreatedAt      time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy      int              `gorm:"default:null" json:"created_by,omitempty"`
	UpdatedAt      time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy      int              `gorm:"default:null" json:"updated_by,omitempty"`
	DeletedAt      *time.Time       `json:"deleted_at,omitempty"`
	DeletedBy      int              `gorm:"default:null" json:"deleted_by,omitempty"`
	RolePermission []RolePermission `gorm:"foreignKey:RoleUUID"`
}

// RoleFaker represent content when generate fake data of role.
type RoleFaker struct {
}

// Roles represent multiple role.
type Roles []Role

// FieldsForRoleDetail represent fields for role detail.
type FieldsForRoleDetail struct {
	UUID string `gorm:"size:36;not null;unique_index;" json:"uuid"`
	Name string `gorm:"size:100;not null;" json:"name"`
}

// FieldsForRoleList represent fields for role list.
type FieldsForRoleList struct {
	CreatedAt time.Time `json:"created_at"`
}

// DetailRole represent format of detail role.
type DetailRole struct {
	FieldsForRoleDetail
}

// DetailRoleList represent format of detail role list.
type DetailRoleList struct {
	FieldsForRoleDetail
	FieldsForRoleList
}

// DetailRole will return formatted user detail of user.
func (r *Role) DetailRole() interface{} {
	return &DetailRole{
		FieldsForRoleDetail: FieldsForRoleDetail{
			UUID: r.UUID,
			Name: r.Name,
		},
	}
}

// DetailRoleList will return formatted user detail of user list.
func (r *Role) DetailRoleList() interface{} {
	return &DetailRoleList{
		FieldsForRoleDetail: FieldsForRoleDetail{
			UUID: r.UUID,
			Name: r.Name,
		},
		FieldsForRoleList: FieldsForRoleList{
			CreatedAt: r.CreatedAt,
		},
	}
}

// DetailRoles will return formatted user detail of multiple role.
func (roles Roles) DetailRoles() []interface{} {
	result := make([]interface{}, len(roles))
	for index, role := range roles {
		result[index] = role.DetailRoleList()
	}
	return result
}
