package entities

import (
	"time"

	"github.com/google/uuid"
)

type Forms struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      int       `db:"status" json:"status"`
	CreatedAt   time.Time `db:"creat_at" json:"creat_at"`
}

// TableName returns the DB table name

func (Forms) TableName() string {
	return "forms"
}

// IsActive checks if the form is active

func (f *Forms) IsActive() bool {
	return f.Status == 1
}

// Deactivate marks the form as inactive
func (f *Forms) Deactivate() {
	f.Status = 0
}
