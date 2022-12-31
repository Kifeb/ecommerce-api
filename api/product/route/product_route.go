package productRoute

import (
	productController "ecommerce_api/api/product/controller"
	"ecommerce_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewProductRouter(c productController.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/product", c.Create)
	router.GET("/api/product", c.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
