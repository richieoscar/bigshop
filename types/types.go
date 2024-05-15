package types

type RegisterRequest struct {
	FirsName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
