package adapters

import (
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/services"
)

func NewAddressServiceAdapter(db *sql.DB) iservices.IAddressService {
	addressRepository := repositories.NewAddressRepository(db)

	addressService := services.NewAddressService(addressRepository)
	return addressService
}
