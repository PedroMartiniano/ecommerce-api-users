package repositories

import (
	"context"
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/irepositories"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) pr.IRoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) ListRoles(c context.Context) ([]models.Role, error) {
	query := `SELECT id, name FROM roles;`

	rows, err := r.db.QueryContext(c, query)
	if err != nil {
		return []models.Role{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return []models.Role{}, configs.NewError(configs.ErrInternalServer, err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}
