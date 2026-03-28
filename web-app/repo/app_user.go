package repo

import (
	"algtutor/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppUserRepo struct {
	dbPool *pgxpool.Pool
}

func NewAppUserRepo(pool *pgxpool.Pool) *AppUserRepo {
	return &AppUserRepo{pool}
}

func (ur AppUserRepo) CreateUser(u domain.AppUser) error {
	ctx := context.Background()
	insertQuery := `
		INSERT INTO app_user(email, role, password_hash, provider, provider_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := ur.dbPool.Exec(
		ctx,
		insertQuery,
		u.Email,
		u.Role,
		u.PasswordHash,
		u.Provider,
		u.ProviderId,
	)
	if err != nil {
		// TODO: custom error handling and stuff
		return err
	}
	return nil
}
