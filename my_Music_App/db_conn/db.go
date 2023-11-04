package db_conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error
var db *sql.DB

// InitDB initializes the database connection and returns a pointer to the DB object.

func InitDB() (*sql.DB, error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error in loading .env file :%s", err.Error())
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatalf("Error in connecting to Database :%s", err.Error())
	}

	fmt.Println("Database Connection Successfully")

	SongInfoTable()

	return db, nil
}

// SongInfoTable creates the 'songInfo' table if it doesn't exist.

func SongInfoTable() {
	_, err := db.Exec(`CREATE TABLE  IF NOT EXISTS songInfo(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100), artists VARCHAR(100), genre VARCHAR(100), publishyear INT, language VARCHAR(50))`)
	if err != nil {
		log.Fatalf("Error in Creating Table songInfo :%s", err.Error())
	}

	fmt.Println("songInfo Table ceatred Successfully")
}
