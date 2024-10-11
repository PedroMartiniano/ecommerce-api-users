package adapters

import (
	"database/sql"

	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/services"
)

func NewUserServiceAdapter(db *sql.DB) ps.IUserService {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	return userService
}
