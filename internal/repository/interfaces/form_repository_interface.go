// internal/repository/interface/form_repository.go
package interfaces

import (
	"context"

	"Skillture_Form/internal/domain/entities"

	"github.com/google/uuid"
)

type FormRepository interface {

	// Create stores a form with its fields (transactional)
	Create(
		ctx context.Context,
		form *entities.Forms,
		fields []entities.FormField,
	) error

	// GetByID retrieves a form with its fields
	GetByID(
		ctx context.Context, id uuid.UUID) (*entities.Forms, []entities.FormField, error)

	// Update updates form metadata (title, description, status)
	Update(ctx context.Context, form *entities.Forms) error

	// Delete removes a form and all related fields
	Delete(ctx context.Context, id uuid.UUID) error
}
