package seeds

import (
	"fmt"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/seed"

	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
)

// userFactory is a function uses to create []seed.Seed.
func userFactory() []seed.Seed {
	fakerFactories := make([]seed.Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.UserFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			fmt.Println(errFaker)
		}
		fakerFactories[i] = seed.Seed{
			Name: fmt.Sprintf("Create %s", a.FirstName),
			Run: func(db *gorm.DB) error {
				errDB := createUser(db, a.UUID, a.FirstName, a.LastName, a.Email, a.Phone, a.Password)
				return errDB
			},
		}
	}

	return fakerFactories
}
