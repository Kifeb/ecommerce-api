package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
