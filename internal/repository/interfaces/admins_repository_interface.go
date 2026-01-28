package interfaces

import (
	"Skillture_Form/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

type AdminRepository interface {
	// Create Stores a new admin
	Create(ctx context.Context, admin *entities.Admin) error
	// GetByID retrieves an admin by its ID
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Admin, error)
	// GetByUsername retrieves an admin by username (For Login)
	GetByUsername(ctx context.Context, username string) (*entities.Admin, error)
}
