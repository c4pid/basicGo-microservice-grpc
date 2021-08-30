package handlers

import (
	"Assignment-micro-service/grpc-flights/models"
	"Assignment-micro-service/grpc-flights/repositories"
	"Assignment-micro-service/pb"
	"context"
	"database/sql"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	pb.UnimplementedFlightsServer
	flightRepository repositories.FlightRepository
}

func NewFlightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
	}, nil
}
func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {

	flight := models.Flight{
		Date: in.Date.AsTime(),
	}

	err := copier.Copy(&flight, &in)
	if err != nil {
		return nil, err
	}

	newFlight, err := h.flightRepository.CreateFlight(ctx, &flight)
	if err != nil {
		return nil, err
	}

	fRes := &pb.Flight{
		Date: in.Date,
	}
	err = copier.Copy(&fRes, &newFlight)
	if err != nil {
		return nil, err
	}
	return fRes, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.flightRepository.GetFlightByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}
	flight.Date = in.Date.AsTime()
	err = copier.Copy(&flight, &in)
	if err != nil {
		return nil, err
	}

	newFlight, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	fRes := &pb.Flight{
		Date: in.Date,
	}

	err = copier.Copy(&fRes, &newFlight)
	if err != nil {
		return nil, err
	}

	return fRes, nil
}

func (h *FlightHandler) SearchFly(ctx context.Context, in *pb.SearchRequest) (*pb.ListFlights, error) {

	var (
		flights []*models.Flight
		err     error
	)

	if in.Name == "" && in.From == "" && in.To == "" && in.Date == nil {
		return nil, status.Error(codes.InvalidArgument, "Please type one field")
	}

	if in.Name != "" {
		flights, err = h.flightRepository.GetFlightsByName(ctx, in.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}

	} else if in.From != "" {
		flights, err = h.flightRepository.GetFlightsByFrom(ctx, in.From)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	} else if in.To != "" {
		flights, err = h.flightRepository.GetFlightsByDestination(ctx, in.To)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	fRes := &pb.ListFlights{
		Flights: []*pb.Flight{},
	}

	err = copier.CopyWithOption(&fRes.Flights, &flights, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}

	return fRes, nil
}

func (h *FlightHandler) SearchFlyByID(ctx context.Context, in *pb.ID) (*pb.Flight, error) {
	flight, err := h.flightRepository.GetFlightByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}

	fRes := &pb.Flight{
		Date: timestamppb.New(flight.Date),
	}

	err = copier.Copy(&fRes, flight)
	if err != nil {
		return nil, err
	}

	return fRes, nil
}
