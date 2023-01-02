package productRoute

import (
	productController "ecommerce_api/api/product/controller"

	"github.com/gorilla/mux"
)

func NewProductRouter(c productController.ProductController) *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/product", c.FindAll).Methods("GET")
	// router.HandleFunc("/api/product", c.Create).Methods("POST")

	// router.POST("/api/product", c.Create)
	// router.GET("/api/product", c.FindAll)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
