package handlers

import (
	"Assignment-micro-service/api/flight-api/requests"
	"Assignment-micro-service/api/flight-api/responses"
	"Assignment-micro-service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	SearchFly(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightsClient
}

func NewFlightHandler(flightClient pb.FlightsClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}

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

	fReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(req.Date),
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	fRes, err := h.flightClient.CreateFlight(c.Request.Context(), fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		Id:            fRes.Id,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvailableSlot: fRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := requests.UpdateFlightRequest{}

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

	fReq := pb.Flight{
		Date: timestamppb.New(req.Date),
	}

	err := copier.Copy(&fReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	fRes, err := h.flightClient.UpdateFlight(c.Request.Context(), &fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		Date: fRes.Date.AsTime(),
	}

	err = copier.Copy(&dto, &fRes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}

func (h *flightHandler) SearchFly(c *gin.Context) {
	req := requests.SearchRequest{}

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

	fReq := pb.SearchRequest{}

	err := copier.Copy(&fReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	fRes, err := h.flightClient.SearchFly(c.Request.Context(), &fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.ListFlights{}

	err = copier.CopyWithOption(&dto, &fRes, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": dto,
		"status":  http.StatusOK,
	})
}
