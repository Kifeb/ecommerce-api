package productService

import (
	"context"
	web "ecommerce_api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, r web.ProductCreateRequest) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
