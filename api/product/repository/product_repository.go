package productRepository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product, userId int) domain.Product
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	// Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	// Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
