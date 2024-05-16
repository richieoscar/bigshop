package mocks

import "github.com/richieoscar/bigshop/types"

type mockUserStore struct{}

func NewUserStorMock() mockUserStore{
	return mockUserStore{}
}

func (s mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return new(types.User), nil

}

func (s mockUserStore) GetUserByID(id int) (*types.User, error) {
	return new(types.User), nil

}

func (s mockUserStore) CreateUser(types.User) (*types.User, error) {
	return new(types.User), nil

}