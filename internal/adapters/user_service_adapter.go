package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-users/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/services"
	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/services"
)
func NewUserServiceAdapter() ps.IUserService {
	userRepository := repositories.NewUseRepository()
	userService := services.NewUserService(userRepository)

	return userService
}