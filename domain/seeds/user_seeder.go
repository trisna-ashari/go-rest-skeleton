package seeds

import (
	"fmt"
	"go-rest-skeleton/domain/entity"

	"github.com/bxcodec/faker"

	"github.com/jinzhu/gorm"
)

// userFactory is a function uses to create []seed.Seed.
func userFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.UserFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			fmt.Println(errFaker)
		}

		user := &entity.User{
			UUID:      a.UUID,
			FirstName: a.FirstName,
			LastName:  a.LastName,
			Email:     a.Email,
			Phone:     a.Phone,
			Password:  a.Password,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.FirstName),
			Run: func(db *gorm.DB) error {
				_, errDB := createUser(db, user)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createUser will create fake user and insert into DB.
func createUser(db *gorm.DB, user *entity.User) (*entity.User, error) {
	err := db.Create(user).Error
	return user, err
}
