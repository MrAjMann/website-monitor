package checker

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"log"

	"github.com/joho/godotenv"
)

func CheckWebSiteStatus(url string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file:", err)
		// Consider whether to continue or return on .env load failure depending on your use case.
	}

	if url == "" {
		return "Invalid URL", fmt.Errorf("empty URL provided")
	}

	port := os.Getenv("PORT_NUMBER")
	if port == "" {
		log.Println("PORT_NUMBER environment variable is not set.")
		port = "80" // Assuming a default port if not specified; adjust as necessary.
	}

	domain := strings.TrimPrefix(url, "https://")
	address := domain + ":" + port
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("[DOWN] %s is unreachable, Error: %v", domain, err)
		return "OFFLINE", err
	}
	defer conn.Close()

	log.Printf("[ONLINE] %s is online. Local: %s, Remote: %s", domain, conn.LocalAddr(), conn.RemoteAddr())
	return "ONLINE", nil
}
