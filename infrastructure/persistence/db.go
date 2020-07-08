package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type Repositories struct {
	User repository.UserRepository
	db *gorm.DB
}

func NewRepositories(dbDriver, dbUser, dbPassword, dbHost, dbName, dbPort string) (*Repositories, error) {
	dbUrl := ""
	switch dbDriver {
	case "postgres":
		dbUrl = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	case "mysql":
		dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	db, err := gorm.Open(dbDriver, dbUrl)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		db: db,
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// migrate all tables
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.User{}).Error
}
