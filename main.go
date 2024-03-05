package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"github.com/MrAjMann/website-monitoring/handlers"
	"time"
)

type Website struct {
	Name   string
	URL    string
	Status string
}

func main() {
	fmt.Println("Server started at http://localhost:8080")
	// Create a file server for static files this includes tailwindcss files
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			log.Fatal(err)
		}
		websites := map[string][]Website{
			"Websites": {
				{Name: "AM Website Solutions", URL: "https://amwebsolutions.com.au", Status: "Online"},
				{Name: "Outback Edge Studio's", URL: "https://outback-edge-studio.vercel.app", Status: "Online"},
			},
		}

		// Execute the template
		tmpl.Execute(w, websites)

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
	http.HandleFunc("/add-website/", AddWebsiteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
