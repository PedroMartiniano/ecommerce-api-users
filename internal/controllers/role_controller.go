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

func (u *roleController) ListHandler(c *gin.Context) {
	users, err := u.roleService.ListExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, users)
}
