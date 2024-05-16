package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/richieoscar/bigshop/service/user"
)

type ApiServer struct {
	address  string
	database *sql.DB
}

func NewApiServer(address string, database *sql.DB) *ApiServer {
	return &ApiServer{
		address:  address,
		database: database,
	}

}

func (server *ApiServer) Run() error {
	var router = mux.NewRouter()
	var subRouter = router.PathPrefix("/api/v1").Subrouter()
	var userStore = user.NewStore(server.database)
	var userHandler = user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on ", server.address)
	return http.ListenAndServe(server.address, router)
}
