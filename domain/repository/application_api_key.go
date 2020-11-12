package repository

import "go-rest-skeleton/domain/entity"

// ApplicationApiKeyRepository is an interface.
type ApplicationApiKeyRepository interface {
	SaveApplicationApiKey(*entity.ApplicationApiKey) (*entity.ApplicationApiKey, map[string]string, error)
	UpdateApplicationApiKey(string, *entity.ApplicationApiKey) (*entity.ApplicationApiKey, map[string]string, error)
	ActivateApplicationApiKey(string) error
	DeactivateApplicationApiKey(string) error
	DeleteApplicationApiKey(string) error
	GetApplicationApiKey(string) (*entity.ApplicationApiKey, error)
	GetApplicationApiKeys(parameters *Parameters) ([]entity.ApplicationApiKey, interface{}, error)
}
