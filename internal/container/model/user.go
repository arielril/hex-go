package model

import (
	"time"
)

type User struct {
	ID           string
	FullName     string
	EmailAddress string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
