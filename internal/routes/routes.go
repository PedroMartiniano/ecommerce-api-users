package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(server *gin.Engine) {
	configs.SwaggerConfigure(docs.SwaggerInfo)

	userRouter := server.Group("/users")
	userRoutes(userRouter)

	roleRouter := server.Group("/roles")
	roleRoutes(roleRouter)

	addressRouter := server.Group("/addresses")
	AddressRoutes(addressRouter)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
