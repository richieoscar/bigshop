package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/richieoscar/bigshop/types"
	"github.com/richieoscar/bigshop/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(route *mux.Router) {
	route.HandleFunc("/login", h.handleLogin).Methods("POST")
	route.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get json paylaod
	var requestDto types.RegisterRequest //the struct we a serializing the json into a go Object
	err := utils.ParseJson(r, requestDto)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
