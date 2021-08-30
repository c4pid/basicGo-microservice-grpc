package handlers

import (
	"Assignment-micro-service/grpc-bookings/models"
	"Assignment-micro-service/grpc-bookings/repositories"
	"Assignment-micro-service/pb"
	"context"
	"database/sql"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	customerClient    pb.CustomerServiceClient
	flightClient      pb.FlightsClient
	bookingRepository repositories.BookingRepositories
	pb.UnimplementedBookingsServer
}

func NewBookingHandler(customerClient pb.CustomerServiceClient,
	flightClient pb.FlightsClient,
	bookingRepository repositories.BookingRepositories) (*BookingHandler, error) {
	return &BookingHandler{
		customerClient:    customerClient,
		flightClient:      flightClient,
		bookingRepository: bookingRepository,
	}, nil
}

func (h *BookingHandler) Booking(ctx context.Context, in *pb.Info) (*pb.Info, error) {
	booking := &models.Booking{
		BookingDate: in.BookingDate.AsTime(),
	}

	err := copier.Copy(&booking, &in)
	if err != nil {
		return nil, err
	}

	newBooking, err := h.bookingRepository.Booking(ctx, booking)
	if err != nil {
		return nil, err
	}

	bRes := &pb.Info{
		BookingDate: in.BookingDate,
	}

	err = copier.Copy(&bRes, &newBooking)
	if err != nil {
		return nil, err
	}
	return bRes, nil
}

func (h *BookingHandler) ViewBooking(ctx context.Context, in *pb.ViewRequest) (*pb.ViewResponse, error) {
	booking, err := h.bookingRepository.GetBookingByID(ctx, in.BookingCode)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "id not exist")
			}
		} else {
			return nil, err
		}
	}

	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindRequest{
		Id: booking.BookingCode,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer not exist")
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	flight, err := h.flightClient.SearchFlyByID(ctx, &pb.ID{
		Id: booking.BookingCode,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "flight not exist")
			}
		} else {
			return nil, err
		}
	}

	bRes := &pb.ViewResponse{
		BookingCode: in.BookingCode,
		BookingDate: timestamppb.New(booking.BookingDate),
		Customer:    customer,
		Flight:      flight,
	}

	return bRes, nil
}

func (h *BookingHandler) CancleBooking(ctx context.Context, in *pb.ViewRequest) (*pb.Info, error) {
	booking, err := h.bookingRepository.GetBookingByID(ctx, in.BookingCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	booking.Status = "Cancel"

	newBooking, err := h.bookingRepository.CancleBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	bRes := &pb.Info{
		BookingCode: newBooking.BookingCode,
		BookingDate: timestamppb.New(newBooking.BookingDate),
		CustomerId:  newBooking.CustomerId,
		FlightId:    newBooking.FlightId,
		Status:      newBooking.Status,
	}

	return bRes, nil

}
