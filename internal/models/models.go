package models

import "time"

type User struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CPF       string    `json:"cpf"`
	BirthDate time.Time `json:"birth_date"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	ID           string `json:"id"`
	ZipCode      string `json:"zip_code"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
}
