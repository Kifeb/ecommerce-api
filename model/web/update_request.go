package web

type ProductUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}
