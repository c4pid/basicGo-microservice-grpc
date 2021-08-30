package handlers

import (
	"Assignment-micro-service/api/customer-api/requests"
	"Assignment-micro-service/api/customer-api/responses"
	"Assignment-micro-service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.CustomerServiceClient
}

func NewCustomerHandler(customerClient pb.CustomerServiceClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := requests.CreateCustomerRequest{}

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

	cReq := &pb.Customer{
		Name:        req.Name,
		Address:     req.Address,
		LicenseId:   req.LicenseId,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
	}

	cRes, err := h.customerClient.CreateCustomer(c.Request.Context(), cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		Id:          cRes.Id,
		Name:        cRes.Name,
		Password:    cRes.Password,
		Address:     cRes.Address,
		LicenseId:   cRes.LicenseId,
		PhoneNumber: cRes.PhoneNumber,
		Email:       cRes.Email,
		Active:      cRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"payload": dto,
	})
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	req := requests.UpdateCustomerRequest{}
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

	cReq := &pb.Customer{}

	err := copier.Copy(&cReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	cRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		Id:          req.Id,
		Name:        cRes.Name,
		Address:     cRes.Address,
		LicenseId:   cRes.LicenseId,
		PhoneNumber: cRes.PhoneNumber,
		Email:       cRes.Email,
		Active:      cRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"payload": dto,
	})

}

func (h *customerHandler) ChangePassword(c *gin.Context) {
	req := requests.ChangePasswordRequest{}
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

	passReq := &pb.ChangePasswordRequest{}
	err := copier.Copy(&passReq, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	_, err = h.customerClient.ChangePassword(c.Request.Context(), passReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
