package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MAX_OPEN_CONN, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	if err != nil {
		MAX_OPEN_CONN = 20
	}
	MAX_IDLE_CONN, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	if err != nil {
		MAX_IDLE_CONN = 10
	}
	MAX_CONN_LIFETIME, err := strconv.Atoi(os.Getenv("MAX_CONN_LIFETIME"))
	if err != nil {
		MAX_CONN_LIFETIME = 5
	}

	dsn := fmt.Sprintf("host=%s password=%s port=%s user=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to open database connection: %v\n", err)
	}

	db.SetMaxOpenConns(MAX_OPEN_CONN)

	db.SetMaxIdleConns(MAX_IDLE_CONN)

	db.SetConnMaxLifetime(time.Duration(MAX_CONN_LIFETIME) * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB = db
	fmt.Println("Connected to the database!")
}
