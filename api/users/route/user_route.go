package route

import (
	"ecommerce_api/api/users/controller"

	"github.com/gorilla/mux"
)

func NewUserRouter(c controller.UserController) *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/users", c.Create).Methods("POST")
	// router.HandleFunc("/api/users", c.FindAll).Methods("GET")
	// router.HandleFunc("/api/users/{userId}", c.GetProductByUser).Methods("GET")
	// router.POST("/api/users", c.Create)
	// router.GET("/api/users", c.FindAll)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
