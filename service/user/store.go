package user

import (
	"database/sql"
	"fmt"

	"github.com/richieoscar/bigshop/types"
)

type Store struct {
	database *sql.DB
}

func NewStore(database *sql.DB) *Store {
	return &Store{
		database: database,
	}

}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.database.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.User) //create user pointer variable
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

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.database.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := new(types.User) //create user pointer variable
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

func (s *Store) CreateUser(user types.User) (*types.User, error) {

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

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

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
