package middlewares

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyJWT(c *gin.Context) {
	userID, err := utils.VerifyJWT(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
