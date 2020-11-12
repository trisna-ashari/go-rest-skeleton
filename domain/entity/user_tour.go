package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserTour represent schema of table user_roles.
type UserTour struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"uuid"`
	UserUUID  string    `gorm:"size:36;not null;index;" json:"user_uuid"`
	TourUUID  string    `gorm:"size:100;not null;index;" json:"role_uuid"`
	Tour      Tour      `gorm:"foreignKey:TourUUID;association_foreignKey:TourUUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// UserTours represent multiple user_role.
type UserTours []UserTour

// TableName return name of table.
func (ut *UserTour) TableName() string {
	return "user_tours"
}

// BeforeCreate handle uuid generation.
func (ut *UserTour) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if ut.UUID == "" {
		ut.UUID = generateUUID.String()
	}
	return nil
}

// GetUserTour will return multiple role detail.
func (ut UserTours) GetUserTour() []interface{} {
	result := make([]interface{}, len(ut))
	for index, role := range ut {
		result[index] = role.Tour.DetailTour()
	}
	return result
}
