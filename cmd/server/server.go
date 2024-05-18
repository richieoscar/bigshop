package server

import (
	"database/sql"
	"github.com/richieoscar/bigshop/handler"
	"github.com/richieoscar/bigshop/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address  string
	database *sql.DB
}

func NewServer(address string, database *sql.DB) *Server {
	return &Server{
		address:  address,
		database: database,
	}

}

func (server *Server) Run() error {
	var router = mux.NewRouter()
	var subRouter = router.PathPrefix("/server/v1").Subrouter()
	var userStore = repository.NewUserRepository(server.database)
	var userHandler = handler.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on ", server.address)
	return http.ListenAndServe(server.address, router)
}
