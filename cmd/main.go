package main

import (
	"log"

	"github.com/richieoscar/bigshop/cmd/server"
	"github.com/richieoscar/bigshop/config"
)

func main() {
	//init domain
	db, _ := config.InitDB()
	server := server.NewServer(":8085", db)
	var serverErr = server.Run()
	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
