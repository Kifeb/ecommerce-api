package repository

import (
	"context"
	"database/sql"
	"ecommerce_api/model/domain"
	"log"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO users(username, email, password, role, phone) VALUES (?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Role, user.Phone)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.Id = int(id)
	return user
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, username, email, role, phone FROM users"
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	var users []domain.User
	for result.Next() {
		user := domain.User{}
		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.Phone)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}
