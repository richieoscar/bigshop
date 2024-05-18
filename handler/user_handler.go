package handler

import (
	"fmt"
	"github.com/richieoscar/bigshop/domain"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/richieoscar/bigshop/dto"
	"github.com/richieoscar/bigshop/utils"
)

type Handler struct {
	store dto.UserRepository
}

func NewHandler(store dto.UserRepository) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(route *mux.Router) {
	route.HandleFunc("/login", h.handleLogin).Methods("POST")
	route.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	//get json paylaod
	var requestDto dto.RegisterRequest //the struct we a serializing the json into a go Object
	err := utils.ParseJson(r, &requestDto)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = h.store.GetUserByEmail(requestDto.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	newUser := createUser(requestDto)

	u, err := h.store.CreateUser(newUser)
	if err != nil {
		log.Fatal(err)
		utils.WriteError(w, http.StatusOK, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, u)

}

func createUser(request dto.RegisterRequest) domain.User {
	hasPass, err := hash(request.Password)
	if err != nil {
		return domain.User{}
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  hasPass,
	}
	return user

}

func hash(password string) (string, error) {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}
	return hash, nil
}
