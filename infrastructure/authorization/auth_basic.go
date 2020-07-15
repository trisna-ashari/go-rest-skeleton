package authorization

import (
	"go-rest-skeleton/application"
)

// BasicAuth is represent needed dependencies.
type BasicAuth struct {
	us application.UserAppInterface
}

// NewBasicAuth is a constructor.
func NewBasicAuth(us application.UserAppInterface) *BasicAuth {
	return &BasicAuth{us: us}
}
