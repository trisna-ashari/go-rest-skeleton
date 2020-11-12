package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserForgotPassword represent schema of table user_login.
type UserForgotPassword struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;"`
	UserUUID  string    `gorm:"size:36;not null;index;"`
	Token     string    `gorm:"size:255;"`
	Password  string    `gorm:"size:255;index;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
	ExpiredAt time.Time
}

// TableName return name of table.
func (ufp *UserForgotPassword) TableName() string {
	return "user_forgot_password"
}

// BeforeCreate handle uuid generation.
func (ufp *UserForgotPassword) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if ufp.UUID == "" {
		ufp.UUID = generateUUID.String()
	}
	return nil
}
