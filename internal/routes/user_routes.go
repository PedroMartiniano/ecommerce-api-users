package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/controllers"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	userService := adapters.NewUserServiceAdapter(configs.DB)
	userController := controllers.NewUserController(userService)

	router.POST("/", userController.CreateHandler)
	router.GET("/", userController.ListHandler)
	router.GET("/:id", userController.GetByIDHandler)
	router.GET("/me", middlewares.VerifyJWT, userController.GetMeHandler)
	router.POST("/login", userController.AuthUserHandler)
}
