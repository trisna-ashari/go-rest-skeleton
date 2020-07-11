package seeds

import (
	"go-rest-skeleton/infrastructure/seed"
)

// All will seeds all defined seeder.
func All() []seed.Seed {
	b := Prepare()
	return b
}

// Prepare will prepare fake data based on entity's faker struct.
func Prepare() []seed.Seed {
	roleFactories := RoleFactory()
	userFactories := UserFactory()

	var (
		allFactories []seed.Seed
	)
	allFactories = append(allFactories, roleFactories...)
	allFactories = append(allFactories, userFactories...)

	return allFactories
}
