package repositories

import (
	"Assignment-micro-service/database"
	"Assignment-micro-service/grpc-bookings/models"
	"context"

	"gorm.io/gorm"
)

type BookingRepositories interface {
	Booking(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	GetBookingByID(ctx context.Context, id int64) (*models.Booking, error)
	CancleBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Booking{})

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) Booking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	if err := m.Save(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *dbmanager) GetBookingByID(ctx context.Context, id int64) (*models.Booking, error) {
	booking := &models.Booking{}
	if err := m.Where(&models.Booking{BookingCode: id}).First(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *dbmanager) CancleBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {

	if err := m.Where(&models.Booking{BookingCode: booking.BookingCode}).Save(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}
