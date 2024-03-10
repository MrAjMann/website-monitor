package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MrAjMann/website-monitor/internal/handler"
	"github.com/MrAjMann/website-monitor/internal/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	initEnv()
	db := initDB()
	defer db.Close()

	websiteRepo := repository.NewWebsiteRepository(db)
	StartStatusChecker(websiteRepo)
	initHTTPHandlers(websiteRepo)

}
func initEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func initDB() *sql.DB {
	log.Println("Connecting to DB...")
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	return db
}

func initHTTPHandlers(websiteRepo *repository.WebsiteRepository) {
	websiteHandler := handler.NewWebsiteHandler(websiteRepo, nil)

	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))
	http.HandleFunc("/", websiteHandler.ServerHome)
	http.HandleFunc("/get-websites/", websiteHandler.GetAllWebsites)
	http.HandleFunc("/get-website-status", websiteHandler.GetWebsiteStatusById)
	http.HandleFunc("/delete-website", websiteHandler.DeleteWebsiteById)
	http.HandleFunc("/add-website/", websiteHandler.AddWebsite)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
