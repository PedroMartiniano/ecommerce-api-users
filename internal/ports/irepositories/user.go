package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IUserRepository interface {
	Create(context.Context, models.User) (models.User, error)
	List(context.Context) ([]models.User, error)
	GetByID(context.Context, string) (models.User, error)
	GetByEmail(context.Context, string) (models.User, error)
}
