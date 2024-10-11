package services

import (
	"context"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/irepositories"
	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
)

type roleService struct {
	roleRepository pr.IRoleRepository
}

func NewRoleService(roleRepository pr.IRoleRepository) ps.IRoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (u *roleService) ListExecute(c context.Context) ([]models.Role, error) {
	roles, err := u.roleRepository.ListRoles(c)

	return roles, err
}
