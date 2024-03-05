package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/MrAjMann/website-monitor/internal/handler/websitehandler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Connecting to DB...")
	databaseURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Error connecting to the DB", err)
	}
	defer db.Close()
	// Create a file server for static files this includes tailwindcss files
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			log.Fatal(err)
		}

		// Execute the template
		tmpl.Execute(w, nil)

	}

	Check := func(destination string, port string) string {

		address := destination + ":" + port
		fmt.Println("Checking ", address)
		timeout := time.Duration(5 * time.Second)
		conn, err := net.DialTimeout("tcp", address, timeout)
		var status string

		if err != nil {
			status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
		} else {
			status = fmt.Sprintf("ONLINE %v is online, \n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
		}

		fmt.Println("des", address)
		fmt.Println("status", status)
		return status
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-website/", websitehandler.AddWebsite)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
