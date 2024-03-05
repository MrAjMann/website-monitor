package handlers

import (
	"fmt"
	"html/template"
	model "htmx-webserver/internal/models"
	"net/http"
)

type websiteHandler struct {
	
}

func (h *websiteHandler) AddWebsiteHandler(w http.ResponseWriter, r *http.Request) {

	website := r.PostFormValue("website-title")
	url := r.PostFormValue("website-url")
	status := "pending"
	// domain := strings.TrimPrefix(url, "https://")

	htmlTempl := template.Must(template.ParseFiles("src/index.html"))
	// status := h.Check(domain, "80")
	fmt.Println(website, url, status)

	htmlTempl.ExecuteTemplate(w, "website-list-element", model.Website{Name: website, URL: url, Status: status})

}
