package mocks

import "github.com/richieoscar/bigshop/dto"

type mockUserStore struct{}

func NewUserStorMock() mockUserStore {
	return mockUserStore{}
}

func (s mockUserStore) GetUserByEmail(email string) (*dto.User, error) {
	return new(dto.User), nil

}

func (s mockUserStore) GetUserByID(id int) (*dto.User, error) {
	return new(dto.User), nil

}

func (s mockUserStore) CreateUser(dto.User) (*dto.User, error) {
	return new(dto.User), nil

}
