package repositories

import (
	"context"
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	"github.com/google/uuid"
)

type AddressesRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressesRepository {
	return &AddressesRepository{db: db}
}

func (r *AddressesRepository) Create(c context.Context, address models.Address) (models.Address, error) {
	query := `
		INSERT INTO addresses (id, user_id, zip_code, state, city, neighborhood, street, number, complement)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	id, _ := uuid.NewV7()
	address.ID = id.String()

	_, err := r.db.ExecContext(
		c,
		query,
		address.ID,
		address.UserID,
		address.ZipCode,
		address.State,
		address.City,
		address.Neighborhood,
		address.Street,
		address.Number,
		address.Complement,
	)
	if err != nil {
		return models.Address{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return address, nil
}

func (r *AddressesRepository) GetByUserID(c context.Context, userID string) ([]models.Address, error) {
	query := `SELECT id, user_id, zip_code, state, city, neighborhood, street, number, complement FROM addresses WHERE user_id = $1`

	rows, err := r.db.QueryContext(c, query, userID)
	if err != nil {
		return nil, configs.NewError(configs.ErrInternalServer, err)
	}

	defer rows.Close()

	addresses := []models.Address{}
	for rows.Next() {
		var address models.Address
		err = rows.Scan(
			&address.ID,
			&address.UserID,
			&address.ZipCode,
			&address.State,
			&address.City,
			&address.Neighborhood,
			&address.Street,
			&address.Number,
			&address.Complement,
		)
		if err != nil {
			return nil, configs.NewError(configs.ErrInternalServer, err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (r *AddressesRepository) GetByID(c context.Context, id string) (models.Address, error) {
	query := `SELECT id, user_id, zip_code, state, city, neighborhood, street, number, complement FROM addresses WHERE id = $1`

	var address models.Address
	err := r.db.QueryRowContext(c, query, id).Scan(
		&address.ID,
		&address.UserID,
		&address.ZipCode,
		&address.State,
		&address.City,
		&address.Neighborhood,
		&address.Street,
		&address.Number,
		&address.Complement,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Address{}, configs.NewError(configs.ErrNotFound, err)
		}
		return models.Address{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return address, nil
}
