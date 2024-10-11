package iservices

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IRoleService interface {
	ListExecute(context.Context) ([]models.Role, error)
}
