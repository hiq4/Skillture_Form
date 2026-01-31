package entities

import (
	"time"

	"github.com/google/uuid"
)

// Response represents a single form submission by a user
type Response struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	FormID      uuid.UUID      `db:"form_id" json:"form_id"`
	Respondent  map[string]any `db:"respondent" json:"respondent"` // JSONB: {"email": "...", "name": "...", "phone": "..."}
	SubmittedAt time.Time      `db:"submitted_at" json:"submitted_at"`
}

// TableName returns the DB table name
func (Response) TableName() string {
	return "responses"
}

// GetEmail returns the email of the respondent if exists
func (r *Response) GetEmail() string {
	if email, ok := r.Respondent["email"].(string); ok {
		return email
	}
	return ""
}

// GetName returns the name of the respondent if exists
func (r *Response) GetName() string {
	if name, ok := r.Respondent["name"].(string); ok {
		return name
	}
	return ""
}

// SetEmail sets or updates the email in the respondent JSON
func (r *Response) SetEmail(email string) {
	if r.Respondent == nil {
		r.Respondent = make(map[string]any)
	}
	r.Respondent["email"] = email
}

// SetName sets or updates the name in the respondent JSON
func (r *Response) SetName(name string) {
	if r.Respondent == nil {
		r.Respondent = make(map[string]any)
	}
	r.Respondent["name"] = name
}
