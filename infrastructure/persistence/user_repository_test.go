package persistence_test

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/google/uuid"

	"github.com/bxcodec/faker"

	"github.com/stretchr/testify/assert"
)

func TestSaveUser_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	var user = entity.User{}
	var userFaker = entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user.Email = userFaker.Email
	user.Name = userFaker.Name
	user.Password = userFaker.Password
	user.Phone = userFaker.Phone

	repo := persistence.NewUserRepository(conn)

	u, _, errSave := repo.SaveUser(&user)
	assert.NoError(t, errSave)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.Name, userFaker.Name)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
	assert.NotEqual(t, u.Password, userFaker.Password)
}

func TestUpdateUser_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, userFaker, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}

	repo := persistence.NewUserRepository(conn)
	userData := entity.User{
		Name:  "Updated " + userFaker.Name,
		Email: "Updated " + userFaker.Email,
	}
	u, _, errUpdate := repo.UpdateUser(user.UUID, &userData)

	assert.NoError(t, errUpdate)
	assert.EqualValues(t, u.Name, "Updated "+userFaker.Name)
	assert.EqualValues(t, u.Email, "Updated "+userFaker.Email)
}

func TestDeleteUser_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, _, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserRepository(conn)
	errGet := repo.DeleteUser(user.UUID)

	assert.NoError(t, errGet)
}

func TestGetUser_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, userFaker, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserRepository(conn)
	u, errGet := repo.GetUser(user.UUID)

	assert.NoError(t, errGet)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.Name, userFaker.Name)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
}

func TestGetUsers_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	_, _, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserRepository(conn)
	params := repository.Parameters{
		Offset:  0,
		Limit:   1,
		PerPage: 1,
		Page:    1,
		Order:   "desc",
	}
	users, _, errGet := repo.GetUsers(&params)

	assert.NoError(t, errGet)
	assert.EqualValues(t, len(users), 1)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, userFaker, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserRepository(conn)
	userEmailAndPassword := entity.User{Email: user.Email, Password: userFaker.Password}
	u, _, errGet := repo.GetUserByEmailAndPassword(&userEmailAndPassword)

	assert.NoError(t, errGet)
	assert.EqualValues(t, u.Email, userFaker.Email)
	assert.EqualValues(t, u.Name, userFaker.Name)
	assert.EqualValues(t, u.Phone, userFaker.Phone)
}

func TestUpdateUserAvatar_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	user, _, errSeed := seedUser(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}

	repo := persistence.NewUserRepository(conn)
	avatarUUID := uuid.New().String()
	userData := entity.User{
		AvatarUUID: avatarUUID,
	}
	u, _, errUpdate := repo.UpdateUserAvatar(user.UUID, &userData)

	assert.NoError(t, errUpdate)
	assert.EqualValues(t, u.AvatarUUID, avatarUUID)
}
