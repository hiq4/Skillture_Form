package admin

import (
	"context"
	"errors"
	"time"

	"Skillture_Form/internal/domain/entities"
	repo "Skillture_Form/internal/repository/interfaces"

	"github.com/google/uuid"
)

// AdminUseCase is the exported struct for admin business logic
type AdminUseCase struct {
	adminRepo repo.AdminRepository
}

// NewAdminUseCase creates a new instance of AdminUseCase
func NewAdminUseCase(adminRepo repo.AdminRepository) *AdminUseCase {
	return &AdminUseCase{adminRepo: adminRepo}
}

// Create creates a new admin with username and password
func (uc *AdminUseCase) Create(ctx context.Context, username, password string) (*entities.Admin, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	// Example password hash (replace with bcrypt in production)
	hashed := HashPassword(password)

	admin := &entities.Admin{
		ID:             uuid.New(),
		Username:       username,
		HashedPassword: hashed,
		CreatedAt:      time.Now(),
	}

	if err := uc.adminRepo.Create(ctx, admin); err != nil {
		return nil, err
	}

	return admin, nil
}

// Authenticate validates an admin login attempt
func (uc *AdminUseCase) Authenticate(ctx context.Context, username, password string) (*entities.Admin, error) {
	admin, err := uc.adminRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !CheckPassword(password, admin.HashedPassword) {
		return nil, errors.New("invalid username or password")
	}

	return admin, nil
}

// List retrieves all admins
func (uc *AdminUseCase) List(ctx context.Context) ([]*entities.Admin, error) {
	return uc.adminRepo.List(ctx)
}

// Delete removes an admin by ID
func (uc *AdminUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adminRepo.Delete(ctx, id)
}

// --- Helper functions ---
// HashPassword is a placeholder function for hashing passwords
func HashPassword(password string) string {
	return password // Replace with proper hashing (bcrypt)
}

// CheckPassword compares plain password with hashed password
func CheckPassword(password, hashed string) bool {
	return password == hashed // Replace with proper hash comparison
}
