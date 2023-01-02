package controller

import (
	"net/http"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	GetProductByUser(w http.ResponseWriter, r *http.Request, id int)
	GetProducyById(w http.ResponseWriter, r *http.Request, userId int, product int)
	UpdateProductBySeller(w http.ResponseWriter, r *http.Request, userId int, productId int)
	Purchase(w http.ResponseWriter, r *http.Request, userId int, productId int)
}
