package mock

import "go-rest-skeleton/infrastructure/authorization"

// AuthInterface is a mock auth.
type AuthInterface struct {
	CreateAuthFn    func(string, *authorization.TokenDetails) error
	FetchAuthFn     func(string) (string, error)
	DeleteRefreshFn func(string) error
	DeleteTokensFn  func(*authorization.AccessDetails) error
}

func (f *AuthInterface) DeleteRefresh(refreshUUID string) error {
	return f.DeleteRefreshFn(refreshUUID)
}
func (f *AuthInterface) DeleteTokens(authD *authorization.AccessDetails) error {
	return f.DeleteTokensFn(authD)
}
func (f *AuthInterface) FetchAuth(uuid string) (string, error) {
	return f.FetchAuthFn(uuid)
}
func (f *AuthInterface) CreateAuth(uuid string, authD *authorization.TokenDetails) error {
	return f.CreateAuthFn(uuid, authD)
}
