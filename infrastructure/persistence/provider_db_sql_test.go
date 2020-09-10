package persistence_test

import (
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
)

func TestNewDBConnection_Success(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	dbConn, errConn := persistence.NewDBConnection(conf.DBConfig)
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}

	var db *gorm.DB
	assert.NoError(t, errConn)
	assert.IsType(t, db, dbConn)
}

func TestNewDBConnection_Failed(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = "invalid password"
	_, errConn := persistence.NewDBConnection(conf.DBConfig)
	assert.Error(t, errConn)
}

func TestNewDBService_Success(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}

	var typeRepositories *persistence.Repositories
	assert.NoError(t, errDBService)
	assert.IsType(t, typeRepositories, dbService)
}

func TestNewDBService_Failed(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = "invalid password"
	_, errConn := persistence.NewDBService(conf.DBConfig)
	assert.Error(t, errConn)
}

func TestClose(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}
	errClose := dbService.Close()

	assert.NoError(t, errClose)
}

func TestAutoMigrate(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	conf.DBConfig.DBLog = false
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}
	errClose := dbService.AutoMigrate()

	assert.NoError(t, errClose)
}

func TestSeeds(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	conf.DBConfig.DBLog = false
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}
	errClose := dbService.Seeds()

	assert.NoError(t, errClose)
}

func TestInitialSeeds(t *testing.T) {
	conf := InitConfig()
	conf.DBConfig.DBDriver = conf.DBTestConfig.DBDriver
	conf.DBConfig.DBHost = conf.DBTestConfig.DBHost
	conf.DBConfig.DBPort = conf.DBTestConfig.DBPort
	conf.DBConfig.DBName = conf.DBTestConfig.DBName
	conf.DBConfig.DBUser = conf.DBTestConfig.DBUser
	conf.DBConfig.DBPassword = conf.DBTestConfig.DBPassword
	conf.DBConfig.DBLog = false
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}
	errClose := dbService.InitialSeeds()

	assert.NoError(t, errClose)
}
