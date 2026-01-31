package postgres

import (
	"context"
	"errors"

	"Skillture_Form/internal/domain/entities"
	"Skillture_Form/internal/repository/interfaces"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type adminRepository struct {
	*BaseRepository // embed BaseRepository
}

// Constructor
func NewAdminRepository(pool *pgxpool.Pool) interfaces.AdminRepository {
	return &adminRepository{
		BaseRepository: NewBaseRepository(pool),
	}
}

// Create admin
func (r *adminRepository) Create(ctx context.Context, admin *entities.Admin) error {
	query := `
		INSERT INTO admins (id, username, password, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.Exec(ctx, query, admin.ID, admin.Username, admin.HashedPassword, admin.CreatedAt)
	return err
}

// GetByID
func (r *adminRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Admin, error) {
	query := `
		SELECT id, username, password, created_at
		FROM admins
		WHERE id = $1
	`
	row := r.QueryRow(ctx, query, id)
	var admin entities.Admin
	err := row.Scan(&admin.ID, &admin.Username, &admin.HashedPassword, &admin.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// GetByUsername
func (r *adminRepository) GetByUsername(ctx context.Context, username string) (*entities.Admin, error) {
	query := `
		SELECT id, username, password, created_at
		FROM admins
		WHERE username = $1
	`
	row := r.QueryRow(ctx, query, username)
	var admin entities.Admin
	err := row.Scan(&admin.ID, &admin.Username, &admin.HashedPassword, &admin.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
