package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/controllers"
	"github.com/gin-gonic/gin"
)

func AddressRoutes(r *gin.RouterGroup) {
	addressService := adapters.NewAddressServiceAdapter(configs.DB)
	addressController := controllers.NewAddressController(addressService)

	r.POST("/:user_id", addressController.CreateHandler)
	r.GET("/user/:user_id", addressController.GetByUserIDHandler)
	r.GET("/:address_id", addressController.GetByIDHandler)
}
