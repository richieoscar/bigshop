package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
}

func NewMysqlStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db,nil
}
