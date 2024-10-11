package routes

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {
	userRouter := server.Group("/users")
	userRoutes(userRouter)

	roleRouter := server.Group("/roles")
	roleRoutes(roleRouter)
}
