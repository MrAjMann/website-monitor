package handler

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"strconv"

	"net/http"

	model "github.com/MrAjMann/website-monitor/internal/models"
	"github.com/MrAjMann/website-monitor/internal/repository"
)

type WebsiteHandler struct {
	repo *repository.WebsiteRepository
	tmpl *template.Template
}

func NewWebsiteHandler(repo *repository.WebsiteRepository, tmpl *template.Template) *WebsiteHandler {
	return &WebsiteHandler{repo: repo, tmpl: tmpl}
}

func (h *WebsiteHandler) ServerHome(w http.ResponseWriter, r *http.Request) {
	websites, err := h.repo.GetAllWebsites()
	if err != nil {
		http.Error(w, "Error getting websites", http.StatusInternalServerError)
		log.Printf("Error getting websites: %v\n", err)
		return
	}

	tmpl, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Printf("Error loading template: %v", err)
		return
	}
	err = tmpl.Execute(w, websites)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
		return
	}
}

func (h *WebsiteHandler) AddWebsite(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Method not allowed: %v\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		log.Printf("Error parsing form: %v\n", err)
		return
	}

	website := model.Website{
		Name:   r.PostFormValue("website-title"),
		URL:    r.PostFormValue("website-url"),
		Status: "pending",
	}

	websiteId, err := h.repo.AddWebsite(website)
	if err != nil {
		http.Error(w, "Error adding website", http.StatusInternalServerError)
		log.Printf("Error adding website: %v\n", err)
		return
	}
	htmlTempl := template.Must(template.ParseFiles("src/index.html"))
	websiteIdAsInt, err := strconv.Atoi(websiteId)
	if err != nil {
		http.Error(w, "Error converting website id to int", http.StatusInternalServerError)
		log.Printf("Error converting website id to int: %v\n", err)
		return
	}
	err = htmlTempl.ExecuteTemplate(w, "website-list-element", model.Website{Id: websiteIdAsInt, Name: website.Name, URL: website.URL, Status: website.Status})
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
		return
	}

}

func (h *WebsiteHandler) GetAllWebsites(w http.ResponseWriter, r *http.Request) {
	websites, err := h.repo.GetAllWebsites()
	if err != nil {
		http.Error(w, "Error getting websites", http.StatusInternalServerError)
		log.Printf("Error getting websites: %v\n", err)
		return
	}

	err = h.tmpl.ExecuteTemplate(w, "website-list-element", websites)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}

func (h *WebsiteHandler) GetWebsiteStatusById(w http.ResponseWriter, r *http.Request) {
	// Parse the website ID from the query parameter

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing website ID", http.StatusBadRequest)
		return
	}

	// Convert idParam to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid website ID", http.StatusBadRequest)
		return
	}

	// Fetch the website's current status from the database
	// Assuming you have a function in your repo to do this
	status, err := h.repo.GetWebsiteStatusById(id)
	if err != nil {
		// Log the error and return a server error response
		log.Printf("Failed to get website status for ID %d: %v", id, err)
		http.Error(w, "Failed to get website status", http.StatusInternalServerError)
		return
	}

	// Return the status as a simple HTML snippet or plain text
	// This example returns a span with the status, but you could adjust based on your needs
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<span>%s</span>`, html.EscapeString(status))
}

func (h *WebsiteHandler) DeleteWebsiteById(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing website ID", http.StatusBadRequest)
		return
	}

	// Convert idParam to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid website ID", http.StatusBadRequest)
		return
	}

	_, err = h.repo.DeleteWebsiteById(id)
	if err != nil {
		// Log the error and return a server error response
		log.Printf("Failed to delete website with ID %d: %v", id, err)
		http.Error(w, "Failed to delete website", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
