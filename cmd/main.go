package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/richieoscar/bigshop/cmd/api"
	"github.com/richieoscar/bigshop/config"
	"github.com/richieoscar/bigshop/db"
)

func main() {
	//init db
	db, dbErr := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	connecDb(db)

	server := api.NewApiServer(":8080", db)
	var serverErr = server.Run()
	if serverErr != nil {
		log.Fatal(serverErr)
	}

}

func connecDb(db *sql.DB) {
	connectionErr := db.Ping()
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}
	log.Println("Connected to Datbase")

}
