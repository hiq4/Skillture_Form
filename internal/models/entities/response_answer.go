package entities

import (
	"time"

	"github.com/google/uuid"
)

type ResponseAnswer struct {
	ID         uuid.UUID      `db:"id" json:"id"`
	ResponseID uuid.UUID      `db:"response_id" json:"response_id"`
	FieldID    uuid.UUID      `db:"field_id" json:"field_id"`
	Value      map[string]any `db:"value" json:"value"` // JSONB
	CreatedAt  time.Time      `db:"created_at" json:"created_at"`
}

// TableName returns the DB table name
func (ResponseAnswer) TableName() string {
	return "response_answers"
}

// HasValue checks if the answer contains any value
func (a *ResponseAnswer) HasValue() bool {
	return len(a.Value) > 0
}

// GetValue returns the value for a given key in JSONB
func (a *ResponseAnswer) GetValue(key string) any {
	if val, ok := a.Value[key]; ok {
		return val
	}
	return nil
}
