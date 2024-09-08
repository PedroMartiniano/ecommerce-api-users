package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IUserService interface {
	CreateExecute(context.Context, models.User) (models.User, error)
	ListExecute(context.Context) ([]models.User, error)
	GetByIDExecute(context.Context, string) (models.User, error)
	GetByEmailExecute(context.Context, string) (models.User, error)
}
