package controllers

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/ports/services"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) CreateHandler(c *gin.Context) {
	var request createUserRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Email:     request.Email,
		Password:  request.Password,
		RoleID:    request.RoleID,
		Name:      request.Name,
		Phone:     request.Phone,
		CPF:       request.CPF,
		BirthDate: request.BirthDate,
	}

	newUser, err := u.userService.CreateExecute(c.Request.Context(), user)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusCreated, newUser)
}

func (u *userController) GetByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, "missing id parameter")
		return
	}

	user, err := u.userService.GetByIDExecute(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, user)
}

func (u *userController) ListHandler(c *gin.Context) {
	users, err := u.userService.ListExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, users)
}

func (u *userController) AuthUserHandler(c *gin.Context) {
	var request authUserRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.userService.GetByEmailExecute(c.Request.Context(), request.Email)
	if err != nil {
		sendError(c, http.StatusUnauthorized, "Email or password incorrect")
		return
	}

	err = utils.CompareHashAndPassword(user.Password, request.Password)
	if err != nil {
		sendError(c, http.StatusUnauthorized, "Email or password incorrect")
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
	}

	sendSuccess(c, http.StatusOK, token)
}

func (u *userController) GetMeHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		sendError(c, http.StatusUnauthorized, "You need to be logged in")
		return
	}

	strUserID, ok := userID.(string)
	if !ok {
		sendError(c, http.StatusUnauthorized, "You need to be logged in")
		return
	}
	user, err := u.userService.GetByIDExecute(c.Request.Context(), strUserID)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, user)
}
