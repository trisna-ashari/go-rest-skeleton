package entity

import (
	"fmt"
	"go-rest-skeleton/pkg/response"
	"go-rest-skeleton/pkg/validator"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Tour represent schema of table tours.
type Tour struct {
	UUID      string     `gorm:"size:36;not null;unique_index;primary_key" json:"uuid"`
	Name      string     `gorm:"size:100;not null;index;" json:"name" form:"name"`
	Slug      string     `gorm:"size:100;not null;uniqueIndex;" json:"slug" form:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int        `gorm:"default:null" json:"created_by,omitempty"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy int        `gorm:"default:null" json:"updated_by,omitempty"`
	ExpiredAt *time.Time `gorm:"default:null" json:"expired_at"`
	DeletedAt gorm.DeletedAt
	DeletedBy int `gorm:"default:null" json:"deleted_by,omitempty"`
}

// Tours represent multiple Tour.
type Tours []Tour

// TableName return name of table.
func (t *Tour) TableName() string {
	return "tours"
}

// FilterableFields return fields.
func (t *Tour) FilterableFields() []interface{} {
	return []interface{}{"name", "slug"}
}

// BeforeCreate handle uuid generation.
func (t *Tour) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if t.UUID == "" {
		t.UUID = generateUUID.String()
	}
	return nil
}

// FieldsForTourDetail represent fields for Tour detail.
type FieldsForTourDetail struct {
	UUID string `gorm:"size:36;not null;unique_index;" json:"uuid"`
	Name string `gorm:"size:100;not null;" json:"name"`
	Slug string `gorm:"size:100;not null;" json:"slug"`
}

// FieldsForTourList represent fields for Tour list.
type FieldsForTourList struct {
	CreatedAt time.Time  `json:"created_at"`
	ExpiredAt *time.Time `json:"expired_at"`
}

// DetailTour represent format of detail Tour.
type DetailTour struct {
	FieldsForTourDetail
}

// DetailTourList represent format of detail Tour list.
type DetailTourList struct {
	FieldsForTourDetail
	FieldsForTourList
}

// DetailTour will return formatted user detail of user.
func (t *Tour) DetailTour() interface{} {
	return &DetailTour{
		FieldsForTourDetail: FieldsForTourDetail{
			UUID: t.UUID,
			Name: t.Name,
			Slug: t.Slug,
		},
	}
}

// DetailTourList will return formatted user detail of user list.
func (t *Tour) DetailTourList() interface{} {
	return &DetailTourList{
		FieldsForTourDetail: FieldsForTourDetail{
			UUID: t.UUID,
			Name: t.Name,
			Slug: t.Slug,
		},
		FieldsForTourList: FieldsForTourList{
			CreatedAt: t.CreatedAt,
			ExpiredAt: t.ExpiredAt,
		},
	}
}

// DetailRoles will return formatted user detail of multiple tour.
func (tours Tours) DetailTours() []interface{} {
	result := make([]interface{}, len(tours))
	for index, tour := range tours {
		result[index] = tour.DetailTourList()
	}
	return result
}

// ValidateSaveTour will validate create a new Tour request.
func (t *Tour) ValidateSaveTour() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", t.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("slug", t.Slug, validation.AddRule().Required().IsLowerAlphaUnderscore().Length(3, 64).Apply()).
		Set("expired_at", t.ExpiredAt, validation.AddRule().When(t.ExpiredAt != nil,
			validation.AddRule(),
		).Apply())

	fmt.Println(validation)

	return validation.Validate()
}

// ValidateUpdateTour will validate update Tour request.
func (t *Tour) ValidateUpdateTour() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", t.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("slug", t.Slug, validation.AddRule().Required().IsLowerAlphaUnderscore().Length(3, 64).Apply()).
		Set("expired_at", t.ExpiredAt, validation.AddRule().When(t.ExpiredAt != nil,
			validation.AddRule().IsTime("2020-01-01 20:01:01"),
		).Apply())

	return validation.Validate()
}
