package productRepository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
	"log"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product, userId int) domain.Product {
	query := "INSERT INTO product(name, price, picture, category, quantity, user_id) VALUES (?, ?, ?, ?, ?, ?) JOIN "

	result, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Picture, product.Category, product.Quantity, product.User_Id)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	product.Id = int(id)
	return product
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "SELECT id, name, picture, category, quantity FROM product"
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var products []domain.Product
	for result.Next() {
		product := domain.Product{}
		err := result.Scan(&product.Id, &product.Name, &product.Picture, &product.Category, &product.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return products
}
