package controllers

import "time"

type createUserRequest struct {
	RoleID    string    `json:"role_id" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	CPF       string    `json:"cpf" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

type authUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateAddressRequest struct {
	ZipCode      string `json:"zip_code" binding:"required"`
	State        string `json:"state" binding:"required"`
	City         string `json:"city" binding:"required"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	Street       string `json:"street" binding:"required"`
	Number       string `json:"number" binding:"required"`
	Complement   string `json:"complement" binding:"required"`
}
