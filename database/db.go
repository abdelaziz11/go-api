package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
	"database/sql"
	"github.com/lib/pq"
)

var db *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database.")
}

func GetDB() *sql.DB {
	return db
}
