package controller

import (
	"ecommerce_api/api/users/service"
	"ecommerce_api/helpers"
	web "ecommerce_api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userResponses := c.UserService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   userResponses,
	}

	helpers.WriteToReqBody(w, webResponse)
}
