package entity

import (
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/security"
	"go-rest-skeleton/pkg/validator"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// User represent schema of table users.
type User struct {
	UUID               string    `gorm:"size:36;not null;unique_index;primary_key;" json:"uuid"`
	Name               string    `gorm:"size:100;not null;" json:"name" form:"name"`
	Email              string    `gorm:"size:100;not null;uniqueIndex;" json:"email" form:"email"`
	Phone              string    `gorm:"size:100;" json:"phone,omitempty" form:"phone"`
	Password           string    `gorm:"size:100;not null;index;" json:"password" form:"password"`
	AvatarUUID         string    `gorm:"size:36;" json:"avatar_uuid"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	DeletedAt          gorm.DeletedAt
	UserRoles          []UserRole           `gorm:"foreignKey:UserUUID"`
	UserLogins         []UserLogin          `gorm:"foreignKey:UserUUID"`
	UserForgotPassword []UserForgotPassword `gorm:"foreignKey:UserUUID"`
}

// UserResetPassword represent payload for reset password request.
type UserResetPassword struct {
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

// UserFaker represent content when generate fake data of user.
type UserFaker struct {
	UUID     string `faker:"uuid_hyphenated"`
	Name     string `faker:"name"`
	Email    string `faker:"email"`
	Phone    string `faker:"phone_number"`
	Password string `faker:"password"`
}

// Users represent multiple User.
type Users []User

// DetailUser represent format of detail User.
type DetailUser struct {
	UserFieldsForDetail
	Role []interface{} `json:"roles,omitempty"`
}

// DetailUserList represent format of DetailUser for User list.
type DetailUserList struct {
	UserFieldsForDetail
	UserFieldsForList
}

// UserFieldsForDetail represent fields of detail User.
type UserFieldsForDetail struct {
	UUID   string      `json:"uuid"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Phone  interface{} `json:"phone,omitempty"`
	Avatar interface{} `json:"avatar,omitempty"`
}

// UserFieldsForList represent fields of detail User for User list.
type UserFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *User) TableName() string {
	return "users"
}

// Prepare will prepare submitted data of user.
func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	u.Password = string(hashPassword)
	return nil
}

// DetailUsers will return formatted user detail of multiple user.
func (users Users) DetailUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.DetailUserList()
	}
	return result
}

// DetailUser will return formatted user detail of user.
func (u *User) DetailUser() interface{} {
	return &DetailUser{
		UserFieldsForDetail: UserFieldsForDetail{
			UUID:   u.UUID,
			Name:   u.Name,
			Email:  u.Email,
			Phone:  u.Phone,
			Avatar: nil,
		},
		Role: UserRoles.GetUserRole(u.UserRoles),
	}
}

// DetailUserAvatar will return formatted user detail of user.
func (u *User) DetailUserAvatar(url interface{}) interface{} {
	return &DetailUser{
		UserFieldsForDetail: UserFieldsForDetail{
			UUID:   u.UUID,
			Name:   u.Name,
			Email:  u.Email,
			Phone:  u.Phone,
			Avatar: url,
		},
		Role: UserRoles.GetUserRole(u.UserRoles),
	}
}

// DetailUserList will return formatted user detail of user for user list.
func (u *User) DetailUserList() interface{} {
	return &DetailUserList{
		UserFieldsForDetail: UserFieldsForDetail{
			UUID:  u.UUID,
			Name:  u.Name,
			Email: u.Email,
			Phone: u.Phone,
		},
		UserFieldsForList: UserFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveUser will validate create a new user request.
func (u *User) ValidateSaveUser() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("email", u.Email, validation.AddRule().Required().Length(6, 128).IsEmail().Apply()).
		Set("phone", u.Phone, validation.AddRule().Required().Length(6, 16).IsPhone().Apply()).
		Set("password", u.Password, validation.AddRule().Required().Length(6, 32).Apply())

	return validation.Validate()
}

// ValidateUpdateUser will validate create a new user request.
func (u *User) ValidateUpdateUser() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("email", u.Email, validation.AddRule().Required().Length(6, 128).IsEmail().Apply()).
		Set("phone", u.Phone, validation.AddRule().Required().Length(6, 16).IsPhone().Apply())

	return validation.Validate()
}

// ValidateLogin will validate login request.
func (u *User) ValidateLogin() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("email", u.Email, validation.AddRule().Required().IsEmail().Apply()).
		Set("password", u.Password, validation.AddRule().Required().Apply())

	return validation.Validate()
}

// ValidateForgotPassword will validate forgot password request.
func (u *User) ValidateForgotPassword() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("email", u.Email, validation.AddRule().Required().IsEmail().Apply())

	return validation.Validate()
}

// ValidateResetPassword will validate reset password request.
func (u *UserResetPassword) ValidateResetPassword() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("new_password", u.NewPassword, validation.AddRule().Required().Length(6, 32).Apply()).
		Set("confirm_password", u.ConfirmPassword, validation.AddRule().Required().EqualTo("new_password", u.NewPassword).Length(6, 32).Apply())

	return validation.Validate()
}
