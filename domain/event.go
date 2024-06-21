package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       string    `json:"rating"`
	ImageURL     string    `json:"image_url"`
	Price        float64   `json:"price"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var alias struct {
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		Organization string  `json:"organization"`
		Date         string  `json:"date"`
		Price        float64 `json:"price"`
		Rating       string  `json:"rating"`
		ImageURL     string  `json:"image_url"`
		CreatedAt    string  `json:"created_at"`
		Location     string  `json:"location"`
	}

	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	e.ID = alias.ID
	e.Name = alias.Name
	e.Organization = alias.Organization
	e.Price = alias.Price
	e.Rating = alias.Rating
	e.ImageURL = alias.ImageURL
	e.Location = alias.Location

	var err error
	e.Date, err = time.Parse("2006-01-02T15:04:05", alias.Date)
	if err != nil {
		return fmt.Errorf("error parsing Date: %w", err)
	}

	e.CreatedAt, err = time.Parse("2006-01-02T15:04:05", alias.CreatedAt)
	if err != nil {
		return fmt.Errorf("error parsing CreatedAt: %w", err)
	}

	return nil
}
