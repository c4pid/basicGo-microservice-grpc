package handlers

import (
	"Assignment-micro-service/api/booking-api/requests"
	"Assignment-micro-service/api/booking-api/responses"
	"Assignment-micro-service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler interface {
	Booking(c *gin.Context)
	ViewBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingsClient
}

func NewBookingClient(bookingClient pb.BookingsClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

func (h *bookingHandler) Booking(c *gin.Context) {
	req := requests.CreateBooking{}

	if err := c.ShouldBindJSON(&req); err != nil {
		if validatoErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validatoErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	bReq := &pb.Info{
		BookingDate: timestamppb.New(req.BookingDate),
		CustomerId:  req.CustomerId,
		FlightId:    req.FlightId,
		Status:      req.Status,
	}

	bRes, err := h.bookingClient.Booking(c.Request.Context(), bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		BookingCode: bRes.BookingCode,
		BookingDate: bRes.BookingDate.AsTime(),
		CustomerId:  bRes.CustomerId,
		FlightId:    bRes.FlightId,
		Status:      bRes.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"payload": dto,
	})
}

func (h *bookingHandler) ViewBooking(c *gin.Context) {
	req := requests.ViewBooking{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validatoErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validatoErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	bReq := &pb.ViewRequest{
		BookingCode: req.BookingCode,
	}

	bRes, err := h.bookingClient.ViewBooking(c.Request.Context(), bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.ViewBookingResponse{
		BookingCode: bRes.BookingCode,
		BookingDate: bRes.BookingDate.AsTime(),
		Customer: responses.ViewCustomerResponse{
			Id:      bRes.Customer.Id,
			Name:    bRes.Customer.Name,
			Address: bRes.Customer.Address,
			Email:   bRes.Customer.Email,
		},
		Flight: responses.ViewFlightResponse{
			Id:            int(bRes.Flight.Id),
			Name:          bRes.Flight.Name,
			From:          bRes.Flight.From,
			To:            bRes.Flight.To,
			Date:          bRes.Flight.Date.AsTime(),
			Status:        bRes.Flight.Status,
			AvailableSlot: bRes.Flight.AvailableSlot,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	req := requests.ViewBooking{}

	if err := c.ShouldBindJSON(&req); err != nil {
		if validatoErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validatoErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	id := &pb.ViewRequest{
		BookingCode: req.BookingCode,
	}
	bRes, err := h.bookingClient.CancleBooking(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		BookingCode: bRes.BookingCode,
		BookingDate: bRes.BookingDate.AsTime(),
		CustomerId:  bRes.CustomerId,
		FlightId:    bRes.FlightId,
		Status:      bRes.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}
