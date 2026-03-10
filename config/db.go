package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dbURL := os.Getenv("DB_URL")

	DB, err := sql.opne("postgres", dbURL)
	if err != nil {
		log.Fatal("Error Opening the database: ", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("database connection los.Cannot connect to the Database:", err)
	}
	fmt.Println("Connected To The Database Successfully.\n")
}
