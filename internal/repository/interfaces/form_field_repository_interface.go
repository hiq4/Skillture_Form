package interfaces

import (
	"Skillture_Form/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

// Filter object
type FormFieldFilter struct {
	FormID *uuid.UUID
}

type FormFieldRepository interface {
	// Create saves
	Create(ctx context.Context, form *entities.Forms) error
	// GetByID retrieves an admin by their ID
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Forms, error)
	// Update modifies admin details
	Update(ctx context.Context, form *entities.Forms) error
	// Delete removes an admin
	Delete(ctx context.Context, id uuid.UUID) error
	// List retrieves forms based on optional filter
	List(ctx context.Context, filter FormFieldFilter) ([]*entities.FormField, error)
}
