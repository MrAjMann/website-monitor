package handler

import (
	"html/template"
	"log"

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
	// domain := strings.TrimPrefix(url, "https://")
	websiteId, err := h.repo.AddWebsite(website)
	if err != nil {
		http.Error(w, "Error adding website", http.StatusInternalServerError)
		log.Printf("Error adding website: %v\n", err)
		return
	}
	htmlTempl := template.Must(template.ParseFiles("src/index.html"))

	err = htmlTempl.ExecuteTemplate(w, "website-list-element", model.Website{Id: websiteId, Name: website.Name, URL: website.URL, Status: website.Status})

}
