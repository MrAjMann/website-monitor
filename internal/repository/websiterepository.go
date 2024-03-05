package repository

import (
	"database/sql"
	model "website-monitor/internal/models"
)

type WebsiteRepository struct {
	db *sql.DB
}

func (repo *WebsiteRepository) AddWebsite(website *model.Website) (string, error) {
	// AddWebsite adds a new website to the database
	var websiteId string
	err := repo.db.QueryRow("INSERT INTO websites (name, url, status) VALUES ($1, $2, $3) RETURNING id", website.Name, website.URL, website.Status).Scan(&websiteId)

	if err != nil {
		return "ERROR  IN REPO", err
	}
	return websiteId, nil
}
