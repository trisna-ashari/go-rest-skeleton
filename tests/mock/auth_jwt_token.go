package mock

import (
	"go-rest-skeleton/infrastructure/authorization"
	"net/http"
)

// TokenInterface is a mock token interface.
type TokenInterface struct {
	CreateTokenFn          func(UUID string) (*authorization.TokenDetails, error)
	ExtractTokenMetadataFn func(*http.Request) (*authorization.AccessDetails, error)
}

func (f *TokenInterface) CreateToken(UUID string) (*authorization.TokenDetails, error) {
	return f.CreateTokenFn(UUID)
}
func (f *TokenInterface) ExtractTokenMetadata(r *http.Request) (*authorization.AccessDetails, error) {
	return f.ExtractTokenMetadataFn(r)
}
