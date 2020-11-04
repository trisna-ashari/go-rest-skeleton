package persistence

import (
	"errors"
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
	"go-rest-skeleton/infrastructure/message/exception"

	"gorm.io/gorm"
)

// TourRepo is a struct to store db connection.
type TourRepo struct {
	db *gorm.DB
}

// NewTourRepository will initialize Tour repository.
func NewTourRepository(db *gorm.DB) *TourRepo {
	return &TourRepo{db}
}

// TourRepo implements the repository.TourRepository interface.
var _ repository.TourRepository = &TourRepo{}

// SaveTour will create a new Tour.
func (r TourRepo) SaveTour(Tour *entity.Tour) (*entity.Tour, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Tour).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Tour, nil, nil
}

func (r TourRepo) UpdateTour(uuid string, tour *entity.Tour) (*entity.Tour, map[string]string, error) {
	panic("implement me")
}

func (r TourRepo) DeleteTour(uuid string) error {
	panic("implement me")
}

func (r TourRepo) GetTour(uuid string) (*entity.Tour, error) {
	var tour entity.Tour
	err := r.db.Where("uuid = ?", uuid).Take(&tour).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrorTextRoleNotFound
	}
	return &tour, nil
}

func (r TourRepo) GetTourBySlug(t *entity.Tour) (*entity.Tour, map[string]string, error) {
	var tour entity.Tour
	errDesc := map[string]string{}
	err := r.db.Where("slug = ?", t.Slug).Take(&tour).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["slug"] = exception.ErrorTextTourSlugAlreadyExists.Error()
			return nil, errDesc, exception.ErrorTextTourSlugAlreadyExists
		}
		return nil, errDesc, err
	}

	return &tour, nil, nil
}

func (r TourRepo) GetTours(p *repository.Parameters) ([]entity.Tour, interface{}, error) {
	var total int64
	var tours []entity.Tour
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&tours).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&tours).Error
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
	return tours, meta, nil
}
