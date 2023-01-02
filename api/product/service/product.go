package productService

import (
	"context"
	"database/sql"
	productRepository "ecommerce_api/api/product/repository"
	"ecommerce_api/helpers"
	"ecommerce_api/model/domain"
	web "ecommerce_api/model/web"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository productRepository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository productRepository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, r web.ProductCreateRequest) web.ProductResponse {
	err := s.Validate.Struct(r)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	product := domain.Product{
		Name:     r.Name,
		Price:    r.Price,
		Picture:  r.Picture,
		Category: r.Category,
		Quantity: r.Quantity,
		User_Id:  r.User_Id,
	}

	product = s.ProductRepository.Save(ctx, tx, product)

	return helpers.ToProductResponse(product)
}

func (s *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.CommitOrRollback(tx)

	products := s.ProductRepository.FindAll(ctx, tx)

	return helpers.ToProductResponses(products)
}
