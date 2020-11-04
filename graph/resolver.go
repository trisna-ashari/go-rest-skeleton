package graph

import (
	"go-rest-skeleton/infrastructure/persistence"
)

type Resolver struct {
	DBServices *persistence.Repositories
}
