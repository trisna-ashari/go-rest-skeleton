package mock

import "go-rest-skeleton/infrastructure/authorization"

// AuthInterface is a mock of authorization.AuthInterface.
type AuthInterface struct {
	CreateAuthFn    func(string, *authorization.TokenDetails) error
	FetchAuthFn     func(string) (string, error)
	DeleteRefreshFn func(string) error
	DeleteTokensFn  func(*authorization.AccessDetails) error
}

// DeleteRefresh calls the DeleteRefreshFn.
func (f *AuthInterface) DeleteRefresh(refreshUUID string) error {
	return f.DeleteRefreshFn(refreshUUID)
}

// DeleteTokens calls the DeleteTokensFn.
func (f *AuthInterface) DeleteTokens(authD *authorization.AccessDetails) error {
	return f.DeleteTokensFn(authD)
}

// FetchAuth calls the FetchAuthFn.
func (f *AuthInterface) FetchAuth(uuid string) (string, error) {
	return f.FetchAuthFn(uuid)
}

// CreateAuth calls the CreateAuthFn.
func (f *AuthInterface) CreateAuth(uuid string, authD *authorization.TokenDetails) error {
	return f.CreateAuthFn(uuid, authD)
}
