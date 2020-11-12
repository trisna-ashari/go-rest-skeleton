package repository

import "go-rest-skeleton/domain/entity"

// ApplicationOauthClientRepository is an interface.
type ApplicationOauthClientRepository interface {
	SaveApplicationOauthClient(*entity.ApplicationOauthClient) (*entity.ApplicationOauthClient, map[string]string, error)
	UpdateApplicationOauthClient(string, *entity.ApplicationOauthClient) (*entity.ApplicationOauthClient, map[string]string, error)
	DeleteApplicationOauthClient(string) error
	GetApplicationOauthClient(string) (*entity.ApplicationOauthClient, error)
	GetApplicationOauthClients(parameters *Parameters) ([]entity.ApplicationOauthClient, interface{}, error)
}
