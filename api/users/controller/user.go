package controller

import (
	"ecommerce_api/api/users/service"
	"ecommerce_api/helpers"
	web "ecommerce_api/model/web"
	"fmt"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userCreateReq := web.UserCreateRequest{}
	helpers.ReadFromReqBody(r, &userCreateReq)

	userResponse := c.UserService.Create(r.Context(), userCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   userResponse,
	}

	helpers.WriteToReqBody(w, webResponse)
}

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	userResponses := c.UserService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   userResponses,
	}

	helpers.WriteToReqBody(w, webResponse)
}

func (c *UserControllerImpl) GetProductByUser(w http.ResponseWriter, r *http.Request, id int) {
	// userId := mux.Vars(r)["userId"]
	// id, _ := strconv.Atoi(userId)
	productResponse := c.UserService.GetProductByUser(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}
	helpers.WriteToReqBody(w, webResponse)
}

func (c *UserControllerImpl) GetProducyById(w http.ResponseWriter, r *http.Request, userId int, productId int) {
	productResponse := c.UserService.GetProductById(r.Context(), userId, productId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helpers.WriteToReqBody(w, webResponse)
}

func (c *UserControllerImpl) UpdateProductBySeller(w http.ResponseWriter, r *http.Request, userId int, productId int) {

	productUpdateRequest := web.ProductUpdateRequest{}
	helpers.ReadFromReqBody(r, &productUpdateRequest)

	productUpdateRequest.Id = productId

	productResponse := c.UserService.UpdateProductByUserSeller(r.Context(), productUpdateRequest, userId)
	fmt.Println(productResponse)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helpers.WriteToReqBody(w, webResponse)
}

func (c *UserControllerImpl) Purchase(w http.ResponseWriter, r *http.Request, userId int, productId int) {
	productResponse := c.UserService.Purchase(r.Context(), userId, productId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helpers.WriteToReqBody(w, webResponse)
}
