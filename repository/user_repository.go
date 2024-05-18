package repository

import (
	"database/sql"
	"fmt"
	"github.com/richieoscar/bigshop/domain"
)

type UserRepository struct {
	database *sql.DB
}

type UserRepo interface {
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(*domain.User) (*domain.User, error)
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}

}

func (s *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	rows, err := s.database.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(domain.User) //create user pointer variable
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("error occurred")
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil

}

func (s *UserRepository) GetUserByID(id int) (*domain.User, error) {
	rows, err := s.database.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := new(domain.User) //create user pointer variable
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil

}

func (s *UserRepository) CreateUser(user domain.User) (*domain.User, error) {

	exec, err := s.database.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, fmt.Errorf("error creating user")
	}
	execID, err := exec.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error creating user")
	}
	fmt.Println(execID)
	return &user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*domain.User, error) {
	user := new(domain.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil

}
