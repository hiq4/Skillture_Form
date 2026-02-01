package entities

import (
	"Skillture_Form/internal/domain/enums"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Forms struct {
	ID          uuid.UUID        `db:"id" json:"id"`
	Title       string           `db:"title" json:"title"`
	Description string           `db:"description" json:"description"`
	Status      enums.FormStatus `db:"status" json:"status"`
	CreatedAt   time.Time        `db:"creat_at" json:"creat_at"`
}

var ErrInvalidFormStatus = errors.New("invalid form status")

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

// IsValid validates domain rules
func (f *Forms) IsValid() error {
	if !f.Status.IsValid() {
		return ErrInvalidFormStatus
	}
	return nil
}
