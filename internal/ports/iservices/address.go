package iservices

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IAddressService interface {
	CreateExecute(context.Context, models.Address) (models.Address, error)
	GetByUserIDExecute(context.Context, string) ([]models.Address, error)
	GetByIDExecute(context.Context, string) (models.Address, error)
}
