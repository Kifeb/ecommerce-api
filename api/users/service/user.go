package service

import (
	"context"
	"database/sql"
	"ecommerce_api/api/users/repository"
	"ecommerce_api/exception"
	"ecommerce_api/helpers"
	"ecommerce_api/model/domain"
	web "ecommerce_api/model/web"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, r web.UserCreateRequest) web.UserResponse {
	err := s.Validate.Struct(r)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
		Phone:    r.Phone,
	}

	user = s.UserRepository.Save(ctx, tx, user)

	return helpers.ToUserResponse(user)
}

func (s *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	users := s.UserRepository.FindAll(ctx, tx)

	return helpers.ToUserResponses(users)
}

func (s *UserServiceImpl) GetProductById(ctx context.Context, userId int, productId int) web.ProductResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)
	product, err := s.UserRepository.GetProductById(ctx, tx, userId, productId)
	if err != nil {
		panic(err)
	}

	return helpers.ToProductResponse(product)
}

func (s *UserServiceImpl) GetProductByUser(ctx context.Context, userId int) []web.ProductResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	products := s.UserRepository.GetProductByUser(ctx, tx, userId)

	return helpers.ToProductResponses(products)

}

func (s *UserServiceImpl) UpdateProductByUserSeller(ctx context.Context, request web.ProductUpdateRequest, userId int) web.ProductResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)
	product, err := s.UserRepository.GetProductById(ctx, tx, userId, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	fmt.Println(product.Quantity)
	product.Quantity = request.Quantity
	product.Price = request.Price

	product = s.UserRepository.UpdateProductByUserSeller(ctx, tx, product, userId)
	fmt.Println(product)

	return helpers.ToProductResponse(product)
}

func (s *UserServiceImpl) Purchase(ctx context.Context, userId int, productId int) web.ProductResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	products, err := s.UserRepository.Purchase(ctx, tx, userId, productId)

	return helpers.ToProductResponse(products)
}
