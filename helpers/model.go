package helpers

import (
	"ecommerce_api/model/domain"
	web "ecommerce_api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Phone:    user.Phone,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Picture:  product.Picture,
		Category: product.Category,
		Quantity: product.Quantity,
	}
}

func ToProductResponses(users []domain.Product) []web.ProductResponse {
	var userResponses []web.ProductResponse

	for _, user := range users {
		userResponses = append(userResponses, ToProductResponse(user))
	}

	return userResponses
}
