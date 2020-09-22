package seeds

import (
	"gorm.io/gorm"
)

// Seed represent it self.
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}
