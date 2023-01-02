package repository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
	"errors"
	"fmt"
	"log"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO users(username, email, password, role, phone) VALUES (?, ?, ?, ?, ?)"

	go func(query string) {
		result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Role, user.Phone)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		user.Id = int(id)
	}(query)
	return user
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, username, email, role, phone FROM users"
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Close()

	var users []domain.User
	for result.Next() {
		user := domain.User{}
		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.Phone)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	return users
}

func (r *UserRepositoryImpl) GetProductById(ctx context.Context, tx *sql.Tx, userId int, productId int) (domain.Product, error) {
	query := "SELECT p.id, p.name, p.category, p.picture, p.quantity, p.price FROM product p INNER JOIN users u ON p.user_id = u.id WHERE u.id = ? AND p.id = ? AND u.role = 'seller'"

	result, err := tx.QueryContext(ctx, query, userId, productId)
	if err != nil {
		fmt.Println(err)
	}

	product := domain.Product{}
	if result.Next() {
		err := result.Scan(&product.Id, &product.Name, &product.Category, &product.Picture, &product.Quantity, &product.Price)
		if err != nil {
			fmt.Println(err)
		}
		return product, nil
	} else {
		return product, errors.New("Product Not Found")
	}
}

func (r *UserRepositoryImpl) GetProductByUser(ctx context.Context, tx *sql.Tx, userId int) []domain.Product {
	query := "SELECT p.id, p.name, p.category, p.picture, p.quantity, p.price FROM product p INNER JOIN users u ON p.user_id = u.id WHERE u.id = ?"
	result, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Close()

	var products []domain.Product
	for result.Next() {
		product := domain.Product{}
		err := result.Scan(&product.Id, &product.Name, &product.Category, &product.Picture, &product.Quantity, &product.Price)
		if err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}

	return products
}

func (r *UserRepositoryImpl) UpdateProductByUserSeller(ctx context.Context, db *sql.Tx, product domain.Product, userId int) (domain.Product, error) {

	query := "UPDATE product p JOIN users u ON p.user_id = u.id SET p.quantity = ?, p.price = ? WHERE u.id = ? AND u.role = 'seller'"

	_, err := db.ExecContext(ctx, query, product.Quantity, product.Price, userId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *UserRepositoryImpl) Purchase(ctx context.Context, tx *sql.Tx, userId int, productId int) (domain.Product, error) {
	query := "SELECT p.id, p.name, p.category, p.picture, p.quantity, p.price FROM product p INNER JOIN users u ON p.user_id = u.id WHERE u.id = ? AND p.id = ?"

	result, err := tx.QueryContext(ctx, query, userId, productId)
	if err != nil {
		fmt.Println(err)
	}

	product := domain.Product{}
	if result.Next() {
		err := result.Scan(&product.Id, &product.Name, &product.Category, &product.Picture, &product.Quantity, &product.Price)
		if err != nil {
			fmt.Println(err)
		}
		return product, nil
	}

	return product, errors.New("Product not found")
}
