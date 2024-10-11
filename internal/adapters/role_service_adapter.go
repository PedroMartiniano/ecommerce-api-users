package adapters

import (
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/services"
)

func NewRoleServiceAdapter(db *sql.DB) iservices.IRoleService {
	roleRepository := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepository)

	return roleService
}