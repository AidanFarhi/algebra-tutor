package repo

import (
	"algtutor/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppUserRepo struct {
	db *pgxpool.Pool
}

func NewAppUserRepo(db *pgxpool.Pool) *AppUserRepo {
	return &AppUserRepo{db}
}

func (aur AppUserRepo) CreateUser(au domain.AppUser) error {
	ctx := context.Background()
	query := `
		INSERT INTO app_user(email, role, password_hash, provider, provider_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := aur.db.Exec(
		ctx,
		query,
		au.Email,
		au.Role,
		au.PasswordHash,
		au.Provider,
		au.ProviderId,
	)
	return err
}

func (aur AppUserRepo) GetByID(id int) (domain.AppUser, error) {
	ctx := context.Background()
	query := `
		SELECT id, email, role, password_hash, provider, provider_id, created_at
		FROM app_user
		WHERE id = $1
	`
	var au domain.AppUser
	err := aur.db.QueryRow(ctx, query, id).Scan(
		&au.Id,
		&au.Email,
		&au.Role,
		&au.PasswordHash,
		&au.Provider,
		&au.ProviderId,
		&au.CreatedAt,
	)
	return au, err
}

func (aur AppUserRepo) GetByEmail(email string) (domain.AppUser, error) {
	ctx := context.Background()
	query := `
		SELECT id, email, role, password_hash, provider, provider_id, created_at
		FROM app_user
		WHERE email = $1
	`
	var au domain.AppUser
	err := aur.db.QueryRow(ctx, query, email).Scan(
		&au.Id,
		&au.Email,
		&au.Role,
		&au.PasswordHash,
		&au.Provider,
		&au.ProviderId,
		&au.CreatedAt,
	)
	return au, err
}
