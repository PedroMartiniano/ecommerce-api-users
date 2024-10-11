package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/controllers"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func roleRoutes(router *gin.RouterGroup) {
	roleService := adapters.NewRoleServiceAdapter(configs.DB)
	roleController := controllers.NewRoleController(roleService)

	router.GET("/", middlewares.VerifyJWT, roleController.ListHandler)
}
