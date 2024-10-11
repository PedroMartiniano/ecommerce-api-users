package controllers

import (
	"net/http"

	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleService ps.IRoleService
}

func NewRoleController(roleService ps.IRoleService) *roleController {
	return &roleController{
		roleService: roleService,
	}
}

// @BasePath /roles
// @Summary Get all roles
// @Security BearerAuth
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {object} roleResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /roles/ [get]
func (u *roleController) ListHandler(c *gin.Context) {
	users, err := u.roleService.ListExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, users)
}
