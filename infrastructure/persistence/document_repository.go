package persistence

import (
	"errors"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"

	"gorm.io/gorm"
)

// DocumentRepo is a struct to store db connection.
type DocumentRepo struct {
	db *gorm.DB
}

// NewDocumentRepository will initialize Document repository.
func NewDocumentRepository(db *gorm.DB) *DocumentRepo {
	return &DocumentRepo{db}
}

// DocumentRepo implements the repository.DocumentRepository interface.
var _ repository.DocumentRepository = &DocumentRepo{}

// SaveDocument will create a new document.
func (r DocumentRepo) SaveDocument(document *entity.Document) (*entity.Document, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&document).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}

	return document, nil, nil
}

// UpdateDocument will update specified document.
func (r DocumentRepo) UpdateDocument(uuid string, document *entity.Document) (*entity.Document, map[string]string, error) {
	errDesc := map[string]string{}
	documentData := &entity.Document{
		Title: document.Title,
	}
	err := r.db.First(&document, "uuid = ?", uuid).Updates(documentData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextDocumentInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextDocumentNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return document, nil, nil
}

// DeleteDocument will delete document.
func (r DocumentRepo) DeleteDocument(uuid string) error {
	var document entity.Document
	err := r.db.Where("uuid = ?", uuid).Take(&document).Delete(&document).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextUserNotFound
		}
		return err
	}
	return nil
}

// GetDocument will return a document.
func (r DocumentRepo) GetDocument(uuid string) (*entity.Document, error) {
	var document entity.Document
	err := r.db.Where("uuid = ?", uuid).Take(&document).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrorTextDocumentNotFound
	}

	return &document, nil
}

// GetDocuments will return document list.
func (r DocumentRepo) GetDocuments(p *repository.Parameters) ([]entity.Document, interface{}, error) {
	var total int64
	var documents []entity.Document
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&documents).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&documents).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}

	if errList != nil {
		return nil, nil, errList
	}

	if errors.Is(errList, gorm.ErrRecordNotFound) {
		return nil, nil, errList
	}

	meta := repository.NewMeta(p, total)
	return documents, meta, nil
}
