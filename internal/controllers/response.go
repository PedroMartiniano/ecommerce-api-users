package controllers

import (
	"errors"
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	"github.com/gin-gonic/gin"
)

var logger = configs.GetLogger()

func httpError(err error) (code int, message string) {
	var myErr configs.Error

	if errors.As(err, &myErr) {
		message = myErr.AppError().Error()

		switch myErr.TypeError() {
		case configs.ErrInternalServer:
			code = http.StatusInternalServerError
		case configs.ErrNotFound:
			code = http.StatusNotFound
		case configs.ErrBadRequest:
			code = http.StatusBadRequest
		}
	} else {
		code = http.StatusInternalServerError
		message = configs.ErrInternalServer.Error()
	}

	return
}

func sendError(c *gin.Context, code int, err string) {
	logger.Error(err)
	c.JSON(code, gin.H{"success": false, "message": err})
}

func sendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{"success": true, "data": data})
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type userResponse1 struct {
	Success bool        `json:"success"`
	Data    models.User `json:"data"`
}

type userResponse2 struct {
	Success bool          `json:"success"`
	Data    []models.User `json:"data"`
}

type userResponse3 struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

type roleResponse1 struct {
	Success bool        `json:"success"`
	Data    models.Role `json:"data"`
}

type addressResponse1 struct {
	Success bool        `json:"success"`
	Data    models.Address `json:"data"`
}

type addressResponse2 struct {
	Success bool           `json:"success"`
	Data    []models.Address `json:"data"`
}

