package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Connecting to DB...")

	if err != nil {
		log.Printf("Database connection error: %v", err)
		return
	}

	databaseURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	createWebsitesTableSQL := `
			CREATE TABLE IF NOT EXISTS websites (
					id SERIAL PRIMARY KEY,
					name TEXT,
					url TEXT,
					status TEXT,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);
`

	_, err = db.Exec(createWebsitesTableSQL)
	if err != nil {
		log.Fatal("Error creating websites table: ", err)
	}

}
