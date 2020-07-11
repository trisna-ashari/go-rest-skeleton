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

// RoleFactory is a function uses to create []seed.Seed.
func RoleFactory() []seed.Seed {
	roles := []role{
		{UUID: uuid.New().String(), name: "Administrator"},
		{UUID: uuid.New().String(), name: "User"},
	}

	fakerFactories := make([]seed.Seed, 2)
	for i, role := range roles {
		role := role
		fakerFactories[i] = seed.Seed{
			Name: fmt.Sprintf("Create %s", role.name),
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, role.UUID, role.name)
				return err
			},
		}
	}

	return fakerFactories
}
