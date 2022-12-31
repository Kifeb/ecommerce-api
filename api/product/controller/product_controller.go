package productController

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
