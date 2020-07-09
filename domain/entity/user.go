package entity

import (
	"go-rest-skeleton/infrastructure/security"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

// User represent schema of table user.
type User struct {
	UUID      string     `gorm:"size:36;not null;unique_index;" json:"uuid"`
	FirstName string     `gorm:"size:100;not null;" json:"first_name"`
	LastName  string     `gorm:"size:100;not null;" json:"last_name"`
	Email     string     `gorm:"size:100;not null;unique;index:email" json:"email"`
	Phone     string     `gorm:"size:100;" json:"phone,omitempty"`
	Password  string     `gorm:"size:100;not null;index:password" json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// UserFaker represent content when generate fake data of user.
type UserFaker struct {
	UUID      string `faker:"uuid_hyphenated"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Email     string `faker:"email"`
	Phone     string `faker:"phone_number"`
	Password  string `faker:"password"`
}

// Users represent multiple user.
type Users []User

// DetailUser represent format of detail user.
type DetailUser struct {
	UUID      string `gorm:"size:36;not null;unique_index;" json:"uuid"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
	Email     string `gorm:"size:100;not null;unique;index:email" json:"email"`
	Phone     string `gorm:"size:100;" json:"phone,omitempty"`
}

// Prepare will prepare submitted data of user.
func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeSave handle uuid generation and password hashing.
func (u *User) BeforeSave() error {
	generateUUID := uuid.New()
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.UUID = generateUUID.String()
	u.Password = string(hashPassword)
	return nil
}

// DetailUsers will return formatted user detail of multiple user.
func (users Users) DetailUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.DetailUser()
	}
	return result
}

// DetailUser will return formatted user detail of user.
func (u *User) DetailUser() interface{} {
	return &DetailUser{
		UUID:      u.UUID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}

// Validate will validate any action related to user.
func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["email"] = "please provide a valid email"
			}
		}

	case "login":
		if u.Password == "" {
			errorMessages["password"] = "password is required"
		}
		if u.Email == "" {
			errorMessages["email"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["email"] = "please provide a valid email"
			}
		}
	case "forgot_password":
		if u.Email == "" {
			errorMessages["email"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["first_name"] = "first_name is required"
		}
		if u.LastName == "" {
			errorMessages["last_name"] = "last_name is required"
		}
		if u.Password == "" {
			errorMessages["password"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["email"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}
	return errorMessages
}
