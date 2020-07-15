package seeds

import (
	"fmt"
	"go-rest-skeleton/infrastructure/seed"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type role struct {
	UUID string
	name string
}

// roleFactory is a function uses to create []seed.Seed.
func roleFactory() []seed.Seed {
	roles := []role{
		{UUID: uuid.New().String(), name: "Administrator"},
		{UUID: uuid.New().String(), name: "User"},
	}

	fakerFactories := make([]seed.Seed, 2)
	for i, r := range roles {
		cr := r
		fakerFactories[i] = seed.Seed{
			Name: fmt.Sprintf("Create %s", cr.name),
			Run: func(db *gorm.DB) error {
				err := createRole(db, cr.UUID, cr.name)
				return err
			},
		}
	}

	return fakerFactories
}
