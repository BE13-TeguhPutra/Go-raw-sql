package config

import (
	"database/sql"
	"fmt"
	"log"
)

func Connection() *sql.DB {
	var connectionString = "root:Teguh12345@tcp(127.0.0.1:3306)/db_be13_teguh"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("Error connect to db", errPing.Error())
	} else {
		fmt.Println("Koneksi berhasil")
	}
	return db

}
