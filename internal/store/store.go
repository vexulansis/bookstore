package store

import (
	"database/sql"
	"fmt"
	"os"
	. "project/internal"

	"github.com/joho/godotenv"
)

type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}
func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		FatalLog("Error loading .env file")
	}
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		FatalLog("Error opening database")
	}
	InfoLog("Database connected")
	return db
}
