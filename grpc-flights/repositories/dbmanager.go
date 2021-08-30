package repositories

import (
	"Assignment-micro-service/database"
	"Assignment-micro-service/grpc-flights/models"
	"context"

	"gorm.io/gorm"
)

type FlightRepository interface {
	GetFlightByID(ctx context.Context, id int64) (*models.Flight, error)
	CreateFlight(ctx context.Context, flight *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, flight *models.Flight) (*models.Flight, error)
	GetFlightsByName(ctx context.Context, name string) ([]*models.Flight, error)
	GetFlightsByFrom(ctx context.Context, from string) ([]*models.Flight, error)
	GetFlightsByDestination(ctx context.Context, destination string) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Flight{})

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, flight *models.Flight) (*models.Flight, error) {
	if err := m.Save(flight).Error; err != nil {
		return nil, err
	}
	return flight, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, flight *models.Flight) (*models.Flight, error) {
	if err := m.Where(&models.Flight{Id: flight.Id}).Save(flight).Error; err != nil {
		return nil, err
	}

	return flight, nil
}

func (m *dbmanager) GetFlightByID(ctx context.Context, id int64) (*models.Flight, error) {
	flight := models.Flight{}

	if err := m.Where(&models.Flight{Id: uint(id)}).First(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (m *dbmanager) GetFlightsByName(ctx context.Context, name string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if err := m.Where(&models.Flight{Name: name}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (m *dbmanager) GetFlightsByFrom(ctx context.Context, from string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if err := m.Where(&models.Flight{From: from}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (m *dbmanager) GetFlightsByDestination(ctx context.Context, destination string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if err := m.Where(&models.Flight{To: destination}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}
