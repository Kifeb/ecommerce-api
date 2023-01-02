package service

import (
	"context"
	web "ecommerce_api/model/web"
)

type UserService interface {
	Create(ctx context.Context, r web.UserCreateRequest) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	GetProductById(ctx context.Context, userId int, productId int) web.ProductResponse
	GetProductByUser(ctx context.Context, userId int) []web.ProductResponse
	UpdateProductByUserSeller(ctx context.Context, request web.ProductUpdateRequest, userId int) web.ProductResponse
	Purchase(ctx context.Context, request web.ProductUpdateRequest, userId int) web.ProductResponse
}
