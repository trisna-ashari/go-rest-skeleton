package persistence_test

import (
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteRole_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	roles, seedErr := seedRoles(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}

	role := roles[0]
	repo := persistence.NewRoleRepository(conn)
	getErr := repo.DeleteRole(role.UUID)

	assert.Nil(t, getErr)
}

func TestGetRole_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	roles, seedErr := seedRoles(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	role := roles[0]
	repo := persistence.NewRoleRepository(conn)
	r, getErr := repo.GetRole(role.UUID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, r.UUID, role.UUID)
	assert.EqualValues(t, r.Name, role.Name)
}

func TestGetRoles_Success(t *testing.T) {
	SkipThis(t)

	conn, connErr := DBConn()
	if connErr != nil {
		t.Fatalf("want non error, got %#v", connErr)
	}
	_, seedErr := seedRoles(conn)
	if seedErr != nil {
		t.Fatalf("want non error, got %#v", seedErr)
	}
	repo := persistence.NewRoleRepository(conn)
	params := repository.Parameters{
		Offset:  0,
		Limit:   3,
		PerPage: 3,
		Page:    1,
		Order:   "desc",
	}
	r, _, getErr := repo.GetRoles(&params)

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(r), 3)
}
