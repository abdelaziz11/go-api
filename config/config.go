package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to connect to database.")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Database ping failed.")
	}

	fmt.Println("Successfully connected to database.")
}

func GetDB() *sql.DB {
	return db
}
