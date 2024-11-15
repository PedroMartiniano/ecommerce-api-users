package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/irepositories"
	ps "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/iservices"
)

type AddressService struct {
	addressesRepository pr.IAddressesRepository
}

func NewAddressService(addressesRepository pr.IAddressesRepository) ps.IAddressService {
	return &AddressService{
		addressesRepository,
	}
}

func (s *AddressService) CreateExecute(c context.Context, address models.Address) (models.Address, error) {
	return s.addressesRepository.Create(c, address)
}

func (s *AddressService) GetByUserIDExecute(c context.Context, userID string) ([]models.Address, error) {
	return s.addressesRepository.GetByUserID(c, userID)
}

func (s *AddressService) GetByIDExecute(c context.Context, id string) (models.Address, error) {
	return s.addressesRepository.GetByID(c, id)
}

