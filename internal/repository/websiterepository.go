package repository

import (
	"database/sql"
	"log"

	model "github.com/MrAjMann/website-monitor/internal/models"
)

type WebsiteRepository struct {
	db *sql.DB
}

func NewWebsiteRepository(db *sql.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (repo *WebsiteRepository) GetAllWebsites() ([]model.Website, error) {
	// GetAllWebsites returns all the websites from the database
	rows, err := repo.db.Query("SELECT id, name, url, status FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	websites := []model.Website{}
	for rows.Next() {
		website := model.Website{}
		err := rows.Scan(&website.Id, &website.Name, &website.URL, &website.Status)
		if err != nil {
			return nil, err
		}
		websites = append(websites, website)
	}
	return websites, nil
}

func (repo *WebsiteRepository) AddWebsite(website model.Website) (string, error) {
	// AddWebsite adds a new website to the database
	var websiteId string
	err := repo.db.QueryRow("INSERT INTO websites (name, url, status) VALUES ($1, $2, $3) RETURNING id", website.Name, website.URL, website.Status).Scan(&websiteId)

	if err != nil {
		return "ERROR  IN REPO", err
	}
	return websiteId, nil
}

func (repo *WebsiteRepository) UpdateWebsiteStatusById(status string, id int) (string, int, error) {
	// UpdateWebsiteStatus updates the status of a website in the database
	log.Printf("Website Status: %v\n", status)
	log.Printf("Website id: %v\n", id)
	result, err := repo.db.Exec("UPDATE websites SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		log.Printf("Error updating website status: %v\n", err)
		return "ERROR IN REPO", id, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v\n", err)
		return "ERROR IN REPO", id, err
	}
	if rowsAffected == 0 {
		log.Printf("No rows updated")
		return "NO UPDATE", id, nil // Or handle as you see fit
	}
	return "UPDATED", id, nil
}

func (repo *WebsiteRepository) GetWebsiteStatusById(id int) (string, error) {
	var status string
	err := repo.db.QueryRow("SELECT status FROM websites WHERE id = $1", id).Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (repo *WebsiteRepository) DeleteWebsiteById(Id int) (int, error) {
	// DeleteWebsiteById deletes a website from the database
	result, err := repo.db.Exec("DELETE FROM websites WHERE id = $1", Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	log.Printf("Rows affected: %v\n", rowsAffected)
	return int(rowsAffected), nil

}
