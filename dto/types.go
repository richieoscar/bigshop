package dto

import "github.com/richieoscar/bigshop/domain"

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(domain.User) (*domain.User, error)
}

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"passsword" validate:"required"`
}
