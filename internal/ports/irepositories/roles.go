package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
)

type IRoleRepository interface {
	ListRoles(c context.Context) ([]models.Role, error)
}
