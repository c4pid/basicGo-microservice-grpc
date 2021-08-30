package repositories

import (
	"Assignment-micro-service/database"
	"Assignment-micro-service/gprc-customers/models"
	"Assignment-micro-service/gprc-customers/requests"
	"context"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerByID(ctx context.Context, id int64) (*models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	ChangePassword(ctx context.Context, request *requests.ChangePasswordRequest) error
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Customer{})

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if err := m.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (m *dbmanager) GetCustomerByID(ctx context.Context, id int64) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Id: uint(id)}).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (m *dbmanager) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if err := m.Where(&models.Customer{Id: customer.Id}).Updates(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (m *dbmanager) ChangePassword(ctx context.Context, request *requests.ChangePasswordRequest) error {
	if err := m.Where(&models.Customer{Id: uint(request.Id)}).Updates(models.Customer{Password: request.ConfirmPassword}).Error; err != nil {
		return err
	}
	return nil
}
