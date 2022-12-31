package route

import (
	"ecommerce_api/api/users/controller"
	"ecommerce_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewUserRouter(c controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", c.Create)
	router.GET("/api/users", c.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
