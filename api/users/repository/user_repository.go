package repository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	GetProductById(ctx context.Context, tx *sql.Tx, userId int, productId int) (domain.Product, error)
	GetProductByUser(ctx context.Context, tx *sql.Tx, userId int) []domain.Product
	UpdateProductByUserSeller(ctx context.Context, db *sql.Tx, product domain.Product, userId int) domain.Product
	Purchase(ctx context.Context, tx *sql.Tx, userId int, productId int) (domain.Product, error)
	// Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	// Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
