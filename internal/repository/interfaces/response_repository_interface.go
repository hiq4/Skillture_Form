package interfaces

import (
	"context"

	"Skillture_Form/internal/domain/entities"

	"github.com/google/uuid"
)

type ResponseRepository interface {

	// Create saves a new response
	Create(ctx context.Context, response *entities.Response) error

	// GetByID retrieves a response by its ID
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Response, error)

	// GetByFormID retrieves all responses for a specific form
	GetByFormID(ctx context.Context, formID uuid.UUID) ([]entities.Response, error)

	// GetByEmail retrieves all responses submitted with a specific email
	GetByEmail(ctx context.Context, email string) ([]entities.Response, error)

	// Delete removes a response (optional)
	Delete(ctx context.Context, id uuid.UUID) error
}
