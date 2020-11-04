package application

import (
	"go-rest-skeleton/domain/entity"
	"go-rest-skeleton/domain/repository"
)

type tourApp struct {
	tr repository.TourRepository
}

// tourApp implement the TourAppInterface.
var _ TourAppInterface = &tourApp{}

// TourAppInterface is an interface.
type TourAppInterface interface {
	SaveTour(*entity.Tour) (*entity.Tour, map[string]string, error)
	UpdateTour(UUID string, tour *entity.Tour) (*entity.Tour, map[string]string, error)
	DeleteTour(UUID string) error
	GetTours(p *repository.Parameters) ([]entity.Tour, interface{}, error)
	GetTour(UUID string) (*entity.Tour, error)
	GetTourBySlug(tour *entity.Tour) (*entity.Tour, map[string]string, error)
}

func (t tourApp) SaveTour(tour *entity.Tour) (*entity.Tour, map[string]string, error) {
	return t.tr.SaveTour(tour)
}

func (t tourApp) UpdateTour(UUID string, tour *entity.Tour) (*entity.Tour, map[string]string, error) {
	return t.tr.UpdateTour(UUID, tour)
}

func (t tourApp) DeleteTour(UUID string) error {
	return t.tr.DeleteTour(UUID)
}

func (t tourApp) GetTours(p *repository.Parameters) ([]entity.Tour, interface{}, error) {
	return t.tr.GetTours(p)
}

func (t tourApp) GetTour(UUID string) (*entity.Tour, error) {
	return t.tr.GetTour(UUID)
}

func (t tourApp) GetTourBySlug(tour *entity.Tour) (*entity.Tour, map[string]string, error) {
	return t.tr.GetTourBySlug(tour)
}
