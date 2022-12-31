package productController

import (
	productService "ecommerce_api/api/product/service"
	"ecommerce_api/helpers"
	web "ecommerce_api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService productService.ProductService
}

func NewProductController(productService productService.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateReq := web.ProductCreateRequest{}
	helpers.ReadFromReqBody(r, &userCreateReq)

	productResponse := c.ProductService.Create(r.Context(), userCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helpers.WriteToReqBody(w, webResponse)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productResponses := c.ProductService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   productResponses,
	}

	helpers.WriteToReqBody(w, webResponse)
}
