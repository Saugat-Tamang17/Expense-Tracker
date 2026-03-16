package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", dsn)

	//if error occurs here, that means the driver is invalid or the dsn isnt parseble //
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	//if error has happened here , that means the database is offline or the connection lost heheh//
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to the db : %v", err)
	}

	log.Println("Database connected successfully")
	return db
}
