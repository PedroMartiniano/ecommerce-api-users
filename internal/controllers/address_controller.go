package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/gin-gonic/gin"
)

type AddressController struct {
	addressService ps.IAddressService
}

func NewAddressController(addressService ps.IAddressService) *AddressController {
	return &AddressController{addressService}
}

// @BasePath /addresses
// @Summary Create a new address
// @Security BearerAuth
// @Tags Addresses
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param address body CreateAddressRequest true "Request Body"
// @Success 201 {object} addressResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /addresses/ [post]
func (a *AddressController) CreateHandler(c *gin.Context) {
	var request CreateAddressRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.Param("user_id")
	if userID == "" {
		sendError(c, http.StatusBadRequest, "user_id is required")
		return
	}

	address, err := a.addressService.CreateExecute(c.Request.Context(), models.Address{
		UserID:       userID,
		ZipCode:      request.ZipCode,
		State:        request.State,
		City:         request.City,
		Neighborhood: request.Neighborhood,
		Street:       request.Street,
		Number:       request.Number,
		Complement:   request.Complement,
	})
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusCreated, address)
}

// @BasePath /addresses
// @Summary Get addresses by user ID
// @Security BearerAuth
// @Tags Addresses
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} addressResponse2
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /addresses/user/{user_id} [get]
func (a *AddressController) GetByUserIDHandler(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		sendError(c, http.StatusBadRequest, "user_id is required")
		return
	}

	addresses, err := a.addressService.GetByUserIDExecute(c.Request.Context(), userID)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, addresses)
}

// @BasePath /addresses
// @Summary Get address by ID
// @Security BearerAuth
// @Tags Addresses
// @Accept json
// @Produce json
// @Param address_id path string true "Address ID"
// @Success 200 {object} addressResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /addresses/{address_id} [get]
func (a *AddressController) GetByIDHandler(c *gin.Context) {
	addressID := c.Param("address_id")
	if addressID == "" {
		sendError(c, http.StatusBadRequest, "address_id is required")
		return
	}

	address, err := a.addressService.GetByIDExecute(c.Request.Context(), addressID)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, address)
}
