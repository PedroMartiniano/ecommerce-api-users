package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IAddressesRepository interface {
	Create(context.Context, models.Address) (models.Address, error)
	GetByUserID(context.Context, string) ([]models.Address, error)
	GetByID(context.Context, string) (models.Address, error)
}
