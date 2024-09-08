package services

import (
	"context"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/ports/services"
)

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) services.IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateExecute(c context.Context, user models.User) (models.User, error) {
	newUser, err := u.userRepository.Create(c, user)

	return newUser, err
}

func (u *userService) GetByIDExecute(c context.Context, userID string) (models.User, error) {
	user, err := u.userRepository.GetByID(c, userID)

	return user, err
}

func (u *userService) ListExecute(c context.Context) ([]models.User, error) {
	users, err := u.userRepository.List(c)

	return users, err
}

func (u *userService) GetByEmailExecute(ctx context.Context, s string) (models.User, error) {
	user, err := u.userRepository.GetByEmail(ctx, s)

	return user, err
}
