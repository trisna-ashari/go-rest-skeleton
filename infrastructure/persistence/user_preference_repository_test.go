package persistence_test

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserPreference_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	userPreference, errSeed := seedUserPreference(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserPreferenceRepository(conn)
	r, errGet := repo.GetUserPreference(userPreference.UserUUID)

	assert.Nil(t, errGet)
	assert.EqualValues(t, r.UUID, userPreference.UUID)
	assert.EqualValues(t, r.UserUUID, userPreference.UserUUID)
	assert.EqualValues(t, r.Preference, userPreference.Preference)
}

func TestUpdateUserPreference_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}

	userPreference, errSeed := seedUserPreference(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserPreferenceRepository(conn)
	userPreferenceData := entity.DetailUserPreference{
		Language: "id",
		DarkMode: false,
	}

	r, _, errUpdate := repo.UpdateUserPreference(userPreference.UserUUID, &userPreferenceData)
	assert.NoError(t, errUpdate)
	assert.EqualValues(t, r.Preference, userPreference.BuildPatchPreference(&userPreferenceData))
}

func TestResetUserPreference_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}

	userPreference, errSeed := seedUserPreference(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewUserPreferenceRepository(conn)
	userPreferenceData := entity.DetailUserPreference{
		Language: "id",
		DarkMode: false,
	}

	u, _, errUpdate := repo.UpdateUserPreference(userPreference.UserUUID, &userPreferenceData)
	r, errReset := repo.ResetUserPreference(userPreference.UserUUID)
	assert.NoError(t, errUpdate)
	assert.NoError(t, errReset)
	assert.NotEqualValues(t, u.Preference, r.Preference)
}
