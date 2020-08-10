package persistence_test

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/bxcodec/faker"

	"github.com/stretchr/testify/assert"
)

func TestSaveUser_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	var user = entity.User{}
	var userFaker = entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user.Email = userFaker.Email
	user.FirstName = userFaker.FirstName
	user.LastName = userFaker.LastName
	user.Password = userFaker.Password
	user.Phone = userFaker.Phone

	repo := persistence.NewUserRepository(conn)

	u, saveErr, _ := repo.SaveUser(&user)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.FirstName, userFaker.FirstName)
	assert.EqualValues(t, u.LastName, userFaker.LastName)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
	assert.NotEqual(t, u.Password, userFaker.Password)
}

func TestUpdateUser_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, userFaker, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewUserRepository(conn)
	userData := entity.User{
		FirstName: "Updated " + userFaker.FirstName,
		LastName:  "Updated " + userFaker.LastName,
		Email:     "Updated " + userFaker.Email,
	}
	u, updateErr, _ := repo.UpdateUser(user.UUID, &userData)

	assert.Nil(t, updateErr)
	assert.EqualValues(t, u.FirstName, "Updated "+userFaker.FirstName)
	assert.EqualValues(t, u.LastName, "Updated "+userFaker.LastName)
	assert.EqualValues(t, u.Email, "Updated "+userFaker.Email)
}

func TestDeleteUser_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, _, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewUserRepository(conn)
	getErr := repo.DeleteUser(user.UUID)

	assert.Nil(t, getErr)
}

func TestGetUser_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, userFaker, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewUserRepository(conn)
	u, getErr := repo.GetUser(user.UUID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.FirstName, userFaker.FirstName)
	assert.EqualValues(t, u.LastName, userFaker.LastName)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
}

func TestGetUsers_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	_, _, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewUserRepository(conn)
	params := repository.Parameters{
		Offset:  0,
		Limit:   1,
		PerPage: 1,
		Page:    1,
		Order:   "desc",
	}
	users, _, getErr := repo.GetUsers(&params)

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(users), 1)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	user, userFaker, seedErr := seedUser(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewUserRepository(conn)
	userEmailAndPassword := entity.User{Email: user.Email, Password: userFaker.Password}
	u, _, getErr := repo.GetUserByEmailAndPassword(&userEmailAndPassword)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.FirstName, userFaker.FirstName)
	assert.EqualValues(t, u.LastName, userFaker.LastName)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
}
