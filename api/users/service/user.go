package service

import (
	"context"
	"database/sql"
	"ecommerce_api/api/users/repository"
	"ecommerce_api/helpers"
	"ecommerce_api/model/domain"
	web "ecommerce_api/model/web"

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
	defer helpers.CommirOrRollback(tx)

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
	defer helpers.CommirOrRollback(tx)

	users := s.UserRepository.FindAll(ctx, tx)

	return helpers.ToUserResponses(users)
}
