package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/MrAjMann/website-monitor/internal/handler"
	"github.com/MrAjMann/website-monitor/internal/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Connecting to DB...")

	// Database connection code
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
	// Create a file server for static files this includes tailwindcss files

	// Templates
	websiteRepo := repository.NewWebsiteRepository(db)
	if websiteRepo == nil {
		println("Creating customers table")
	}

	websitehandler := handler.NewWebsiteHandler(websiteRepo, nil)

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			log.Fatal(err)
		}

		// Execute the template
		tmpl.Execute(w, nil)

	}

	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-website/", websitehandler.AddWebsite)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
