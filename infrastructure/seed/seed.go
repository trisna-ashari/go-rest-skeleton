package seed

import (
	"github.com/jinzhu/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}