package main // ensure this matches the package name of your main.go

import (
	"fmt"
	"log"
	"time"

	// Assuming "checker" and "repo" are properly imported or accessible
	"github.com/MrAjMann/website-monitor/internal/checker"
	"github.com/MrAjMann/website-monitor/internal/repository"
)

func StartStatusChecker(repo *repository.WebsiteRepository) {

	go func() {
		ticker := time.NewTicker(60 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			websites, err := repo.GetAllWebsites()
			log.Printf("Websites: %v\n", websites)
			if err != nil {
				log.Printf("Error getting websites: %v\n", err)
				continue
			}

			for _, website := range websites {
				status, err := checker.CheckWebSiteStatus(website.URL)
				if err != nil {
					fmt.Printf("Error checking status for %s: %v\n", website.URL, err)
					status = "OFFLINE" // Default status on error
				}

				_, _, err = repo.UpdateWebsiteStatusById(status, website.Id)
				if err != nil {
					fmt.Printf("Error updating website status for %s: %v\n", website.URL, err)
				} else {
					fmt.Printf("Successfully updated status for %s to %s\n", website.URL, status)
				}
			}
		}
	}()
}
