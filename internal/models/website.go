package model

import (
	"time"
)

type Website struct {
	Id        string
	Name      string
	URL       string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
