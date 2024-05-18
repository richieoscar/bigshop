package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func newMysqlStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func InitDB() (*sql.DB, error) {
	db, dbErr := newMysqlStorage(mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	db, err := connectDb(db)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil

}

func connectDb(db *sql.DB) (*sql.DB, error) {
	connectionErr := db.Ping()
	if connectionErr != nil {
		log.Fatal(connectionErr)
		return nil, nil
	}
	log.Println("Connected to Database")
	return db, nil

}
