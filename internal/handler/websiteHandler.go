package handler

import (
	"fmt"
	"html/template"
	model "website-monitor/internal/models"
	"website-monitor/internal/repository"

	"net/http"
)

type WebsiteHandler struct {
	repo *repository.WebsiteRepository
}

func (h *WebsiteHandler) AddWebsite(w http.ResponseWriter, r *http.Request) {

	website := r.PostFormValue("website-title")
	url := r.PostFormValue("website-url")
	status := "pending"
	// domain := strings.TrimPrefix(url, "https://")

	htmlTempl := template.Must(template.ParseFiles("src/index.html"))
	// status := h.Check(domain, "80")
	fmt.Println(website, url, status)

	htmlTempl.ExecuteTemplate(w, "website-list-element", model.Website{Name: website, URL: url, Status: status})

}
