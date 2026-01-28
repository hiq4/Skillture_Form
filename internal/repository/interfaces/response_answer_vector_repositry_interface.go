package interfaces

import (
	"context"

	"Skillture_Form/internal/domain/entities"

	"github.com/google/uuid"
)

type ResponseAnswerVectorRepository interface {

	// Create stores a vector for a response answer
	Create(ctx context.Context, vector *entities.ResponseAnswerVector) error

	// GetByResponseAnswerID retrieves all vectors for a specific answer
	GetByResponseAnswerID(ctx context.Context, responseAnswerID uuid.UUID) ([]entities.ResponseAnswerVector, error)

	// DeleteByResponseAnswerID removes all vectors for an answer
	DeleteByResponseAnswerID(ctx context.Context, responseAnswerID uuid.UUID) error
}
