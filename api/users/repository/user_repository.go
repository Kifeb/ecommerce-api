package repository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	// Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	// Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
