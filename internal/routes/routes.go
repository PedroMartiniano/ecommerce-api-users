package routes

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {
	userRouter := server.Group("/users")
	UserRoutes(userRouter)
}