package model

import (
	"time"
)

// User struct
type User struct {
	ID           string
	FullName     string
	EmailAddress string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
