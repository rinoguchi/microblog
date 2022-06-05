package models

import "time"

type Comment struct {
	ID        int
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
