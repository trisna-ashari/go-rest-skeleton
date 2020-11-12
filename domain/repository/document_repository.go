package repository

import "go-rest-skeleton/domain/entity"

// DocumentRepository is an interface.
type DocumentRepository interface {
	SaveDocument(*entity.Document) (*entity.Document, map[string]string, error)
	UpdateDocument(string, *entity.Document) (*entity.Document, map[string]string, error)
	DeleteDocument(string) error
	GetDocument(string) (*entity.Document, error)
	GetDocuments(parameters *Parameters) ([]entity.Document, interface{}, error)
}
