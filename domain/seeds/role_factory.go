package seeds

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type role struct {
	UUID string
	name string
}

// roleFactory is a function uses to create []seed.Seed.
func roleFactory() []Seed {
	roles := []role{
		{UUID: uuid.New().String(), name: "Administrator"},
		{UUID: uuid.New().String(), name: "User"},
	}

	fakerFactories := make([]Seed, 2)
	for i, r := range roles {
		cr := r
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", cr.name),
			Run: func(db *gorm.DB) error {
				err := createRole(db, cr.UUID, cr.name)
				return err
			},
		}
	}

	return fakerFactories
}
