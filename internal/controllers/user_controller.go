package controllers

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userService ps.IUserService
}

func NewUserController(userService ps.IUserService) *userController {
	return &userController{
		userService: userService,
	}
}

// @BasePath /users
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body createUserRequest true "Request Body"
// @Success 201 {object} userResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/ [post]
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

// @BasePath /users
// @Summary Get an user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "The user ID"
// @Success 200 {object} userResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [get]
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

// @BasePath /users
// @Summary Get all users
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} userResponse2
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/ [get]
func (u *userController) ListHandler(c *gin.Context) {
	users, err := u.userService.ListExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, users)
}

// @BasePath /users
// @Summary Authenticate an user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body authUserRequest true "Request Body"
// @Success 200 {object} userResponse3
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/login [post]
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
		return
	}

	sendSuccess(c, http.StatusOK, token)
}

// @BasePath /users
// @Summary Get user by token
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} userResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/me [get]
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
