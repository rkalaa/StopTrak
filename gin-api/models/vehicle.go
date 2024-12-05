package models

import "time"

// Vehicle represents a vehicle in the system
type Vehicle struct {
	ID           int       `json:"id"`
	LicensePlate string    `json:"license_plate"`
	Model        string    `json:"model"`
	Manufacturer string    `json:"manufacturer"`
	Year         int       `json:"year"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
