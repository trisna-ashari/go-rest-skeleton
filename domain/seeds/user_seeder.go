package seeds

import (
	"errors"
	"fmt"
	"go-rest-skeleton/domain/entity"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// userFactory is a function uses to create []seed.Seed.
func userFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.UserFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		user := &entity.User{
			UUID:     a.UUID,
			Name:     a.Name,
			Email:    a.Email,
			Phone:    a.Phone,
			Password: a.Password,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Name),
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
	var userExists entity.User
	err := db.Where("email = ?", user.Email).Take(&userExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(user).Error
			if err != nil {
				return user, err
			}
			return user, err
		}
		return user, err
	}
	return user, err
}
