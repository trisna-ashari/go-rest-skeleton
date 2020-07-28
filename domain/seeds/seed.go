package seeds

import (
	"github.com/jinzhu/gorm"
)

// Seed represent it self.
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}
