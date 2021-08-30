package handlers

import (
	"Assignment-micro-service/gprc-customers/models"
	"Assignment-micro-service/gprc-customers/repositories"
	"Assignment-micro-service/gprc-customers/requests"
	"Assignment-micro-service/pb"
	"context"
	"database/sql"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepository,
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	customer := &models.Customer{}

	err := copier.Copy(&customer, &in)
	if err != nil {
		return nil, err
	}

	newCustomer, err := h.customerRepository.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	cResponse := &pb.Customer{
		Id: int64(newCustomer.Id),
	}
	err = copier.Copy(&cResponse, &newCustomer)
	if err != nil {
		return nil, err
	}
	return cResponse, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	err = copier.Copy(&customer, &in)
	if err != nil {
		return nil, err
	}

	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	cResponse := pb.Customer{
		Id: in.Id,
	}

	err = copier.Copy(&cResponse, &newCustomer)
	if err != nil {
		return nil, err
	}
	return &cResponse, nil
}

func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}

	if in.OldPassword != customer.Password {
		return nil, status.Error(codes.PermissionDenied, "Your old password is incorrect")
	}

	if in.NewPassword != in.ConfirmPassword {
		return nil, status.Error(codes.Canceled, "Makes sure your password match")
	}

	passwordRequest := &requests.ChangePasswordRequest{}
	err = copier.Copy(&passwordRequest, &in)
	if err != nil {
		return nil, err
	}
	err = h.customerRepository.ChangePassword(ctx, passwordRequest)
	if err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{}, nil
}

func (h *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindRequest) (*pb.FindResponse, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}

	cRes := pb.FindResponse{
		Id:      in.Id,
		Name:    customer.Name,
		Address: customer.Address,
		Email:   customer.Email,
	}

	return &cRes, nil
}
