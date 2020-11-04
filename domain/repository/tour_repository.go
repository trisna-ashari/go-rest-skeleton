package repository

import "go-rest-skeleton/domain/entity"

// TourRepository is an interface.
type TourRepository interface {
	SaveTour(tour *entity.Tour) (*entity.Tour, map[string]string, error)
	UpdateTour(UUID string, tour *entity.Tour) (*entity.Tour, map[string]string, error)
	DeleteTour(UUID string) error
	GetTour(UUID string) (*entity.Tour, error)
	GetTourBySlug(tour *entity.Tour) (*entity.Tour, map[string]string, error)
	GetTours(parameters *Parameters) ([]entity.Tour, interface{}, error)
}
