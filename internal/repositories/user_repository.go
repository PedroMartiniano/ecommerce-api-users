package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/utils"
	"time"

	"github.com/PedroMartiniano/ecommerce-api-users/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-users/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-users/internal/ports/repositories"
	"github.com/google/uuid"
)

type userRepository struct{}

func NewUseRepository() pr.IUserRepository {
	return &userRepository{}
}

func (u *userRepository) Create(c context.Context, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (id, role_id, email, password, name, phone, cpf, birth_date, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.Status = true

	hashPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		return models.User{}, configs.NewError(configs.ErrInternalServer, err)
	}
	_, err = configs.DB.ExecContext(
		c,
		query,
		user.ID,
		user.RoleID,
		user.Email,
		hashPassword,
		user.Name,
		user.Phone,
		user.CPF,
		user.BirthDate,
		user.Status,
		user.CreatedAt,
	)
	if err != nil {
		return models.User{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return user, nil
}

func (u *userRepository) GetByID(c context.Context, userID string) (models.User, error) {
	query := `
		SELECT id, role_id, address_id, email, password, name, phone, cpf, birth_date, status, created_at
		FROM users
		WHERE id = $1;
	`

	var user models.User
	var addrID sql.NullString

	err := configs.DB.QueryRowContext(c, query, userID).Scan(
		&user.ID,
		&user.RoleID,
		&addrID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Phone,
		&user.CPF,
		&user.BirthDate,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, configs.NewError(configs.ErrNotFound, err)
		}
		return models.User{}, configs.NewError(configs.ErrInternalServer, err)
	}

	if addrID.Valid {
		user.AddrID = addrID.String
	}

	return user, nil
}

func (u *userRepository) List(c context.Context) ([]models.User, error) {
	query := `
		SELECT id, role_id, address_id, email, password, name, phone, cpf, birth_date, status, created_at
		FROM users;
	`

	rows, err := configs.DB.QueryContext(c, query)
	if err != nil {
		return []models.User{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		var addrID sql.NullString

		err := rows.Scan(
			&user.ID,
			&user.RoleID,
			&addrID,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Phone,
			&user.CPF,
			&user.BirthDate,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			return []models.User{}, configs.NewError(configs.ErrInternalServer, err)
		}

		if addrID.Valid {
			user.AddrID = addrID.String
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetByEmail(c context.Context, email string) (models.User, error) {
	query := `
		SELECT id, role_id, address_id, email, password, name, phone, cpf, birth_date, status, created_at
		FROM users
		WHERE email = $1;
	`
	var user models.User
	var addrID sql.NullString

	err := configs.DB.QueryRowContext(c, query, email).Scan(
		&user.ID,
		&user.RoleID,
		&addrID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Phone,
		&user.CPF,
		&user.BirthDate,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, configs.NewError(configs.ErrNotFound, err)
		}
		return models.User{}, configs.NewError(configs.ErrInternalServer, err)
	}

	if addrID.Valid {
		user.AddrID = addrID.String
	}

	return user, nil
}

func (u *userRepository) GetUsersRoles(c context.Context) ([]models.Role, error) {
	query := `
		SELECT id, name
		FROM roles;
	`

	rows, err := configs.DB.QueryContext(c, query)
	if err != nil {
		return []models.Role{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(
			&role.ID,
			&role.Name,
		)
		if err != nil {
			return []models.Role{}, configs.NewError(configs.ErrInternalServer, err)
		}

		roles = append(roles, role)
	}
	return roles, nil
}
