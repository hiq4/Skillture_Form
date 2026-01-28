package entities

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID          uuid.UUID `db:"id" json:"id"`
	FormID      uuid.UUID `db:"form_id" json:"form_id"`
	Email       string    `db:"email,omitempty" json:"email,omitempty"`
	SubmittedAt time.Time `db:"submitted_at" json:"submitted_at"`
}

// TableName returns the DB table name
func (Response) TableName() string {
	return "responses"
}

// IsAnonymous checks if the response was submitted without an email
func (r *Response) IsAnonymous() bool {
	return r.Email == ""
}
