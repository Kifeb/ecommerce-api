package service

import (
	"context"
	web "ecommerce_api/model/web"
)

type UserService interface {
	Create(ctx context.Context, r web.UserCreateRequest) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
