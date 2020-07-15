package seeds

import (
	"go-rest-skeleton/infrastructure/seed"
)

// All will seeds all defined seeder.
func All() []seed.Seed {
	b := prepare()
	return b
}

// prepare will prepare fake data based on entity's faker struct.
func prepare() []seed.Seed {
	roleFactories := roleFactory()
	userFactories := userFactory()

	var (
		allFactories []seed.Seed
	)
	allFactories = append(allFactories, roleFactories...)
	allFactories = append(allFactories, userFactories...)

	return allFactories
}
