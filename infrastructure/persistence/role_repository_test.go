package persistence_test

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateRole_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}

	role, errSeed := seedRole(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewRoleRepository(conn)
	roleData := entity.Role{Name: "Updated " + role.Name}

	r, _, errUpdate := repo.UpdateRole(role.UUID, &roleData)
	assert.NoError(t, errUpdate)
	assert.EqualValues(t, r.Name, "Updated "+role.Name)
}

func TestSaveRole_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}

	var role = entity.Role{}
	role.Name = "Test Create Role"

	repo := persistence.NewRoleRepository(conn)

	r, _, errSave := repo.SaveRole(&role)
	assert.NoError(t, errSave)
	assert.NotNil(t, r.UUID)
	assert.EqualValues(t, r.Name, role.Name)
}

func TestDeleteRole_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	roles, errSeed := seedRoles(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}

	role := roles[0]
	repo := persistence.NewRoleRepository(conn)
	errGet := repo.DeleteRole(role.UUID)

	assert.Nil(t, errGet)
}

func TestGetRole_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	roles, errSeed := seedRoles(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	role := roles[0]
	repo := persistence.NewRoleRepository(conn)
	r, errGet := repo.GetRole(role.UUID)

	assert.Nil(t, errGet)
	assert.EqualValues(t, r.UUID, role.UUID)
	assert.EqualValues(t, r.Name, role.Name)
}

func TestGetRoles_Success(t *testing.T) {
	SkipThis(t)

	conn, errConn := DBConn()
	if errConn != nil {
		t.Fatalf("want non error, got %#v", errConn)
	}
	_, errSeed := seedRoles(conn)
	if errSeed != nil {
		t.Fatalf("want non error, got %#v", errSeed)
	}
	repo := persistence.NewRoleRepository(conn)
	params := repository.Parameters{
		Offset:  0,
		Limit:   3,
		PerPage: 3,
		Page:    1,
		Order:   "desc",
	}
	r, _, errGet := repo.GetRoles(&params)

	assert.Nil(t, errGet)
	assert.EqualValues(t, len(r), 3)
}
