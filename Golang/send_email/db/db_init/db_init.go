package db_init

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Create data directory if not exists
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		os.Mkdir("./data", 0755)
	}

	// Open database with WAL mode for better concurrency
	dbPath := "./data/newsletter.db"
	DB, err = sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_synchronous=NORMAL")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Set connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println("Database initialized successfully with WAL mode")
}

func GetDB() *sql.DB {
	return DB
}
