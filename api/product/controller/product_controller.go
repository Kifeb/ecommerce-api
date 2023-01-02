package productController

import (
	"net/http"
)

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request, userId int)
	FindAll(w http.ResponseWriter, r *http.Request)
}
