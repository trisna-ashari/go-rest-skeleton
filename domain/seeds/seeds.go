package seeds

// Init will seeds initial seeder.
func Init() []Seed {
	return initFactory()
}

// All will seeds all defined seeder.
func All() []Seed {
	b := prepare()
	return b
}

// prepare will prepare fake data based on entity's faker struct.
func prepare() []Seed {
	roleFactories := roleFactory()
	userFactories := userFactory()

	var (
		allFactories []Seed
	)
	allFactories = append(allFactories, roleFactories...)
	allFactories = append(allFactories, userFactories...)

	return allFactories
}
