package entities

import (
	"time"

	"github.com/google/uuid"
)

type ResponseAnswerVector struct {
	ID               uuid.UUID `db:"id" json:"id"`
	ResponseAnswerID uuid.UUID `db:"response_answer_id" json:"response_answer_id"`
	Embedding        []float32 `db:"embedding" json:"embedding"`
	ModelName        string    `db:"model_name,omitempty" json:"model_name,omitempty"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

// TableName returns the DB table name
func (ResponseAnswerVector) TableName() string {
	return "response_answer_vectors"
}

// HasEmbedding checks if the embedding exists
func (v *ResponseAnswerVector) HasEmbedding() bool {
	return len(v.Embedding) > 0
}
