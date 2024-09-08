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
